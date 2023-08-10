package app

import (
	"github.com/codescalersinternships/todo-diaa/internal/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App struct for handlers
type App struct {
	Router *gin.Engine
	db     *database.Database
}

// NewApp creates a new app
func NewApp(dbPath string) (*App, error) {

	db, err := StartDB(dbPath)
	if err != nil {
		return nil, err
	}

	return &App{Router: gin.Default(), db: db}, nil
}

// StartDB starts the database
func StartDB(path string) (*database.Database, error) {

	db, err := database.NewDB(path)

	if err != nil {
		return nil, err
	}

	return db, db.Migrate()
}

// CORSMiddleware is the middleware that handles CORS
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

// Run runs the app
func (a *App) Run(address string) error {

	a.SetAPIs()

	defer a.db.CloseDB()

	return a.Router.Run(address)
}

// SetAPIs sets the APIs
func (a *App) SetAPIs() {
	a.Router.Use(CORSMiddleware())

	a.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todoGroup := a.Router.Group("/todo")
	todoGroup.GET("", a.FindAll)

	todoGroup.POST("", a.AddItem)

	todoGroup.GET("/:id", a.GetById)

	todoGroup.PUT("", a.UpdateItem)

	todoGroup.DELETE("/:id", a.DeleteItem)
}
