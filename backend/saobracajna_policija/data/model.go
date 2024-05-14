package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_Name *string            `bson:"first_name" json:"name" validate:"required"`
	Last_Name  *string            `bson:"last_name" json:"last_name"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
	Address    *string            `bson:"address" json:"address"`
	Role       string             `bson:"role" json:"role"`
}

type Prekrsaj struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`

	//Vreme        primitive.DateTime       `bson:"vreme,omitempty" json:"vreme"`
	//TipPrekrsaja TipSaobracajnogPrekrsaja `bson:"tipPrekrsaja,omitempty" json:"tipPrekrsaja"`
}

type Nesreca struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
	//Vreme      primitive.DateTime    `bson:"vreme,omitempty" json:"vreme"`
	//TipNesrece TipSaobracajneNesrece `bson:"tipNesrece,omitempty" json:"tipNesrece"`
}

type Users []*User

// Functions to encode and decode products to json and from json.
// If we inject an interface, we achieve dependency injection, meaning that anything that implements this interface can be passed down
// For us, it will be ResponseWriter, but it also may be a file writer or something similar.
func (u *Users) ToJSON(w io.Writer) error {
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
}

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
