package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todoapp "github.com/baza04/todoApp"
)

// @Summary Create TODO Item
// @Security ApiKeyAuth
// @Description will create new item in list of current user
// @ID create-todo-item
// @Tags items
// @Accepted json
// @Produce json
// @Param payload body todoapp.TodoItem true "item info"
// @Param id path int true "id"
// @Success 201 {string} string "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todoapp.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}

type getAllItemsResponse struct {
	Data []todoapp.TodoItem
}

// @Summary Get All TODO Items
// @Security ApiKeyAuth
// @Description return all items from choosen list of current user
// @ID get-all-todo-items
// @Tags items
// @Accepted json
// @Produce json
// @Param list_id path int true "list_id"
// @Success 200 {object} handler.getAllItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id}/items [get]
func (h *Handler) getAllItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

// @Summary Get TODO Items By ID
// @Security ApiKeyAuth
// @Description return item by ID of current user
// @ID get-todo-item-by-id
// @Tags items
// @Accepted json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} todoapp.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetByID(userID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update TODO Items By ID
// @Security ApiKeyAuth
// @Description update item by ID of current user
// @ID update-todo-item
// @Tags items
// @Accepted json
// @Produce json
// @Param payload body todoapp.UpdateItemInput true "updated item info"
// @Param id path int true "id"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	var (
		userID, itemID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if itemID, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	var input todoapp.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err = h.services.TodoItem.Update(userID, itemID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete TODO Items By ID
// @Security ApiKeyAuth
// @Description delete item by ID of current user
// @ID delete-todo-item
// @Tags items
// @Accepted json
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	var (
		userID, itemID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if itemID, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err = h.services.TodoItem.Delete(userID, itemID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
