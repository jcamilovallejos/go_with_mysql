package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_with_mysql/controllers"
	"go_with_mysql/database"
	"go_with_mysql/models"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}()

	// Listar todos los contactos
	contacts, err := controllers.ListContacts(db)
	if err != nil {
		log.Fatalf("Error listing contacts: %v", err)
	}
	for _, c := range contacts {
		log.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s", c.Id, c.Name, c.Email, c.Phone)
	}

	// Obtener un contacto por ID
	contact, err := controllers.GetContactByID(db, 1)
	if err != nil {
		log.Fatalf("Error getting contact by ID: %v", err)
	}
	log.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s", contact.Id, contact.Name, contact.Email, contact.Phone)

	// Crear un nuevo contacto con email aleatorio
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	email := fmt.Sprintf("%s%d@gmail.com", "newcontact", randomNumber)
	newContact := models.Contact{
		Name:  "New Random Contact",
		Email: email,
		Phone: "123-4567",
	}
	newContactID, err := controllers.CreateContact(db, newContact)
	if err != nil {
		log.Fatalf("Error creating contact: %v", err)
	}
	log.Printf("Created contact with ID: %d", newContactID)

	// Actualizar un contacto existente
	updatedContact := models.Contact{
		Id:    newContactID,
		Name:  "Updated Name",
		Email: "updated_email@gmail.com",
		Phone: "888-9999",
	}
	updated, err := controllers.UpdateContact(db, updatedContact)
	if err != nil {
		log.Fatalf("Error updating contact: %v", err)
	}
	log.Printf("Updated contact: ID: %d, Nombre: %s, Email: %s, Teléfono: %s", updated.Id, updated.Name, updated.Email, updated.Phone)

	// Eliminar un contacto
	deletedID, err := controllers.DeleteContact(db, newContactID)
	if err != nil {
		log.Fatalf("Error deleting contact: %v", err)
	}
	log.Printf("Deleted contact with ID: %d", deletedID)
}
