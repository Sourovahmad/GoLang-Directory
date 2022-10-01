package main

import (
	"gin-book-management/database"
	"gin-book-management/routes"
)

func main() {
	database.DatabaseMigration()
	routes.HandleRoutes()
}
