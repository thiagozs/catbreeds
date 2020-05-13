package libs

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/require"
)

func TestFetchData(t *testing.T) {

	configs := func(cfg *FetchDataConfig) {
		cfg.URL = "https://api.thecatapi.com/v1/breeds/search"
		cfg.TokenHeader = "2f3a8279-8e3f-482b-867d-1315d2b22c6f"
		cfg.TimeOut = 40
	}

	data := NewFetchData(configs)

	result, err := data.GetJSON("sian")
	if err != nil {
		t.Error(err)
	}

	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result.Array()), 0)

}
