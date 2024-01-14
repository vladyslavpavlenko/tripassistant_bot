package models

// User is the user model
type User struct {
	UserID   int64  `firestore:"user_id"`
	UserName string `firestore:"user_name"`
}

// Trip is the trip model
type Trip struct {
	TripID     int64   `firestore:"trip_id"`
	TripPlaces []Place `firestore:"trip_places"`
}

// Place is the place model
type Place struct {
	PlaceID        string  `firestore:"place_id"`
	PlaceTitle     string  `firestore:"place_title"`
	PlaceLatitude  float64 `firestore:"place_latitude"`
	PlaceLongitude float64 `firestore:"place_longitude"`
	PlaceAddress   string  `firestore:"place_address"`
}
