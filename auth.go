package main

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", username, string(hashedPassword))
	if err != nil {
		fmt.Println("❌ Username already exists or DB error:", err)
		return
	}
	fmt.Println("✅ Registration successful!")
}

func Login(username, password string) {
	var hashed string
	err := DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashed)

	if err == sql.ErrNoRows {
		fmt.Println("❌ User not found!")
		return
	} else if err != nil {
		fmt.Println("❌ DB error:", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		fmt.Println("❌ Incorrect password!")
	} else {
		fmt.Printf("✅ Welcome, %s!\n", username)
	}
}

