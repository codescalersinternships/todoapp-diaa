package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/codescalersinternships/todo-diaa/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	tempDir := t.TempDir()
	app, err := NewApp(path.Join(tempDir, "todo.db"))
	assert.Nil(t, err)
	defer app.db.CloseDB()
	app.RegisterHandlers()

	testCases := []struct {
		name                 string
		expectedResponseCode int
		requestBody          map[string]string
	}{
		{
			name:                 "valid request body",
			expectedResponseCode: http.StatusCreated,
			requestBody: map[string]string{
				"name": "my todo",
			},
		},
		{
			name:                 "invalid request body",
			expectedResponseCode: http.StatusBadRequest,
			requestBody: map[string]string{
				"invalid": "diaa",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			jsonRequestBody, err := json.Marshal(tc.requestBody)

			assert.Nil(t, err)
			request, err := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(jsonRequestBody))

			request.Header.Set("Content-Type", "application/json")

			assert.Nil(t, err)

			response := httptest.NewRecorder()

			app.Router.ServeHTTP(response, request)

			assert.Equal(t, tc.expectedResponseCode, response.Code, "failed in response code")
		})
	}
}

func TestFindAll(t *testing.T) {
	tempDir := t.TempDir()
	app, err := NewApp(path.Join(tempDir, "todo.db"))
	assert.Nil(t, err)

	defer app.db.CloseDB()
	app.RegisterHandlers()

	request, err := http.NewRequest(http.MethodGet, "/todo", nil)

	assert.Nil(t, err)

	response := httptest.NewRecorder()

	app.Router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindById(t *testing.T) {
	tempDir := t.TempDir()
	app, err := NewApp(path.Join(tempDir, "todo.db"))
	assert.Nil(t, err)

	defer app.db.CloseDB()
	app.RegisterHandlers()

	insertItemInDB(t, app)

	testCases := []struct {
		name                 string
		expectedResponseCode int
		testID               int
	}{
		{
			name:                 "id exists",
			expectedResponseCode: http.StatusOK,
			testID:               1,
		},
		{
			name:                 "id does not exist",
			expectedResponseCode: http.StatusNotFound,
			testID:               2,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo/%d", tc.testID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.Nil(t, err)

			response := httptest.NewRecorder()

			app.Router.ServeHTTP(response, request)

			assert.Equal(t, tc.expectedResponseCode, response.Code)
		})
	}

}

func TestUpdateItem(t *testing.T) {
	tempDir := t.TempDir()
	app, err := NewApp(path.Join(tempDir, "todo.db"))
	assert.Nil(t, err)

	defer app.db.CloseDB()
	app.RegisterHandlers()

	insertItemInDB(t, app)

	testCases := []struct {
		name                 string
		expectedResponseCode int
		itemData             database.TodoItem
	}{
		{
			name:                 "should update successfully",
			expectedResponseCode: http.StatusCreated,
			itemData: database.TodoItem{
				ID:       1,
				Name:     "new name",
				Finished: 0,
			},
		},

		{
			name:                 "item id not found",
			expectedResponseCode: http.StatusNotFound,
			itemData: database.TodoItem{
				Name:     "todo",
				ID:       2,
				Finished: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tc.itemData)
			assert.Nil(t, err)
			request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/todo/%d", tc.itemData.ID), bytes.NewReader(requestBody))
			assert.Nil(t, err)

			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			app.Router.ServeHTTP(response, request)

			assert.Equal(t, tc.expectedResponseCode, response.Code)
		})
	}

}

func TestDeleteItem(t *testing.T) {
	tempDir := t.TempDir()
	app, err := NewApp(path.Join(tempDir, "todo.db"))
	assert.Nil(t, err)

	defer app.db.CloseDB()
	app.RegisterHandlers()

	insertItemInDB(t, app)

	testCases := []struct {
		name                 string
		expectedResponseCode int
		testID               int
	}{
		{
			name:                 "id exists",
			expectedResponseCode: http.StatusOK,
			testID:               1,
		},
		{
			name:                 "id does not exist",
			expectedResponseCode: http.StatusNotFound,
			testID:               2,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo/%d", tc.testID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			assert.Nil(t, err)

			response := httptest.NewRecorder()

			app.Router.ServeHTTP(response, request)

			assert.Equal(t, tc.expectedResponseCode, response.Code)
		})
	}
}

func insertItemInDB(t *testing.T, a *App) {
	requestBody := database.TodoItem{
		Name: "my todo",
	}
	bodyInJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)
	request, err := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(bodyInJson))
	assert.Nil(t, err)
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusCreated)

}
