package api

import (
	"github.com/cas/grupos-query/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(repo storage.GroupRepository) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := NewHandlers(repo)

	v1 := e.Group("/v1")
	{
		grupos := v1.Group("/grupos")
		grupos.GET("", h.GetGrupos)
		grupos.GET("/:id", h.GetGrupoByID)
	}

	return e
}
