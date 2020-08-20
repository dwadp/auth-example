package registry

import (
	"context"

	"github.com/dwadp/auth-example/app/middlewares"
	authController "github.com/dwadp/auth-example/auth/controller/rest"
	authRepository "github.com/dwadp/auth-example/auth/repository"
	authService "github.com/dwadp/auth-example/auth/service"
	"github.com/dwadp/auth-example/pkg/hash/bcrypt"
	"github.com/dwadp/auth-example/pkg/jwt"
	userController "github.com/dwadp/auth-example/user/controller/rest"
	userRepository "github.com/dwadp/auth-example/user/repository"
	userService "github.com/dwadp/auth-example/user/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(app *gin.Engine, mongoDB *mongo.Database, redisDB *redis.Client, ctx context.Context) {
	// Initialize all packages
	hash := bcrypt.NewBcrypt()
	jwtAuth := jwt.NewJWTAuth()

	// Initialize all repositories
	userRepo := userRepository.NewMongoUserRepository(mongoDB, ctx)
	authRepo := authRepository.NewAuthRedisRepository(redisDB, ctx)

	// Initialize all services
	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo, authRepo, hash, jwtAuth)

	// Initialize all middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtAuth, authRepo)

	// Registers all controller and map all of their routes
	userController.New(app, userService, authMiddleware).MapRoutes()
	authController.New(app, authService, userService, authMiddleware).MapRoutes()
}
