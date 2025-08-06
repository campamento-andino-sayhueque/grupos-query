package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cas/grupos-query/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll(ctx context.Context) ([]domain.Grupo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Grupo), args.Error(1)
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (*domain.Grupo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Grupo), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, group *domain.Grupo) error {
	args := m.Called(ctx, group)
	return args.Error(0)
}

func (m *MockRepository) Update(ctx context.Context, group *domain.Grupo) error {
	args := m.Called(ctx, group)
	return args.Error(0)
}

func TestGetGrupos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(MockRepository)
	grupos := []domain.Grupo{{ID: "1", Nombre: "Grupo 1"}}
	mockRepo.On("GetAll", mock.Anything).Return(grupos, nil)

	h := NewHandlers(mockRepo)

	if assert.NoError(t, h.GetGrupos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"ID":"1","Nombre":"Grupo 1","Dirigentes":null,"CantidadAcampantes":0}]`, rec.Body.String())
	}
}

func TestGetGrupoByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := new(MockRepository)
	grupo := &domain.Grupo{ID: "1", Nombre: "Grupo 1"}
	mockRepo.On("GetByID", mock.Anything, "1").Return(grupo, nil)

	h := NewHandlers(mockRepo)

	if assert.NoError(t, h.GetGrupoByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"ID":"1","Nombre":"Grupo 1","Dirigentes":null,"CantidadAcampantes":0}`, rec.Body.String())
	}
}
