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
	Password       string
	HashedPassword string
}

func (passwordPair *PasswordPair) CheckPasswordRequest() error {
	encoding := encoder.NewBcryptEncoder()
	verify, err := encoding.Verify(passwordPair.HashedPassword, passwordPair.Password)
	if err != nil {
		return err
	}
	if !verify {
		return errors.New("password is incorrect")
	}
	return nil
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    string      `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

type Dob struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Registered struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Id struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type User struct {
	Gender     string     `json:"gender"`
	Name       Name       `json:"name"`
	Location   Location   `json:"location"`
	Email      string     `json:"email"`
	Login      Login      `json:"login"`
	Dob        Dob        `json:"dob"`
	Registered Registered `json:"registered"`
	Phone      string     `json:"phone"`
	Cell       string     `json:"cell"`
	Id         Id         `json:"id"`
	Picture    Picture    `json:"picture"`
	Nat        string     `json:"nat"`
}
