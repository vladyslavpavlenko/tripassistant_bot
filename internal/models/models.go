package models

// User is the user model
type User struct {
	UserID   int64
	UserName string
}

// Trip is the trip model
type Trip struct {
	TripID int64
	// TripTitle  string
	TripPlaces []Place
}

// Place is the place model
type Place struct {
	PlaceTitle     string
	PlaceLatitude  float64
	PlaceLongitude float64
	PlaceAddress   string
}
