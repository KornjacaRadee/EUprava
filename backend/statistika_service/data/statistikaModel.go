package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type ZavodZaStatistiku struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Naziv   string             `bson:"naziv" json:"naziv"`
	Prestup []Prestupi         `bson:"prestup" json:"prestup"`
}

type Prestupi struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VrstaPrestupa         string             `bson:"vrstaPrestupa" json:"vrstaPrestupa"`
	BrojGodisnjihPrestupa int                `bson:"brojGodisnjihPrestupa" json:"brojGodisnjihPrestupa"`
	Statistika            []Statistika       `bson:"statistika" json:"statistika"`
	Vozilo                []Vozila           `bson:"vozilo" json:"vozilo"`
}

type Statistika struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Regija string             `bson:"regija" json:"regija"`
	Datum  primitive.DateTime `bson:"datum" json:"datum"`
	Grad   string             `bson:"grad" json:"grad"`
}

type Vozila struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GodinaProizvodnje string             `bson:"godinaProizvodnje" json:"godinaProizvodnje"`
	VrstaVozila       string             `bson:"vrstaVozila" json:"vrstaVozila"`
}
