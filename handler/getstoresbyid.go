package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ngc13/logger"
	"ngc13/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (r *Repo) GetDetailStore(c echo.Context) error {
	param_id := c.Param("id")
	id, _ := strconv.Atoi(param_id)

	// find store id
	var s model.Stores
	res := r.DB.First(&s, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "there is no such stores")
		}
		logger.Logging(c).Error(res.Error)
		c.JSON(http.StatusInternalServerError, "Internal server Error")
	}

	logger.Logging(c).Info("Store:", s)

	// hit the API

	apiUrl := fmt.Sprintf("https://api.open-meteo.com/v1/dwd-icon?latitude=%.2f&longitude=%2.f&daily=apparent_temperature_max", s.Latitude, s.Longitude)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Logging(c).Error("error creating request:", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logging(c).Error("error getting response:", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Logging(c).Error("error reading body:", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	var weatherResponse model.WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		logger.Logging(c).Error("error unmarshall:", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"store":   s,
			"weather": weatherResponse,
		},
	)

	/*
		example output
		{
			"store": {
				"StoreID": 3,
				"StoreName": "Store C",
				"StoreAddress": "789 Oak St",
				"Longitude": 34.56,
				"Latitude": 78.9,
				"Rating": 3
			},
			"weather": {
				"longitude": 35,
				"latitude": 78.875,
				"generationtime_ms": 0.04303455352783203,
				"utc_offset_seconds": 0,
				"timezone": "GMT",
				"timezone_abbreviation": "GMT",
				"elevation": 0,
				"daily_units": {
					"time": "iso8601",
					"apparent_temperature_max": "°C"
				},
				"daily": {
					"time": [
						"2024-06-05",
						"2024-06-06",
						"2024-06-07",
						"2024-06-08",
						"2024-06-09",
						"2024-06-10",
						"2024-06-11"
						],
						"apparent_temperature_max": [
							-4.2,
							-4.7,
							-5,
							-6.3,
							-5.2,
							-3.9,
							-5.4
							]
						}
					}
				}
			}
		}
	*/
}
