package models

import (
	"github.com/Luncher/go-rest/db"
	"github.com/Luncher/go-rest/forms"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:name`
	Rating float32       `json:rating`
	Desc   string        `json:desc`
}

type MovieModel struct{}

var dbConnect = db.NewConnection("localhost")

func (m *MovieModel) Create(data forms.CreateMovieCommand) error {
	collection := dbConnect.Use("test-mgo", "movies")
	err := collection.Insert(bson.M{"name": data.Name, "rating": data.Rating, "desc": data.Desc})
	return err
}

func (m *MovieModel) Get(id string) (movie Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MovieModel) Find() (list []Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *MovieModel) Update(id string, data forms.UpdateMovieCommand) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return err
}

func (m *MovieModel) Delete(id string) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.RemoveId(id)

	return err
}
