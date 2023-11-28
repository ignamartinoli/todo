package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func getHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func main() {
	connStr := "postgresql://postgres:password@localhost:5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return getHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
