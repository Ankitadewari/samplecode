package login

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid req body")
		return
	}

	_, err = GetUserByName(u.Username)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "User exists")
		return
	}

	err = CreateUser(u.Username, u.Password, u.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error creating user")
		return

	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created")

}
