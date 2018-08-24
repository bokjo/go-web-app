package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

const passwordSalt = "a99VV0Wzmd1C9ujcitZ0fIVNE0I5I61AC47C852RoLTsHDyLCltvP+ZHEkIl/2hkzTOW90c3ZEjtY-kdfTWJ1Q+="

// User struct
type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

// Login function: handles the loging process
func Login(email, password string) (*User, error) {

	user := &User{}

	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))

	pass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// Top-notch security
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname, lastlogin
		  FROM public.user
		 WHERE email = $1
		   AND password = $2`, email, pass)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	currentTime := time.Now()

	_, err = db.Exec(`
		UPDATE public.user
		   SET lastlogin = $1
		 WHERE id = $2`, currentTime, user.ID)

	if err != nil {
		log.Printf("Failed to update the login time for user %v to %v\n [ ERROR ]: %v", user.Email, currentTime, err)
	}

	return user, nil
}
