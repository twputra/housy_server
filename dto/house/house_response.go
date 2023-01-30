package housesdto

type HouseResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" from:"name" validate:"required"`
	CityId    int    `json:"city_id" from:"city_id" validate:"required"`
	Address   string `json:"address" from:"address" validate:"required"`
	Price     int    `json:"price" from:"price" validate:"required"`
	TypeRent  string `json:"type_rent" from:"type_rent" validate:"required"`
	Ameneties string `json:"ameneties" from:"ameneties" validate:"required"`
	BedRoom   string `json:"bed_room" from:"bed_room" validate:"required"`
	BathRoom  string `json:"bath_room" from:"bath_room" validate:"required"`
	Image     string `json:"image" from:"image" validate:"required"`
}
