package model

type Country struct {
	Name           string                   `json:"name" bson:"name"`
	Iso3           string                   `json:"iso3,omitempty" bson:"iso3"`
	Iso2           string                   `json:"iso2" bson:"iso2"`
	NumberCode     string                   `json:"number_code,omitempty" bson:"number_code"`
	PhoneCode      string                   `json:"phone_code,omitempty" bson:"phone_code"`
	Capital        string                   `json:"capital,omitempty" bson:"capital"`
	Currency       string                   `json:"currency,omitempty" bson:"currency"`
	CurrencyName   string                   `json:"currency_name,omitempty" bson:"currency_name"`
	CurrencySymbol string                   `json:"currency_symbol,omitempty" bson:"currency_symbol"`
	Tld            string                   `json:"tld,omitempty" bson:"tld"`
	Native         string                   `json:"native" bson:"native"`
	Region         string                   `json:"region,omitempty" bson:"region"`
	Subregion      string                   `json:"subregion,omitempty" bson:"subregion"`
	Timezones      []map[string]interface{} `json:"timezones,omitempty" bson:"timezones"`
	Latitude       float64                  `json:"latitude,omitempty" bson:"latitude"`
	Longitude      float64                  `json:"longitude,omitempty" bson:"longitude"`
	Emoji          string                   `json:"emoji,omitempty" bson:"emoji"`
	EmojiU         string                   `json:"emojiU,omitempty" bson:"emojiU"`
}

type State struct {
	Name        string  `json:"name" bson:"name"`
	CountryCode string  `json:"country_code,omitempty" bson:"country_code"`
	StateCode   string  `json:"state_code" bson:"state_code"`
	Type        string  `json:"type,omitempty" bson:"type"`
	Latitude    float64 `json:"latitude,omitempty" bson:"latitude"`
	Longitude   float64 `json:"longitude,omitempty" bson:"longitude"`
}

type City struct {
	Name        string  `json:"name" bson:"name"`
	CountryCode string  `json:"country_code,omitempty" bson:"country_code"`
	StateCode   string  `json:"state_code,omitempty" bson:"state_code"`
	Latitude    float64 `json:"latitude,omitempty" bson:"latitude"`
	Longitude   float64 `json:"longitude,omitempty" bson:"longitude"`
}
