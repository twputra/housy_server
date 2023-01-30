package models

import (
	"encoding/json"
	"time"
)

type House struct {
	ID        int             `jsonz:"id" grom:"primary_key:auto_increment"`
	Name      string          `json:"name" from:"name" gorm:"type: varchar(255)"`
	City      CityResponse    `json:"city" gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CityId    int             `json:"city_id"`
	Address   string          `json:"address" from:"address" gorm:": varchar(255)"`
	Price     int             `json:"price"`
	TypeRent  string          `json:"type_rent" from:"type_rent" gorm:"type: varchar(255)"`
	Ameneties json.RawMessage `json:"ameneties" from:"ameneties" gorm:"type: json"`
	BedRoom   string          `json:"bed_room" from:"bed_room" gorm:"type: varchar(255)"`
	BathRoom  string          `json:"bath_room" from:"bath_room" gorm:"type: varchar(255)"`
	Image     string          `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type HouseResponse struct {
	ID        int             `jsonz:"id"`
	Name      string          `json:"name" `
	CityId    int             `json:"city_id"`
	Address   string          `json:"address"`
	Price     int             `json:"price"`
	TypeRent  string          `json:"type_rent"`
	Ameneties json.RawMessage `json:"ameneties"`
	BedRoom   string          `json:"bed_room"`
	BathRoom  string          `json:"bath_room"`
	Image     string          `json:"image"`
}

func (HouseResponse) TableName() string {
	return "houses"
}
