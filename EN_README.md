# Simple Example: Go + MySQL

This project is a basic example showing how to connect a Go application to a MySQL database and perform CRUD operations on a contacts table.

---

## Project Structure

```
/your-project
│
├── controllers/
│   └── controllers.go   # CRUD functions for Contact
│
├── database/
│   └── connect.go       # Database connection logic
│
├── models/
│   └── contact.go       # Contact model
│
├── .env                 # Environment variables for database config
└── main.go              # Usage example
```

---

## Configuration

1. **Environment Variables**

Create a `.env` file in the project root with the following content (adjust values for your environment):

```env
DB_USERNAME=your_user
DB_PASSWORD=your_password
DB_ENVIRONMENT=localhost
DB_PORT=3306
DB_NAME=your_database_name
```

2. **Install Dependencies**
```bash
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
```

---

## Contact Model

File: `models/contact.go`

```go
package models

type Contact struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
```

---

## MySQL Connection

File: `database/connect.go`

```go
package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ENVIRONMENT"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to database")
	return db, nil
}
```

---

## Contacts Controller

File: `controllers/controllers.go`

Includes functions to list, get, create, update, and delete contacts. See usage in `main.go`.

---

## Usage Example (`main.go`)

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"your-module-path/controllers"
	"your-module-path/database"
	"your-module-path/models"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// List all contacts
	contacts, err := controllers.ListContacts(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range contacts {
		log.Printf("ID: %d, Name: %s, Email: %s, Phone: %s", c.Id, c.Name, c.Email, c.Phone)
	}

	// Create a new contact with random email
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	email := fmt.Sprintf("newcontact%d@gmail.com", randomNumber)
	newContact := models.Contact{
		Name:  "New Contact",
		Email: email,
		Phone: "123-4567",
	}
	id, err := controllers.CreateContact(db, newContact)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created contact with ID: %d", id)

	// Update a contact
	update := models.Contact{
		Id:    id,
		Name:  "Updated Contact",
		Email: "updated@gmail.com",
		Phone: "888-9999",
	}
	_, err = controllers.UpdateContact(db, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact updated.")

	// Delete the contact
	_, err = controllers.DeleteContact(db, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact deleted.")
}
```

---

## Notes

- Make sure your MySQL database has the `contact` table:

```sql
CREATE TABLE contact (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) DEFAULT NULL,
  phone VARCHAR(20) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

- Replace `"your-module-path"` with your actual Go module path.

---

## Summary

- Configure your `.env` file.
- Run `go run main.go`.
- The project performs basic CRUD operations on the `contact` table in MySQL using Go.

---

Let me know if you want to add anything else!

