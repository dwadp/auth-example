package rest

import (
	"errors"
	"net/http"

	"github.com/dwadp/auth-example/app/middlewares"
	"github.com/dwadp/auth-example/auth"
	"github.com/dwadp/auth-example/models"
	"github.com/dwadp/auth-example/pkg/validation"
	"github.com/dwadp/auth-example/response"
	"github.com/dwadp/auth-example/user"
	"github.com/dwadp/mantau"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type AuthController struct {
	app            *gin.Engine
	authService    auth.Service
	userService    user.Service
	authMiddleware *middlewares.AuthMiddleware
}

func New(
	app *gin.Engine,
	authService auth.Service,
	userService user.Service,
	authMiddleware *middlewares.AuthMiddleware) *AuthController {
	return &AuthController{
		app:            app,
		authService:    authService,
		userService:    userService,
		authMiddleware: authMiddleware,
	}
}

func (u *AuthController) MapRoutes() {
	router := u.app.Group("/api/v1/auth")

	router.POST("/login", u.login)
	router.POST("/register", u.register)
	router.GET("/me", u.authMiddleware.Handle(), u.me)
	router.GET("/logout", u.authMiddleware.Handle(), u.logout)
}

func (u *AuthController) login(ctx *gin.Context) {
	var userLogin models.UserLogin

	if err := ctx.ShouldBind(&userLogin); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)

		if !ok {
			ctx.JSON(http.StatusBadRequest, response.Error("Bad request", err))
			return
		}

		formattedErrs, err := validation.Format(validationErr, userLogin)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		}

		ctx.JSON(http.StatusUnprocessableEntity, response.WithStatus("The given data was invalid", false, formattedErrs))
		return
	}

	authExpiry := viper.GetInt("auth.expiry")
	authSecret := viper.GetString("auth.secret")

	auth, err := u.authService.Login(userLogin, authSecret, authExpiry)

	if err != nil {
		if err == user.NotFound || err == user.WrongPassword {
			ctx.JSON(http.StatusUnauthorized, response.Error("Unauthenticated", err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		return
	}

	m := mantau.New()
	transformedResult, err := m.Transform(auth,
		mantau.Schema{
			"token": mantau.SchemaField{Key: "token"},
			"user": mantau.SchemaField{
				Key: "user",
				Value: mantau.Schema{
					"name":  mantau.SchemaField{Key: "name"},
					"email": mantau.SchemaField{Key: "email"},
				},
			},
		},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		return
	}

	ctx.JSON(http.StatusOK, response.OK("", transformedResult))
}

func (u *AuthController) register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)

		if !ok {
			ctx.JSON(http.StatusBadRequest, response.Error("Bad request", err))
			return
		}

		formattedErrs, err := validation.Format(validationErr, user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		}

		ctx.JSON(http.StatusUnprocessableEntity, response.WithStatus("The given data was invalid", false, formattedErrs))
		return
	}

	if err := u.authService.Register(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		return
	}

	ctx.JSON(http.StatusCreated, response.OK("User has been registered successfully", nil))
}

func (a *AuthController) me(ctx *gin.Context) {
	xUserID, exists := ctx.Get("X-USER-ID")

	if !exists {
		ctx.JSON(http.StatusNotFound, response.Error("Not found", errors.New("user id cannot be found")))
		return
	}

	result, err := a.userService.GetByID(xUserID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		return
	}

	m := mantau.New()
	transformedResult, err := m.Transform(result, mantau.Schema{
		"name":  mantau.SchemaField{Key: "name"},
		"email": mantau.SchemaField{Key: "email"},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Internal server error", err))
		return
	}

	ctx.JSON(http.StatusOK, response.OK("", transformedResult))
}

func (a *AuthController) logout(ctx *gin.Context) {
	xSessionID, exists := ctx.Get("X-SESSION-ID")

	if !exists {
		ctx.JSON(
			http.StatusNotFound,
			response.Error("Not found", errors.New("user id cannot be found")),
		)
		return
	}

	if err := a.authService.Logout(xSessionID.(string)); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.Error("Internal server error", err),
		)
		return
	}

	ctx.JSON(http.StatusOK, response.OK("Successfully logged you out", nil))
}
