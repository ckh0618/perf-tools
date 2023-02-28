package multitx

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"han-mongodb/log"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"testing"
)

type Book struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}

func TestOrm(t *testing.T) {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	book := NewBook("Pride and Prejudice", 345)

	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err = mgm.Coll(book).Create(book)
	assert.NoError(t, err)

	book2 := &Book{}
	coll := mgm.Coll(book2)

	err = coll.FindOne(context.Background(), bson.M{"name": "Pride and Prejudice"}).Decode(book2)
	assert.NoError(t, err)

	log.Info("Data %v", *book2)

}
