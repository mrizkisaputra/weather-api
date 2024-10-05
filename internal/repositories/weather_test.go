package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load("./../../config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestWeatherRepository_FetchWeather(t *testing.T) {
	var repository = NewWeatherRepository(os.Getenv("WEATHER.API_KEY"))
	//response, err := repository.FetchWeather("jakarta, indonesia", "", "")
	response, err := repository.FetchWeather("jakarta, indonesia", "2024-10-04", "2024-10-04")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEqual(t, "", response)

	fmt.Println(err)
	fmt.Println(response)
}
