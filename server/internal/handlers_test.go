package App

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
	app := NewApp()
	tempDir := t.TempDir()
	err := app.StartDB(path.Join(tempDir, "todo.db"))

	assert.Nil(t, err, "failed to start database")
	defer app.db.CloseDB()
	app.SetAPIs()

	testCases := []struct {
		name                 string
		expectedResponseCode int
		expectedResponseBody database.TodoItem
		requestBody          map[string]string
	}{
		{
			name:                 "valid request body",
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: database.TodoItem{
				ID:       1,
				Name:     "my todo",
				Finished: 0,
			},
			requestBody: map[string]string{
				"name": "my todo",
			},
		},
		{
			name:                 "invalid request body",
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: database.TodoItem{},
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

			var actualResponseBody database.TodoItem
			json.Unmarshal(response.Body.Bytes(), &actualResponseBody)

			if tc.expectedResponseCode == http.StatusCreated {
				assert.Equal(t, tc.expectedResponseBody, actualResponseBody)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	app := NewApp()
	tempDir := t.TempDir()
	err := app.StartDB(path.Join(tempDir, "todo.db"))

	assert.Nil(t, err, "failed to start database")
	defer app.db.CloseDB()
	app.SetAPIs()

	request, err := http.NewRequest(http.MethodGet, "/todo", nil)

	assert.Nil(t, err)

	response := httptest.NewRecorder()

	app.Router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindById(t *testing.T) {
	app := NewApp()
	tempDir := t.TempDir()
	err := app.StartDB(path.Join(tempDir, "todo.db"))

	assert.Nil(t, err, "failed to start database")
	defer app.db.CloseDB()
	app.SetAPIs()

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
	app := NewApp()
	tempDir := t.TempDir()
	err := app.StartDB(path.Join(tempDir, "todo.db"))

	assert.Nil(t, err, "failed to start database")
	defer app.db.CloseDB()
	app.SetAPIs()

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
			expectedResponseCode: http.StatusBadRequest,
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
			request, err := http.NewRequest(http.MethodPut, "/todo", bytes.NewReader(requestBody))
			assert.Nil(t, err)

			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			app.Router.ServeHTTP(response, request)

			assert.Equal(t, tc.expectedResponseCode, response.Code)
		})
	}

}

func TestDeleteItem(t * testing.T){
	app := NewApp()
	tempDir := t.TempDir()
	err := app.StartDB(path.Join(tempDir, "todo.db"))

	assert.Nil(t, err, "failed to start database")
	defer app.db.CloseDB()
	app.SetAPIs()

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

func insertItemInDB(t *testing.T, a *app) {
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