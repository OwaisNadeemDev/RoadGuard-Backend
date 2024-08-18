package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/owaisnadeemdev/roadguard/internal/api/http/models"
	"github.com/owaisnadeemdev/roadguard/internal/api/http/util"
	"github.com/owaisnadeemdev/roadguard/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandle(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.SendJSONResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" || user.Phonenumber == "" {
		util.SendJSONResponse(w, http.StatusBadRequest, false, "username, email, password, phoneNumber are required", nil)
		return
	}

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 OR "phoneNumber"=$2)`
	err = config.DB.QueryRow(query, user.Email, user.Phonenumber).Scan(&exists)
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	if exists {
		util.SendJSONResponse(w, http.StatusBadRequest, false, "Email or Phonenumber is already associated with an Account", nil)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if hashedPassword != nil {
		user.Password = string(hashedPassword)
	}

	query = `INSERT INTO users (username, email, password, "phoneNumber") VALUES ($1, $2, $3, $4)`

	_, err = config.DB.Exec(query, user.Username, user.Email, user.Password, user.Phonenumber)
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, true, "User Registered Successfully", nil)
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.SendJSONResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	if user.Email == "" || user.Password == "" {
		util.SendJSONResponse(w, http.StatusBadRequest, false, "Email/Username and password are required", nil)
		return
	}

	var storedUser models.User
	query := `SELECT id, username, email, password, "phoneNumber" FROM users WHERE email=$1 OR username=$1`
	err := config.DB.QueryRow(query, user.Email).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Email, &storedUser.Password, &storedUser.Phonenumber)

	if err != nil {
		util.SendJSONResponse(w, http.StatusNotFound, false, "Invalid email/username or password", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		util.SendJSONResponse(w, http.StatusNotFound, false, "Invalid email/username or password", nil)
		return
	}

	storedUser.Password = ""

	util.SendJSONResponse(w, http.StatusOK, true, "Login successful", storedUser)
}
