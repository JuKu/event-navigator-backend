package main

import (
	"github.com/JuKu/event-navigator-backend/db"
	"github.com/JuKu/event-navigator-backend/routes"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db.InitDB()

	// for production usage (alternative: export GIN_MODE=release)
	// gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")

	if err != nil {
		return
	}
}
