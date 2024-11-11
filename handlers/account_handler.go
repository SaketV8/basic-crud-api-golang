package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saketV8/basic-crud-api-golang/models"
	"github.com/saketV8/basic-crud-api-golang/sqlite"
)

func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is Homepage",
	})
}

func TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok-tested 👍",
	})
}

func GetAllUsers(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	//calling <All> method of <UserAccDb_Model>
	users, err := UserAccDb_Model.All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get user accounts",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	// extracting the <user_name> from url
	userNameParam := ctx.Param("user_name")

	//calling <GetByUserName> method of <UserAccDb_Model>
	userAcc, err := UserAccDb_Model.GetByUserName(userNameParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get user account",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, userAcc)
}

func CreateUserAccountTest(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	var rowAffected int64

	// creating data to insert :)
	TestUser := models.TestUserStruct{
		UserName:    "mark",
		FirstName:   "mark",
		LastName:    "manson",
		PhoneNumber: "888-888-909",
		Email:       "markmanson@example.com",
	}

	//calling <Insert> method of <UserAccDb_Model>
	rowAffected, err := UserAccDb_Model.Insert(TestUser.UserName, TestUser.FirstName, TestUser.LastName, TestUser.PhoneNumber, TestUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to insert <Test> user account",
			"details": err.Error(),
			"body":    TestUser,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User account successfully created",
		"row-affected": rowAffected,
		"body":         TestUser,
	})
}

func CreateUserAccount(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	// incoming request data (body of post request)
	var body models.CreateUserRequestBody
	var rowAffected int64

	// Bind the JSON request body to the struct
	// err := ctx.BindJSON(&body)
	err := ctx.Bind(&body)
	if err != nil {
		// If binding fails, return a 400 error with the error message
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})
		return
	}

	rowAffected, err = UserAccDb_Model.Insert(body.UserName, body.FirstName, body.LastName, body.PhoneNumber, body.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to insert user account",
			"details": err.Error(),
			"body":    body,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User account successfully created",
		"row-affected": rowAffected,
		"body":         body,
	})
}

func UpdateUserAccount(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	// incoming request data (body of post request)
	var body models.UpdateUserRequestBody
	var rowAffected int64

	// Bind the JSON to the <RequestBody> struct
	// err := ctx.BindJSON(&body)
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})
		return
	}

	// Insert the user data into the database
	rowAffected, err = UserAccDb_Model.Update(body.UserName, body.FirstName, body.LastName, body.PhoneNumber, body.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update user account",
			"details": err.Error(),
			"body":    body,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User account successfully updated",
		"row-affected": rowAffected,
		"body":         body,
	})
}

func DeleteUserAccount(ctx *gin.Context, UserAccDb_Model *sqlite.UserAccountsDbModel) {
	// incoming request data (body of post request)
	var body models.DeleteRequestBody
	var rowAffected int64

	// Bind the JSON to the <DeleteRequestBody> struct
	// err := ctx.BindJSON(&body)
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})
		return
	}
	rowAffected, err = UserAccDb_Model.Delete(body.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete user account",
			"details": err.Error(),
			"body":    body,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User account successfully delete",
		"row-affected": rowAffected,
		"body":         body,
	})
}
