package models

// User is the user model
type User struct {
	UserID int64
}

// Trip is the trip model
type Trip struct {
	TripID     int64
	TripTitle  string
	PlacesList []Place
}

// Place is the place model
type Place struct {
	PlaceTitle       string
	PlaceDescription string
	PlaceLatitude    float64
	PlaceLongitude   float64
	PlaceAddress     string
}
