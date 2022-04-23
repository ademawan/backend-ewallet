package routes

import (
	"backend-ewallet/delivery/controllers/auth"
	"backend-ewallet/delivery/controllers/transaction"
	"backend-ewallet/delivery/controllers/user"

	"backend-ewallet/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	uc *user.UserController,
	ac *auth.AuthController,
	tc *transaction.TransactionController,

) {
	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", ac.Login())
	e.POST("users/logout", ac.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())
	e.GET("/users/me/search", uc.Search())
	// e.GET("/users/me/dummy", uc.Dummy())
	//df

	//TRANSACTION CONTROLLER
	e.POST("users/me/transactions", tc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/transactions", tc.Get(), middlewares.JwtMiddleware())
	e.GET("/users/me/transactions/:transaction_id", tc.GetByID(), middlewares.JwtMiddleware())
	e.GET("/users/me/transactions/received", tc.GetTransactionReceived(), middlewares.JwtMiddleware())
	e.GET("/users/me/transactions/send", tc.GetTransactionSend(), middlewares.JwtMiddleware())
}
