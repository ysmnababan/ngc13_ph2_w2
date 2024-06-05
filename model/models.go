package model

type User struct {
	UserID        uint    `json:"user_id" gorm:"primaryKey"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

type Product struct {
	ProductID uint    `gorm:"primaryKey" json:"product_id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

type ProductRequest struct {
	ProductID uint    `gorm:"primaryKey" json:"product_id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
	StoreID   uint    `json:"store_id"`
}

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	UserID        uint
	ProductID     uint
	Quantity      int
	TotalAmount   float64
	StoreID       uint
}

type Stores struct {
	StoreID      uint
	StoreName    string
	StoreAddress string
	Longitude    float64
	Latitude     float64
	Rating       int
}

type WeatherResponse struct {
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	Gentime      float64   `json:"generationtime_ms"`
	UtcOffset    int       `json:"utc_offset_seconds"`
	Timezone     string    `json:"timezone"`
	TimezoneAbbr string    `json:"timezone_abbreviation"`
	Elevation    float64   `json:"elevation"`
	DailyUnit    DailyUnit `json:"daily_units"`
	Daily        Daily     `json:"daily"`
}

type DailyUnit struct {
	Time    string `json:"time"`
	TempMax string `json:"apparent_temperature_max"`
}

type Daily struct {
	Time    []string  `json:"time"`
	TempMax []float64 `json:"apparent_temperature_max"`
}

// {
//     "latitude": 52.52,
//     "longitude": 13.419998,
//     "generationtime_ms": 0.14400482177734375,
//     "utc_offset_seconds": 0,
//     "timezone": "GMT",
//     "timezone_abbreviation": "GMT",
//     "elevation": 38.0,
//     "daily_units": {
//         "time": "iso8601",
//         "apparent_temperature_max": "°C"
//     },
//     "daily": {
//         "time": [
//             "2024-06-05",
//             "2024-06-06",
//             "2024-06-07",
//             "2024-06-08",
//             "2024-06-09",
//             "2024-06-10",
//             "2024-06-11"
//         ],
//         "apparent_temperature_max": [
//             16.8,
//             16.9,
//             16.6,
//             21.1,
//             16.8,
//             19.5,
//             17.5
//         ]
//     }
// }
