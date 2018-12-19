package controllers

import (
	"encoding/json"
	"hello/repositories"
	"log"
	"net/http"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func (UserController *UserController) Index(w http.ResponseWriter, r *http.Request){
	users := UserController.UserRepository.GetUsers()

	log.Println(users)
	data, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
