package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(w, "Invalid request body")
		return
	}
	existingUser, err := GetUserByName(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	if u.Password == existingUser.Password {
		tokenString, err := CreateToken(time.Hour*24, "your_secret_key_here")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error creating token")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Invalid Credentials")
}
