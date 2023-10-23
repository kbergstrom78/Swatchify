package database

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// Connect function initializes the database connection and returns a *gorm.DB instance.
func Connect() (*gorm.DB, error) {
	dsn := "user=kimbergstrom dbname=swatchify sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	err = db.AutoMigrate(
		&User{},
		&SwatchImage{},
		&GaugeCalculation{},
		&Material{},
		&UserPreference{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Users table
type User struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PasswordHash string `json:"password_hash"`
	DateJoined  time.Time `json:"date_joined"`
	LastLogin   time.Time `json:"last_login"`
}
// SwatchImages table
type SwatchImage struct {
	ImageID    string `json:"image_id"`
	UserID     string `json:"user_id"`
	ImagePath  string `json:"image_path"`
	UploadDate time.Time `json:"upload_date"`
}

// GaugeCalculations table
type GaugeCalculation struct {
	CalculationID string `json:"calculation_id"`
	ImageID       string `json:"image_id"`
	StitchCount   int    `json:"stitch_count"`
	RowCount      int    `json:"row_count"`
	StitchGauge   float64 `json:"stitch_gauge"`
	RowGauge      float64 `json:"row_gauge"`
	CalculationDate time.Time `json:"calculation_date"` // could be time.Time
}

// Materials table
type Material struct {
	MaterialID  string  `json:"material_id"`
	UserID      string  `json:"user_id"`
	YarnType    string  `json:"yarn_type"`
	YarnWeight  string  `json:"yarn_weight"`
	YarnFiber   string  `json:"yarn_fiber"`
	StitchType  string  `json:"stitch_type"`
	NeedleSize  float64 `json:"needle_size"`
}

// UserPreferences table
type UserPreference struct {
	PreferenceID   string  `json:"preference_id"`
	UserID         string  `json:"user_id"`
	DefaultYarnType string  `json:"default_yarn_type"`
	DefaultNeedleSize float64 `json:"default_needle_size"`
}