package api

import (
	"github.com/cas/grupos-query/internal/storage"
	"github.com/gin-gonic/gin"
)

func NewRouter(repo storage.GroupRepository) *gin.Engine {
	r := gin.Default()

	h := NewHandlers(repo)

	v1 := r.Group("/v1")
	{
		grupos := v1.Group("/grupos")
		grupos.GET("", h.GetGrupos)
		grupos.GET("/:id", h.GetGrupoByID)
	}

	return r
}
