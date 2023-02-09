package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type JSONTime struct {
	time.Time
}
type ShortenEntity struct {
	UrlToShorten     string    `json:"urlToShorten"`
	AccessKey        *string   `json:"accessKey"`
	ShortedUrl       string    `json:"shortedUrl" gorm:"primaryKey;<-:create"`
	CreatedAt        JSONTime  `gorm:"type:timestamp;default:current_timestamp;autoCreateTime" json:"-"`
	UpdatedAt        JSONTime  `gorm:"type:timestamp;default:current_timestamp;autoUpdateTime" json:"-"`
	LastAccessedDate *JSONTime `gorm:"type:TIMESTAMP NULL" json:"lastAccessedDate"`
	ExpireDate       *JSONTime `gorm:"type:date;default:null" json:"expireDate"`
}
type ShortenEntityResponse struct {
	UrlToShorten     string  `json:"urlToShorten"`
	AccessKey        *string `json:"accessKey"`
	ShortedUrl       string  `json:"shortedUrl"`
	ExpireDate       *string `json:"expireDate"`
	LastAccessedDate *string `json:"lastAccessedDate"`
}

func (shorten ShortenEntity) MarshalJSON() ([]byte, error) {
	var expireDateStr *string
	if shorten.ExpireDate != nil {
		t := shorten.ExpireDate.Time.Format(time.DateOnly)
		expireDateStr = &t
	} else {
		expireDateStr = nil
	}
	var accessedDateStr *string
	if shorten.LastAccessedDate != nil {
		t := shorten.LastAccessedDate.Time.Format(time.DateTime)
		accessedDateStr = &t
	} else {
		accessedDateStr = nil
	}
	return json.Marshal(ShortenEntityResponse{
		UrlToShorten:     shorten.UrlToShorten,
		AccessKey:        shorten.AccessKey,
		ShortedUrl:       shorten.ShortedUrl,
		ExpireDate:       expireDateStr,
		LastAccessedDate: accessedDateStr,
	})
}
func (JsonTime JSONTime) Value() (driver.Value, error) {
	return JsonTime.Time, nil
}
func (JsonTime *JSONTime) Scan(value interface{}) error {
	if value == nil {
		*JsonTime = JSONTime{time.Now()}
		return nil
	}
	if readTime, err := driver.DefaultParameterConverter.ConvertValue(value); err == nil {
		if convertedTime, ok := readTime.(time.Time); ok {
			*JsonTime = JSONTime{convertedTime}
			return nil
		}
	}
	return nil
}

type ShortenEntityVisitEntity struct {
	ID         int32  `gorm:"primaryKey;autoIncrement"`
	ShortedUrl string `json:"shortedUrl"`
}
