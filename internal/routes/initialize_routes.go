package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xapier14/todo/internal/middlewares"
	"github.com/xapier14/todo/internal/responses/general"
	"github.com/xapier14/todo/internal/routes/auth"
	"github.com/xapier14/todo/internal/routes/todos"
)

var notAuthorizedResponse = general.GenerateRequestNotAuthorizedResponse()

func setupTodos(router *gin.RouterGroup) {
	todosRouter := router.Group("/todos")
	todosRouter.Use(middlewares.RequireAuth(notAuthorizedResponse))
	todosRouter.GET("/:id", todos.GetTodo)
	todosRouter.GET("/", todos.GetTodos)
	todosRouter.POST("/", todos.PostTodo)
	todosRouter.PATCH("/:id", todos.PatchTodo)
	todosRouter.DELETE("/:id", todos.DeleteTodo)
}

func setupAuth(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	authRouter.POST("/login", auth.PostLogin)
	authRouter.POST("/refresh", middlewares.RequireAuth(notAuthorizedResponse), auth.PostRefresh)
	authRouter.POST("/signup", auth.PostSignup)
}

func InitializeRoutes(router *gin.Engine, basePath string) {
	var baseRouter = router.Group("/")
	if basePath != "" {
		baseRouter = router.Group(basePath)
	}

	// setup group routes
	setupTodos(baseRouter)
	setupAuth(baseRouter)

	// setup static
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/robots.txt", "./static/robots.txt")
	router.StaticFile("/", "./static/index.html")
}