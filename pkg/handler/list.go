package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todoapp "github.com/baza04/todoApp"
)

// @Summary Create TODO List
// @Security ApiKeyAuth
// @Description create new list for current user
// @Tags lists
// @ID create-list
// @Accept json
// @Produce json
// @Param payload body todoapp.TodoList true "list info"
// @Success 201 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	var (
		userID, listID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	var input todoapp.TodoList
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if listID, err = h.services.TodoList.Create(userID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"list_id": listID,
	})
}

type getAllListsResponse struct {
	Data []todoapp.TodoList `json:"data"`
}

// @Summary Get All TODO Lists
// @Security ApiKeyAuth
// @Description return all TODO lists of current user
// @Tags lists
// @ID get-all-lists
// @Accepted json
// @Produce json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	var (
		userID int
		lists  []todoapp.TodoList
		err    error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if lists, err = h.services.TodoList.GetAll(userID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get TODO List By ID
// @Security ApiKeyAuth
// @Description Get TODO List By choosen ID
// @ID get-list-by-id
// @Tags lists
// @Accepted json
// @Produce json
// @Param list_id path int true "list_id"
// @Success 200 {object} todoapp.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} handler.errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [get]
func (h *Handler) getListByID(c *gin.Context) {
	var (
		userID, id int
		list       todoapp.TodoList
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if list, err = h.services.TodoList.GetByID(userID, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update TODO List
// @Security ApiKeyAuth
// @Tags lists
// @Description Update choosen list info
// @ID update-list
// @Accept json
// @Produce json
// @Param payload body todoapp.UpdateListInput true "update_info"
// @Param list_id path int true "list_id"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} handler.errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [put]
func (h *Handler) updateList(c *gin.Context) {
	var (
		userID, id int
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todoapp.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(userID, id, &input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete TODO List
// @Security ApiKeyAuth
// @Description Delete TODO List
// @ID delete-todo-list
// @Tags lists
// @Accepted json
// @Produce json
// @Param list_id path int true "list_id"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	var (
		userID, id int
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Delete(userID, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
