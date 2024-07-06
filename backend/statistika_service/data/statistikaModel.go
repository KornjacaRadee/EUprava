package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type StatistikaPrekrsaja struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GodisnjiBrojPrekrsaja string             `bson:"godisnjiBrojPrekrsaja" json:"godisnjibBojPrekrsaja"`
	NajcescaLokacija      string             `bson:"najcescaLokacija" json:"najcescaLokacija"`
	NajcesceVozilo        string             `bson:"najcesceVozilo" json:"najcesceVozilo"`
}
type StatistikaNesreca struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GodisnjiBrojNesreca string             `bson:"godisnjiBrojNesreca" json:"godisnjibBojNesreca"`
	NajcescaLokacija    string             `bson:"najcescaLokacija" json:"najcescaLokacija"`
	NajcesceVozilo      string             `bson:"najcesceVozilo" json:"najcesceVozilo"`
}

type StatistikaVozackihDozvola struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GodisnjiBrojIzdatihDozvola string             `bson:"godisnjiBrojIzdatihDozvola" json:"godisnjiBrojIzdatihDozvola"`
	NajcescaKategorija         string             `bson:"najcescaKategorija" json:"najcescaKategorija"`
	GodisnjiBrojKaznenihBodova string             `bson:"godisnjiBrojKaznenihBodova" json:"godisnjiBrojKaznenihBodova"`
}
type StatistikaRegistrovanihVozila struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BrojRegistrovanihVozila    string             `bson:"brojRegistrovanihVozila" json:"brojRegistrovanihVozila"`
	ProsecniPeriodVazenja      string             `bson:"prosecniPeriodVazenja" json:"prosecniPeriodVazenja"`
	NajcesceRegistrovanoVozilo string             `bson:"najcesceRegistrovanoVozilo" json:"najcesceRegistrovanoVozilo"`
}
type StatistikaAuta struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UkupanBrojVozila          string             `bson:"ukupanBrojVozila" json:"ukupanBrojVozila"`
	NajcesciModel             string             `bson:"NajcesciModel" json:"NajcesciModel"`
	NajcescaGodinaProizvodnje string             `bson:"najcescaGodinaProizvodnje" json:"najcescaGodinaProizvodnje"`
}

type StatistikaNalogaZaPretres struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NajcesciRazlogPretresa     string             `bson:"najcesciRazlogPretresa" json:"najcesciRazlogPretresa"`
	NajcesciKorisnik           string             `bson:"najcesciKorisnik" json:"najcesciKorisnik"`
	BrojIzdatihNalogaZaPretres string             `bson:"brojIzdatihNalogaZaPretres" json:"brojIzdatihNalogaZaPretres"`
	NajcescaLokacija           string             `bson:"najcescaLokacija" json:"najcescaLokacija"`
}
type StatistikaSaslusanja struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NajcesciRazlogSaslusanja  string             `bson:"najcesciRazlogSaslusanja" json:"najcesciRazlogSaslusanja"`
	ProsecnoTrajanjeSalusanja string             `bson:"prosecnoTrajanjeSalusanja" json:"prosecnoTrajanjeSalusanja"`
	BrojGodisnjihSaslusanja   string             `bson:"brojGodisnjihSaslusanja" json:"brojGodisnjihSaslusanja"`
}
type StatistikaPravnogZahteva struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NajcesciRazlog       string             `bson:"najcesciRazlog" json:"najcesciRazlog"`
	NajcesciKorisnik     string             `bson:"najcesciKorisnik" json:"najcesciKorisnik"`
	BrojGodisnjihZahteva string             `bson:"brojGodisnjihZahteva" json:"brojGodisnjihZahteva"`
}
