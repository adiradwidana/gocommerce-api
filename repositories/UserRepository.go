package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"hello/models"
)

type UserRepository struct{}

const DBSERVER = "localhost:27017"
const DBNAME = "gocommerce"
const DOCNAME = "users"

func (userRepository UserRepository) GetUsers() models.Users{
	session, err := mgo.Dial(DBSERVER)

	if err != nil {
		fmt.Println("Failed to connect to database server")
	}

	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := models.Users{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results;
}
