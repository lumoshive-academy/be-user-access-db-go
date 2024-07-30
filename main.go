package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// User struct to map to the users table
type User struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	DB_USER     = "adminlevel1"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

func main() {
	// Connect to the database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	// Create a new user
	newUser := User{Name: "Candil", Email: "candil@example.com", Phone: "08558358358"}
	userID, err := createUser(db, newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New user created with ID: %d\n", userID)

	// // Get a user by ID
	// user, err := getUserByID(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User fetched: %+v\n", user)

	// // Update a user
	// user.Name = "John Doe Jr."
	// err = updateUser(db, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User updated: %+v\n", user)

	// // Delete a user
	// err = deleteUser(db, user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User with ID %d deleted\n", user.ID)
}

// Create a new user
func createUser(db *sql.DB, user User) (int, error) {
	var userID int
	query := `
		INSERT INTO users (name, email, phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := db.QueryRow(query, user.Name, user.Email, user.Phone, time.Now(), time.Now()).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// // Get a user by ID
// func getUserByID(db *sql.DB, id int) (*User, error) {
// 	user := User{}
// 	query := "SELECT id, name, email, phone, created_at, updated_at, deleted_at FROM users WHERE id = $1"
// 	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// // Update a user
// func updateUser(db *sql.DB, user User) error {
// 	query := `
// 		UPDATE users
// 		SET name = $1, email = $2, phone = $3, updated_at = $4
// 		WHERE id = $5`
// 	_, err := db.Exec(query, user.Name, user.Email, user.Phone, time.Now(), user.ID)
// 	return err
// }

// // Delete a user
// func deleteUser(db *sql.DB, id int) error {
// 	query := "DELETE FROM users WHERE id = $1"
// 	_, err := db.Exec(query, id)
// 	return err
// }
