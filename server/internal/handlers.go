package app

import (
	"net/http"
	"strconv"

	"github.com/codescalersinternships/todo-diaa/internal/database"
	"github.com/gin-gonic/gin"
)

// @Summary      Add new item
// @Description  Takes item name in JSON and store in DB. Return saved JSON.
// @Produce      json
// @Tags         todo
// @Param        item  body      database.TodoItem  true  "Item JSON"
// @Success      201   {object}  database.TodoItem
// @Router       / [post]
// @Failure      400   "Invalid request body"
// @Failure      500   "Internal server error"
func (a *App) AddItem(ctx *gin.Context) {

	var item database.TodoItem
	err := ctx.BindJSON(&item)

	if err != nil || item.Name == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	id, err := a.db.Insert(item)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	newStoredTodo, err := a.db.FindById(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, newStoredTodo)
}

// @Summary      Get todo items array
// @Description  Responds with the list of all todo items as JSON.
// @Tags         todo
// @Produce      json
// @Router       / [get]
// @Success      200   {object}  []database.TodoItem
// @Failure      500   "Internal server error"
func (a *App) FindAll(ctx *gin.Context) {

	todoList, err := a.db.FindAll()

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, todoList)
}

// @Summary      Update item
// @Description  Takes item in JSON and update it.
// @Tags	 	 todo
// @Produce      json
// @Param        item  body      database.TodoItem true "Item JSON"
// @Success      201  {object}  string
// @Router       / [put]
// @Failure      400  "Invalid request body"
// @Failure      500  "Internal server error"
func (a *App) UpdateItem(ctx *gin.Context) {
	var item database.TodoItem
	err := ctx.BindJSON(&item)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = a.db.Update(item)

	if err == database.ErrIDNotFound {

		ctx.Status(http.StatusBadRequest)
		return
	}

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, "updated successfully")
}

// @Summary      Delete item
// @Description  Takes item id and delete it.
// @Tags	 	 todo
// @Produce      json
// @Param        id  path      string  true  "search item by id"
// @Success      200  {object}  string
// @Router       /:id [delete]
// @Failure      404  "Item not found"
// @Failure      500  "Internal server error"
func (a *App) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.db.Delete(id)

	if err == database.ErrIDNotFound {
		ctx.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "deleted successfully")

}

// @Summary      Get single item by id
// @Description  Returns the item whose ID value matches the id.
// @Tags         todo
// @Produce      json
// @Param        id  path      string  true  "search item by id"
// @Success      200  {object}  database.TodoItem
// @Router       /:id [get]
// @Failure      400  "Invalid request parameter"
// @Failure      404  "Item not found"
// @Failure      500  "Internal server error"
func (a *App) GetById(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	item, err := a.db.FindById(id)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, item)
}
