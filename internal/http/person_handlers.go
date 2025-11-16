package http

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/DmitriyChubarov/enkod/internal/app"
	"github.com/DmitriyChubarov/enkod/internal/logic"
	"github.com/labstack/echo/v4"
)

type PersonHandler struct {
	service *logic.PersonService
}

func NewPersonHandler(service *logic.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) Register(e *echo.Echo) {
	e.GET("/person", h.ListPersons)
	e.GET("/person/:id", h.GetPerson)
	e.POST("/person", h.CreatePerson)
	e.PUT("/person/:id", h.UpdatePerson)
	e.DELETE("/person/:id", h.DeletePerson)
}

func (h *PersonHandler) ListPersons(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 50
	}
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	search := c.QueryParam("search")
	
	log.WithFields(log.Fields{
		"method": c.Request().Method,
		"path":   c.Path(),
		"limit":  limit,
		"offset": offset,
		"search": search,
	}).Info("HTTP запрос: получение списка пользователей")

	people, err := h.service.ListPersons(c.Request().Context(), limit, offset, search)
	if err != nil {
		log.WithField("error", err.Error()).Error("HTTP ошибка при получении списка пользователей")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	log.WithField("count", len(people)).Info("HTTP ответ: список пользователей успешно отправлен")
	return c.JSON(http.StatusOK, people)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	
	log.WithFields(log.Fields{
		"method": c.Request().Method,
		"path":   c.Path(),
		"id":     id,
	}).Info("HTTP запрос: получение пользователя")

	p, err := h.service.GetPerson(c.Request().Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Warn("HTTP ошибка: пользователь не найден")
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	log.WithField("id", id).Info("HTTP ответ: пользователь успешно отправлен")
	return c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	var req app.Person
	if err := c.Bind(&req); err != nil {
		log.WithFields(log.Fields{
			"method": c.Request().Method,
			"path":   c.Path(),
			"error":  err.Error(),
		}).Warn("HTTP ошибка: невалидный JSON при создании пользователя")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	
	log.WithFields(log.Fields{
		"method":    c.Request().Method,
		"path":      c.Path(),
		"email":     req.Email,
		"firstName": req.FirstName,
		"lastName":  req.LastName,
	}).Info("HTTP запрос: создание пользователя")

	p, err := h.service.CreatePerson(c.Request().Context(), &req)
	if err != nil {
		log.WithFields(log.Fields{
			"email": req.Email,
			"error": err.Error(),
		}).Warn("HTTP ошибка при создании пользователя")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	log.WithFields(log.Fields{
		"id":    p.Id,
		"email": p.Email,
	}).Info("HTTP ответ: пользователь успешно создан")
	return c.JSON(http.StatusCreated, p)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req app.Person
	if err := c.Bind(&req); err != nil {
		log.WithFields(log.Fields{
			"method": c.Request().Method,
			"path":   c.Path(),
			"id":     id,
			"error":  err.Error(),
		}).Warn("HTTP ошибка: невалидный JSON при обновлении пользователя")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	
	log.WithFields(log.Fields{
		"method":    c.Request().Method,
		"path":      c.Path(),
		"id":        id,
		"email":     req.Email,
		"firstName": req.FirstName,
		"lastName":  req.LastName,
	}).Info("HTTP запрос: обновление пользователя")

	p, err := h.service.UpdatePerson(c.Request().Context(), id, &req)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Warn("HTTP ошибка при обновлении пользователя")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	log.WithField("id", id).Info("HTTP ответ: пользователь успешно обновлен")
	return c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	
	log.WithFields(log.Fields{
		"method": c.Request().Method,
		"path":   c.Path(),
		"id":     id,
	}).Info("HTTP запрос: удаление пользователя")

	err := h.service.DeletePerson(c.Request().Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Warn("HTTP ошибка при удалении пользователя")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	log.WithField("id", id).Info("HTTP ответ: пользователь успешно удален")
	return c.NoContent(http.StatusNoContent)
}
