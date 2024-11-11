package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saketV8/basic-crud-api-golang/handlers"
	"github.com/saketV8/basic-crud-api-golang/sqlite"
)

func SetupRouter(UserAccDb_Model *sqlite.UserAccountsDbModel) *gin.Engine {
	rtr := gin.Default()

	rtr.GET("/api/", handlers.HomePage)

	rtr.GET("/api/test-api", handlers.TestApi)

	rtr.GET("/api/users", func(ctx *gin.Context) {
		handlers.GetAllUsers(ctx, UserAccDb_Model)
	})

	rtr.GET("/api/user/:user_name", func(ctx *gin.Context) {
		handlers.GetUser(ctx, UserAccDb_Model)
	})

	rtr.POST("/api/create-user-account-test", func(ctx *gin.Context) {
		handlers.CreateUserAccountTest(ctx, UserAccDb_Model)
	})

	rtr.POST("/api/create-user-account", func(ctx *gin.Context) {
		handlers.CreateUserAccount(ctx, UserAccDb_Model)
	})

	rtr.DELETE("/api/delete-user-account", func(ctx *gin.Context) {
		handlers.DeleteUserAccount(ctx, UserAccDb_Model)
	})

	rtr.PUT("/api/update-user-account", func(ctx *gin.Context) {
		handlers.UpdateUserAccount(ctx, UserAccDb_Model)
	})

	// returning so that we can run the method .Run from main.go
	return rtr
}
