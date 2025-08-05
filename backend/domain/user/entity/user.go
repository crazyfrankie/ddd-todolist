package entity

type User struct {
	UserID int64

	UniqueName string // unique name
	Name       string // nickname
	Email      string // email
	IconURI    string // avatar URI
	IconURL    string // avatar URL

	CreatedAt int64 // creation time
	UpdatedAt int64 // update time
}
