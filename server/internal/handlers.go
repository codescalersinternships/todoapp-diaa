package App

import (
	"net/http"
	"strconv"

	"github.com/codescalersinternships/todo-diaa/internal/database"
	"github.com/gin-gonic/gin"
)

func (a *app) AddItem(ctx *gin.Context) {

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

func (a *app) FindAll(ctx *gin.Context) {

	todoList, err := a.db.FindAll()

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, todoList)
}

func (a *app) UpdateItem(ctx *gin.Context) {
	var item database.TodoItem
	err := ctx.BindJSON(&item)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = a.db.Update(item)

	if err == database.ErrIdNotFound {

		ctx.Status(http.StatusBadRequest)
		return
	}

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, "updated successfully")
}

func (a *app) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.db.Delete(id)

	if err==database.ErrIdNotFound{
		ctx.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "deleted successfully")

}

func (a *app) GetById(ctx *gin.Context) {

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
