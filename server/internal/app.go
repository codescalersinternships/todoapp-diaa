package App

import (
	"github.com/codescalersinternships/todo-diaa/internal/database"
	"github.com/gin-gonic/gin"
)

type app struct {
	Router *gin.Engine
	db     *database.Database
}

func NewApp() *app {

	return &app{Router: gin.Default(), db: nil}
}

func (a *app) StartDB(path string) error {

	db, err := database.NewDB(path)

	if err != nil {
		return err
	}

	a.db = db
	// migration
	err = db.Migrate()
	if err != nil {
		return err
	}
	return nil
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func (a *app) Run(address string) error {

	a.SetAPIs()
	
	defer a.db.CloseDB()

	return a.Router.Run(address)
}

func (a *app) SetAPIs() {
	a.Router.Use(CORSMiddleware())

	a.Router.GET("/todo", a.FindAll)

	a.Router.POST("/todo", a.AddItem)

	a.Router.GET("/todo/:id", a.GetById)

	a.Router.PUT("/todo", a.UpdateItem)

	a.Router.DELETE("/todo/:id", a.DeleteItem)
}
