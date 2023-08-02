package config

import (
	"NOW/logic/Entities"
	"encoding/json"
	"fmt"
	"net/http"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

func VerifyNameMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into the user struct
		var body Entities.User
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		next(w, r)
	}
}
