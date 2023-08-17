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
// @Router       /todo/ [post]
// @Failure      400   "Invalid request body"
// @Failure      500   "Internal server error"
func (a *App) HandleAddItem(ctx *gin.Context) {

	var item database.TodoItem
	err := ctx.BindJSON(&item)

	if err != nil || item.Name == "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	item, err = a.db.Insert(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, item)
}

// @Summary      Get todo items array
// @Description  Responds with the list of all todo items as JSON.
// @Tags         todo
// @Produce      json
// @Router       /todo/ [get]
// @Success      200   {object}  []database.TodoItem
// @Failure      500   "Internal server error"
func (a *App) HandleFindAll(ctx *gin.Context) {

	todoList, err := a.db.FindAll()

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
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
// @Router       /todo/ [put]
// @Failure      400  "Invalid request body"
// @Failure      500  "Internal server error"
func (a *App) HandleUpdateItem(ctx *gin.Context) {
	var item database.TodoItem
	err := ctx.BindJSON(&item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = a.db.Update(item)

	if err == database.ErrIDNotFound {

		ctx.JSON(http.StatusNotFound, database.ErrIDNotFound)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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
// @Router       /todo/:id [delete]
// @Failure      404  "Item not found"
// @Failure      500  "Internal server error"
func (a *App) HandleDeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.db.Delete(id)

	if err == database.ErrIDNotFound {
		ctx.JSON(http.StatusNotFound, database.ErrIDNotFound)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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
// @Router       /todo/:id [get]
// @Failure      400  "Invalid request parameter"
// @Failure      404  "Item not found"
// @Failure      500  "Internal server error"
func (a *App) HandleGetById(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	item, err := a.db.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, item)
}
