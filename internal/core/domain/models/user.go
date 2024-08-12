package models

import (
	"errors"
	"gopkg.in/encoder.v1"
)

type AuthenticatedUser struct {
	Id       int    `json:"id"  form:"id"  db:"id"`
	Username string `json:"username"  form:"username"  db:"username"`
}
type PasswordPair struct {
	OriginalPassword string
	HashedPassword   string
}

func (p *PasswordPair) CheckRequestPassword() error {
	encoding := encoder.NewBcryptEncoder()
	verify, err := encoding.Verify(p.HashedPassword, p.OriginalPassword)
	if err != nil {
		return err
	}
	if !verify {
		return errors.New("password is incorrect")
	}
	return nil
}
func (p *PasswordPair) SetOriginalPasswordAndGetHashedPassword(originalPassword string) (string, error) {
	encryptor := encoder.NewBcryptEncoder()
	hashedPassword, err := encryptor.Encode(originalPassword)
	p.OriginalPassword = originalPassword
	p.HashedPassword = hashedPassword
	return hashedPassword, err
}

type Name struct {
	Title string `json:"title" db:"name_title"`
	First string `json:"first" db:"name_first"`
	Last  string `json:"last" db:"name_last"`
}

type Street struct {
	Number int    `json:"number" db:"street_number"`
	Name   string `json:"name" db:"street_name"`
}

type Coordinates struct {
	Latitude  string `json:"latitude" db:"coordinates_latitude"`
	Longitude string `json:"longitude" db:"coordinates_longitude"`
}

type Timezone struct {
	Offset      string `json:"offset" db:"timezone_offset"`
	Description string `json:"description" db:"timezone_description"`
}

type Location struct {
	Street      Street      `json:"street" db:"location_street"`
	City        string      `json:"city" db:"location_city"`
	State       string      `json:"state" db:"location_state"`
	Country     string      `json:"country" db:"location_country"`
	Postcode    int         `json:"postcode" db:"location_postcode"`
	Coordinates Coordinates `json:"coordinates" db:"location_coordinates"`
	Timezone    Timezone    `json:"timezone" db:"location_timezone"`
}

type Login struct {
	UUID     string `json:"uuid" db:"login_uuid"`
	Username string `json:"username" db:"login_username"`
	Password string `json:"password" db:"login_password"`
	Salt     string `json:"salt" db:"login_salt"`
	MD5      string `json:"md5" db:"login_md5"`
	SHA1     string `json:"sha1" db:"login_sha1"`
	SHA256   string `json:"sha256" db:"login_sha256"`
}

type Dob struct {
	Date string `json:"date" db:"dob_date"`
	Age  int    `json:"age" db:"dob_age"`
}

type Registered struct {
	Date string `json:"date" db:"registered_date"`
	Age  int    `json:"age" db:"registered_age"`
}

type Id struct {
	Name  string `json:"name" db:"id_name"`
	Value string `json:"value" db:"id_value"`
}

type Picture struct {
	Large     string `json:"large" db:"picture_large"`
	Medium    string `json:"medium" db:"picture_medium"`
	Thumbnail string `json:"thumbnail" db:"picture_thumbnail"`
}

type User struct {
	ID         int        `json:"ID" db:"id"`
	Gender     string     `json:"gender" db:"gender"`
	Name       Name       `json:"name"`
	Location   Location   `json:"location"`
	Email      string     `json:"email" db:"email"`
	Login      Login      `json:"login"`
	Dob        Dob        `json:"dob"`
	Registered Registered `json:"registered"`
	Phone      string     `json:"phone" db:"phone"`
	Cell       string     `json:"cell" db:"cell"`
	Id         Id         `json:"id"`
	Picture    Picture    `json:"picture"`
	Nat        string     `json:"nat" db:"nat"`
}

