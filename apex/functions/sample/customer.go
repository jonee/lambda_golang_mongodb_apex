/**
 * Customer model class
 */

package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"log"
	"time"
)

// Customer class or struct definition
type Customer struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

	Username string `json:"username,omitempty" bson:"username,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Customer class Save function
func (r *Customer) Save(mapContext map[string]interface{}) (*Customer, error) {
	customerCol := mapContext["customerCol"].(*mgo.Collection)
	myCol := customerCol

	val := r.ExportArrayPrivate()

	var err error
	if !r.Id.Valid() {
		r.Id = bson.NewObjectId()
		val["_id"] = r.Id

		err = myCol.Insert(&val)

	} else {
		_, err = myCol.Upsert(bson.M{"_id": r.Id}, &val)
	}

	return r, err

	// Save
}

// ensureIndex function
func CustomerEnsureIndex(mapContext map[string]interface{}) error {
	customerCol := mapContext["customerCol"].(*mgo.Collection)

	keyIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := customerCol.EnsureIndex(keyIndex)
	if err != nil {
		log.Println("ensure key index {username} failed for Customer: " + err.Error())
	}

/*
	keyIndex2 := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err2 := customerCol.EnsureIndex(keyIndex2)
	if err2 != nil {
		log.Println("ensure key index {email} failed for Customer: " + err2.Error())
	}

	// return the (first) non nil
	if err == nil && err2 != nil {
		err = err2
	}
*/

	return err

	// CustomerEnsureIndex
}

// exportarrayprivate is saveable
func (r *Customer) ExportArrayPrivate() map[string]interface{} {
	val := make(map[string]interface{})

	if r.Id.Valid() {
		val["_id"] = r.Id
	}

	if r.Username != "" {
		val["username"] = r.Username
	}

	if !r.CreatedAt.IsZero() {
		val["created_at"] = r.CreatedAt
	}

	if !r.UpdatedAt.IsZero() {
		val["updated_at"] = r.UpdatedAt
	}

	return val

	// ExportArrayPrivate
}


