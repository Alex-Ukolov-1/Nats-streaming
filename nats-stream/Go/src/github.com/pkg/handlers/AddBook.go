package handlers

import (
	"io/ioutil"
	"net/http"
	"log"
	"math/rand"
	"encoding/json"
	"github.com/pkg/mocks"
	"github.com/pkg/models"
)

func AddBooks(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
	body,err:=ioutil.ReadAll(r.Body)

	if err!=nil{
      log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	book.ID=rand.Intn(100)
	mocks.Books=append(mocks.Books,book)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode("created")
}