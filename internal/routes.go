package internal

import "github.com/labstack/echo/v4"

func Assign(e *echo.Echo) {
	v1 := e.Group("/wow/api/v1")

	v1Health := v1.Group("/health")
	v1Health.GET("", GetHealth)

	v1Users := v1.Group("/users")
	v1Users.GET("/:id", GetUser)
	v1Users.GET("", GetUsers)
	v1Users.POST("", PostUser)
	v1Users.PUT("/:id", PutUser)
	v1Users.DELETE("/:id", DeleteUser)

	V1Chars := v1Users.Group("/:userId/chars")
	V1Chars.GET("", GetCharacters)
	V1Chars.POST("", PostCharacter)
	V1Chars.DELETE("/:characterId", DeleteCharacter)
}
