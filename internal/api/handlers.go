package api

import (
	"net/http"

	"github.com/cas/grupos-query/internal/domain"
	"github.com/cas/grupos-query/internal/storage"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	repo storage.GroupRepository
}

func NewHandlers(repo storage.GroupRepository) *Handlers {
	return &Handlers{repo: repo}
}

// GetGrupos returns a list of all groups.
// @Summary Get all groups
// @Description Get all groups
// @Tags grupos
// @Produce json
// @Success 200 {array} domain.Grupo
// @Router /v1/grupos [get]
func (h *Handlers) GetGrupos(c echo.Context) error {
	grupos, err := h.repo.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, grupos)
}

// GetGrupoByID returns a group by its ID.
// @Summary Get a group by ID
// @Description Get a group by ID
// @Tags grupos
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} domain.Grupo
// @Router /v1/grupos/{id} [get]
func (h *Handlers) GetGrupoByID(c echo.Context) error {
	id := c.Param("id")
	grupo, err := h.repo.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, grupo)
}
