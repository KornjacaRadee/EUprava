package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

/* type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_Name *string            `bson:"first_name" json:"name" validate:"required"`
	Last_Name  *string            `bson:"last_name" json:"last_name"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
	Address    *string            `bson:"address" json:"address"`
	Role       string             `bson:"role" json:"role"`
} */

type Prekrsaj struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
}

type Nesreca struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
}

/* type Users []*User */

/* func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (p *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
} */

func (o *Prekrsaj) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Prekrsaj) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Nesreca) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Nesreca) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}
