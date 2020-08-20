package rest

import (
	"net/http"

	"github.com/dwadp/auth-example/app/middlewares"
	"github.com/dwadp/auth-example/response"
	"github.com/dwadp/auth-example/user"
	"github.com/dwadp/mantau"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	app            *gin.Engine
	userService    user.Service
	authMiddleware *middlewares.AuthMiddleware
}

func New(
	app *gin.Engine,
	userService user.Service,
	authMiddleware *middlewares.AuthMiddleware) *UserController {
	return &UserController{
		app:            app,
		userService:    userService,
		authMiddleware: authMiddleware,
	}
}

func (u *UserController) MapRoutes() {
	router := u.app.Group("/api/v1/users")

	router.Use(u.authMiddleware.Handle())

	router.GET("/", u.all)
}

func (u *UserController) all(ctx *gin.Context) {
	result, err := u.userService.GetAll()
	m := mantau.New()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(
			"Something went wrong",
			err,
		))
		return
	}

	transformedResult, err := m.Transform(result, mantau.Schema{
		"id":    mantau.SchemaField{Key: "id"},
		"name":  mantau.SchemaField{Key: "name"},
		"email": mantau.SchemaField{Key: "email"},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(
			"Something went wrong",
			err,
		))
	}

	ctx.JSON(http.StatusOK, response.OK(
		"",
		transformedResult,
	))
}
