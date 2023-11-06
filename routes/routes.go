package routes

import (
	"Project/config/db"
	routes "Project/routes/authR"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	db.ConnectDb()
	defer db.CloseDb()
	r := gin.Default()

	routes.AuthRoutes(r)
	return r
}
