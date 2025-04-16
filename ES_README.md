# Ejemplo básico: Go + MySQL

Este proyecto es un ejemplo sencillo de cómo conectar una aplicación Go con una base de datos MySQL y realizar operaciones CRUD sobre una tabla de contactos.

---

## Estructura del Proyecto

```
/your-project
│
├── controllers/
│   └── controllers.go   # Funciones CRUD para Contact
│
├── database/
│   └── connect.go       # Conexión a la base de datos
│
├── models/
│   └── contact.go       # Modelo Contact
│
├── .env                 # Variables de entorno para la base de datos
└── main.go              # Ejemplo de uso
```

---

## Configuración

1. **Variables de entorno**

Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido (ajusta los valores a tu entorno):

```env
DB_USERNAME=tu_usuario
DB_PASSWORD=tu_contraseña
DB_ENVIRONMENT=localhost
DB_PORT=3306
DB_NAME=nombre_base_de_datos
```

2. **Instala las dependencias**
```bash
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
```

---

## Modelo de Contacto

Archivo: `models/contact.go`

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

## Conexión a MySQL

Archivo: `database/connect.go`

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

## Controlador de Contactos

Archivo: `controllers/controllers.go`

Incluye funciones para listar, obtener, crear, actualizar y eliminar contactos. Ejemplo de uso en el main.

---

## Ejemplo de Uso en `main.go`

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

	// Listar todos los contactos
	contacts, err := controllers.ListContacts(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range contacts {
		log.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s", c.Id, c.Name, c.Email, c.Phone)
	}

	// Crear un contacto con email aleatorio
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	email := fmt.Sprintf("newcontact%d@gmail.com", randomNumber)
	newContact := models.Contact{
		Name:  "Nuevo Contacto",
		Email: email,
		Phone: "123-4567",
	}
	id, err := controllers.CreateContact(db, newContact)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contacto creado con ID: %d", id)

	// Actualizar un contacto
	update := models.Contact{
		Id:    id,
		Name:  "Contacto Actualizado",
		Email: "actualizado@gmail.com",
		Phone: "888-9999",
	}
	_, err = controllers.UpdateContact(db, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto actualizado.")

	// Eliminar el contacto
	_, err = controllers.DeleteContact(db, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto eliminado.")
}
```

---

## Notas

- Asegúrate de tener la tabla `contact` creada en tu base de datos MySQL:

```sql
CREATE TABLE contact (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) DEFAULT NULL,
  phone VARCHAR(20) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

- Cambia `"your-module-path"` por el path real de tu módulo Go.

---

## Resumen

- Configura el archivo `.env`
- Ejecuta `go run main.go`
- El proyecto realiza operaciones CRUD básicas sobre la tabla `contact` en MySQL usando Go.

---

¿Necesitas que agregue algo más?

