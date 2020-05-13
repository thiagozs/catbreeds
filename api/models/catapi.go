package models

import "time"

// CatAPI model
type CatAPI struct {
	ID               string  `json:"id" gorm:"primary_key;unique;not null;"`
	Name             string  `json:"name" gorm:"size(200);unique;not null;"`
	VcahospitalsURL  string  `json:"vcahospitals_url" gorm:"size(200)"`
	Temperament      string  `json:"temperament" gorm:"size(200)"`
	Origin           string  `json:"origin" gorm:"size(200)"`
	CountryCodes     string  `json:"country_codes" gorm:"size(10)"`
	CountryCode      string  `json:"country_code" gorm:"size(10)"`
	Description      string  `json:"description" gorm:"size(200)"`
	LifeSpan         string  `json:"life_span" gorm:"size(200)"`
	Indoor           int     `json:"indoor"`
	AltNames         string  `json:"alt_names" gorm:"size(150)"`
	Adaptability     int     `json:"adaptability"`
	AffectionLevel   int     `json:"affection_level"`
	ChildFriendly    int     `json:"child_friendly"`
	DogFriendly      int     `json:"dog_friendly"`
	EnergyLevel      int     `json:"energy_level"`
	Grooming         int     `json:"grooming"`
	HealthIssues     int     `json:"health_issues"`
	Intelligence     int     `json:"intelligence"`
	SheddingLevel    int     `json:"shedding_level"`
	SocialNeeds      int     `json:"social_needs"`
	StrangerFriendly int     `json:"stranger_friendly"`
	Vocalisation     int     `json:"vocalisation"`
	Experimental     int     `json:"experimental"`
	Hairless         int     `json:"hairless"`
	Natural          int     `json:"natural"`
	Rare             int     `json:"rare"`
	Rex              int     `json:"rex"`
	SuppressedTail   int     `json:"suppressed_tail"`
	ShortLegs        int     `json:"short_legs"`
	WikipediaURL     string  `json:"wikipedia_url" gorm:"size(200)"`
	Hypoallergenic   int     `json:"hypoallergenic"`
	Weight           *Weight `json:"weight" gorm:"embedded"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}
type Weight struct {
	Imperial string `json:"imperial" gorm:"size(10)"`
	Metric   string `json:"metric" gorm:"size(10)"`
}

func (l *CatAPI) TableName() string {
	return "catapis"
}