func (u *User) NewUserByUserDB(userDB *UserDB) error {
	u.ID = userDB.ID
	u.Gender = userDB.Gender
	u.Email = userDB.Email
	u.Phone = userDB.Phone
	u.Cell = userDB.Cell
	u.Nat = userDB.Nat

	u.Name = Name{
		Title: userDB.NameTitle,
		First: userDB.NameFirst,
		Last:  userDB.NameLast,
	}

	u.Location = Location{
		Street: Street{
			Number: userDB.StreetNumber,
			Name:   userDB.StreetName,
		},
		City:     userDB.LocationCity,
		State:    userDB.LocationState,
		Country:  userDB.LocationCountry,
		Postcode: userDB.LocationPostcode,
		Coordinates: Coordinates{
			Latitude:  userDB.CoordinatesLatitude,
			Longitude: userDB.CoordinatesLongitude,
		},
		Timezone: Timezone{
			Offset:      userDB.TimeZoneOffset,
			Description: userDB.TimeZoneDescription,
		},
	}

	u.Login = Login{
		UUID:     userDB.LoginUUID,
		Username: userDB.LoginUsername,
		Password: userDB.LoginPassword,
		Salt:     userDB.LoginSalt,
		MD5:      userDB.LoginMD5,
		SHA1:     userDB.LoginSHA1,
		SHA256:   userDB.LoginSHA256,
	}

	u.Dob = Dob{
		Date: userDB.DobDate,
		Age:  userDB.DobAge,
	}

	u.Registered = Registered{
		Date: userDB.RegisteredDate,
		Age:  userDB.RegisteredAge,
	}

	u.Id = Id{
		Name:  userDB.IdName,
		Value: userDB.IdValue,
	}

	u.Picture = Picture{
		Large:     userDB.PictureLarge,
		Medium:    userDB.PictureMedium,
		Thumbnail: userDB.PictureThumbnail,
	}

	return nil
}

type PasswordFields struct {
	UpdatePassword bool   `json:"update_password"`
	NewPassword    string `json:"new_password"`
}

type UserFields struct {
	UpdateUsername bool   `json:"update_user_name"`
	NewUsername    string `json:"new_user_name"`
}

type UserUpdateRequest struct {
	User           User           `json:"user"`
	UserType       string         `json:"user_type"`
	PasswordFields PasswordFields `json:"password_fields"`
	UserFields     UserFields     `json:"user_fields"`
}

type UserFromRequest struct {
	User     User   `json:"user"`
	UserType string `json:"user_type"`
}

type UserDB struct {
	ID                   int    `db:"id"`
	Gender               string `db:"gender"`
	NameTitle            string `db:"name_title"`
	NameFirst            string `db:"name_first"`
	NameLast             string `db:"name_last"`
	LocationId           int    `db:"location_id"`
	StreetNumber         int    `db:"street_number"`
	StreetName           string `db:"street_name"`
	LocationCity         string `db:"city"`
	LocationState        string `db:"state"`
	LocationCountry      string `db:"country"`
	LocationPostcode     int    `db:"postcode"`
	CoordinatesLatitude  string `db:"coordinates_latitude"`
	CoordinatesLongitude string `db:"coordinates_longitude"`
	TimeZoneOffset       string `db:"timezone_offset"`
	TimeZoneDescription  string `db:"timezone_description"`
	Email                string `json:"email" db:"email"`
	LoginUUID            string `db:"login_uuid"`
	LoginUsername        string `db:"login_username"`
	LoginPassword        string `db:"login_password"`
	LoginSalt            string `db:"login_salt"`
	LoginMD5             string `db:"login_md5"`
	LoginSHA1            string `db:"login_sha1"`
	LoginSHA256          string `db:"login_sha256"`
	DobDate              string `db:"dob_date"`
	DobAge               int    `db:"dob_age"`
	RegisteredDate       string `db:"registered_date"`
	RegisteredAge        int    `db:"registered_age"`
	Phone                string `db:"phone"`
	Cell                 string `db:"cell"`
	IdName               string `db:"id_name"`
	IdValue              string `db:"id_value"`
	PictureLarge         string `db:"picture_large"`
	PictureMedium        string `db:"picture_medium"`
	PictureThumbnail     string `db:"picture_thumbnail"`
	Nat                  string `db:"nat"`
}
