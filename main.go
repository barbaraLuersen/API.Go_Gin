package main

import (
	"github.com/barbaraLuersen/GoGin/database"
	"github.com/barbaraLuersen/GoGin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
