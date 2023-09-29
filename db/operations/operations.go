package operations

import (
	"fmt"
	"go-gqlgen/db/connection/dbconn"
	"log"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func CreateUser(u User) {
	q := `INSERT INTO users (name, email) VALUES ($1, $2)`
	db := dbconn.GetDB()

	_, err := db.Exec(q, u.Name, u.Email)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Println("User created successfully!")
}

func FetchAllUsers() {
	var users []User

	query := "SELECT id, name, email, updated_at, created_at FROM users"
	rows, err := dbconn.GetDB().Query(query)
	if err != nil {
		log.Fatalf("Error executing the query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.UpdatedAt, &u.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		users = append(users, u)
		fmt.Printf("ID: %s, Name: %s, Email: %s, UpdatedAt: %v, CreatedAt: %v\n", u.Id, u.Name, u.Email, u.UpdatedAt, u.CreatedAt)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
	}
}

func UpdateUserEmail(id string, newEmail string) {
	q := `UPDATE users SET email = $1 WHERE id = $2`
	db := dbconn.GetDB()

	_, err := db.Exec(q, newEmail, id)
	if err != nil {
		log.Fatalf("Error updating user email: %v", err)
	}

	fmt.Println("User email updated successfully!")
}

func DeleteUser(id string) {
	q := `DELETE FROM users WHERE id = $1`
	db := dbconn.GetDB()

	_, err := db.Exec(q, id)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}

	fmt.Println("User deleted successfully!")
}
