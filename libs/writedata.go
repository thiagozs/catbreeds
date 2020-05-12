package libs

import (
	"encoding/json"
	"hostgator-challenge/database"
	"hostgator-challenge/models"
	"net/url"

	"github.com/tidwall/gjson"
)

func WriteData(db *database.GormRepo, param string) ([]models.CatAPI, error) {

	configs := func(cfg *FetchDataConfig) {
		cfg.URL = "https://api.thecatapi.com/v1/breeds/search"
		cfg.TokenHeader = "2f3a8279-8e3f-482b-867d-1315d2b22c6f"
		cfg.TimeOut = 40
	}

	fetch := NewFetchData(configs)
	var dataApi []models.CatAPI
	result, err := fetch.GetJSON(url.QueryEscape(param))
	if err != nil {
		return []models.CatAPI{}, err
	}

	if result.IsArray() && len(result.Array()) == 1 {
		var catapi models.CatAPI
		var value gjson.Result = result.Array()[0]
		err := json.Unmarshal([]byte(value.String()), &catapi)
		if err != nil {
			return []models.CatAPI{}, err
		}
		if err := db.Create(&catapi); err != nil {
			return []models.CatAPI{}, err
		}
		dataApi = append(dataApi, catapi)

	} else {
		result.ForEach(func(key, value gjson.Result) bool {
			var catapi models.CatAPI
			err := json.Unmarshal([]byte(value.String()), &catapi)
			if err != nil {
				return false //break loop
			}
			if err := db.Create(&catapi); err != nil {
				return false // break loop
			}
			dataApi = append(dataApi, catapi)

			return true // keep iterating
		})
	}

	return dataApi, nil
}
