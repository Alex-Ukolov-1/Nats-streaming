package handlers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/pkg/mocks"
)

func GetBook(w http.ResponseWriter,r *http.Request){
 vars := mux.Vars(r)
 id, _ := strconv.Atoi(vars["id"])
 for _, book := range mocks.Books {
	 if book.ID == id {
		 w.Header().Add("Content-Type", "application/json")
		 w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(book)
		 break
	   }
   }
}