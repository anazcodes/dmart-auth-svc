// payload is a request response struct package
package payload

import "time"

type AdminLogin struct {
	Username string
	Password string
}

type UserAccount struct {
	ID        uint
	Username  string
	Email     string
	Phone     int64 // cant use country code with it
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	LoginInput string // user can login using email, phone or using username
	Password   string
}
type Contact struct {
	Email string
	Phone int64
}

type Address struct {
	Name             string
	PhoneNumber      string
	PostalCode       string
	Locality         string
	AddressLine      string
	District         string
	Landmark         string
	AlternativePhone string
	UserID           uint
	IsDefault        bool
}

type Time struct {
	Now time.Time
}

type response struct {
	status int
	msg    string
	data   interface{}
	err    interface{}
}

// Function for passing response message
func Response(status int, msg string, data, err interface{}) response {

	return response{
		status: status,
		msg:    msg,
		data:   data,
		err:    err,
	}
}
