package route

import (
	"net/http"
	"strconv"
	"time"

	db "todo/db"
	models "todo/models"

	"github.com/labstack/echo"
)

var seq = 1

// Init echo
func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/todos", getTodos)
	e.POST("/todos", createTodos)
	e.PUT("/todos/:id", updateTodos)
	return e
}

func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Todos)
}

func createTodos(c echo.Context) error {
	t := &models.Todo{
		ID:          seq,
		DateCreated: time.Now(),
	}

	if err := c.Bind(t); err != nil {
		return err
	}

	db.AddTodo(t)
	seq++
	return c.JSON(http.StatusCreated, t)
}

func updateTodos(c echo.Context) error {
	t := new(models.Todo)

	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	t.ID = id
	if _, err := db.UpdateTodo(id, t); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, t)
}
