package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"hostgator-challenge/api/database"
	"hostgator-challenge/api/models"

	"github.com/stretchr/testify/assert"
)

func setupTestCase(t *testing.T) (func(t *testing.T), database.IGormRepo) {
	dbFile := filepath.Base("test_acc.db")
	d, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		t.Error(err)
	}
	d.AutoMigrate(models.Login{})

	db := database.NewGormRepo(d)

	return func(t *testing.T) {
		err := os.Remove(dbFile)
		if err != nil {
			t.Error(err)
		}
		d.Close()
	}, db
}

func performRequest(r http.Handler, method,
	path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAccountHandlerOK(t *testing.T) {
	teardownTestCase, db := setupTestCase(t)
	defer teardownTestCase(t)

	ctl := NewCtlRepo(db)

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()

	s.GET("/account/:id", ctl.FindAccount)

	model := models.Login{UserName: "thiagozs", Password: "xxxx"}
	if err := db.Create(&model); err != nil {
		t.Error(err)
	}

	w := performRequest(s, "GET", "/account/1")

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Login
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, response.UserName, model.UserName)
	assert.Equal(t, response.Password, model.Password)
}

func TestAccountHandlerIDError(t *testing.T) {
	teardownTestCase, db := setupTestCase(t)
	defer teardownTestCase(t)

	ctl := NewCtlRepo(db)

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()

	s.GET("/account/:id", ctl.FindAccount)

	w := performRequest(s, "GET", "/account/aaa")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	values, ok := response["error"]

	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Contains(t, values, "parsing")

}
