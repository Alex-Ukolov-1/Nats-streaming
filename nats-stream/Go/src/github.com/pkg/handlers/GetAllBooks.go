package handlers
import (
	"net/http"
	"encoding/json"
	"github.com/pkg/mocks")

func GetAllBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}