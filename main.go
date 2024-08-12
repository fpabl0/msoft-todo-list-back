package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/msoft-g1/todo-list-backend/internal/domain/task"
	"github.com/msoft-g1/todo-list-backend/internal/domain/user"
	"github.com/msoft-g1/todo-list-backend/internal/graph/generated"
	"github.com/msoft-g1/todo-list-backend/internal/graph/resolver"
	"github.com/msoft-g1/todo-list-backend/internal/repository/sqliterepo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// -- Initialize database
	db, err := gorm.Open(sqlite.Open("./data/todo-list.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// -- Initialize repositories
	usersRepo := sqliterepo.NewUsersRepo(db)
	tasksRepo := sqliterepo.NewTasksRepo(db)

	// -- Initialize services
	usersService := user.NewService(usersRepo)
	tasksService := task.NewService(tasksRepo)

	// -- Define graph related handler
	playHandler := playground.Handler("GraphQL playground", "/graphql")
	graphHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			UsersService: usersService,
			TasksService: tasksService,
		},
	}))

	// -- Setting gin router
	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowMethods:    []string{"GET", "POST"},
			AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowAllOrigins: true,
		}),
		user.AuthMiddleware(usersService),
	)

	r.GET("/play", func(c *gin.Context) {
		playHandler.ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/graphql", func(c *gin.Context) {
		graphHandler.ServeHTTP(c.Writer, c.Request)
	})

	// -- Run the server
	r.Run(fmt.Sprintf(":%s", port))
}
