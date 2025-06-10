package handlers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repositories/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

func GetExecsHandler(w http.ResponseWriter, r *http.Request) {
	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDbHandler(execs, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	exec, err := sqlconnect.GetExecByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

func AddExecsHandler(w http.ResponseWriter, r *http.Request) {
	var newExecs []models.Exec
	var rawExecs []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &rawExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	fields := GetFieldNames(models.Exec{})
	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, exec := range rawExecs {
		for key := range exec {
			if _, ok := allowedFields[key]; !ok {
				http.Error(w, "Invalid field: "+key, http.StatusBadRequest)
				return
			}
		}
	}

	err = json.Unmarshal(body, &newExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for _, exec := range newExecs {
		err := CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.AddExecsDBHandler(newExecs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}

	json.NewEncoder(w).Encode(response)
}

func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {
	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecs(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func PatchOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid exec ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedExec, err := sqlconnect.PatchOneExec(id, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExec)
}

func DeleteOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid exec ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec deleted successfully",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Exec
	// Data Validation

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Search for user if user actually exists
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "error updating data")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	user := &models.Exec{}

	err = db.QueryRow("SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?", req.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorHandler(err, "user not found")
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		http.Error(w, "database query error", http.StatusBadRequest)
		return
	}
	// is user active?
	if user.InactiveStatus {
		http.Error(w, "User is inactive", http.StatusForbidden)
		return
	}

	// verify password
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		utils.ErrorHandler(errors.New("invalid encoded hash format"), "invalid password format")
		http.Error(w, "Invalid encoded hash format", http.StatusInternalServerError)
		return
	}

	saltBase64 := parts[0]
	hashedPasswordBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)

	if err != nil {
		utils.ErrorHandler(err, "error decoding salt")
		http.Error(w, "Error decoding salt", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)

	if err != nil {
		utils.ErrorHandler(err, "failed to decode hashed password")
		http.Error(w, "Error decoding hashed password", http.StatusInternalServerError)
		return
	}

	hash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		// do nothing
	} else {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	tokenString, err := utils.SignToken(user.ID, req.Username, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	// Send token as a response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour), // Token valid for 24 hours
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "Testabc",
		Value:    "testing123",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour), // Token valid for 24 hours
	})

	// Response body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	json.NewEncoder(w).Encode(response)
}
