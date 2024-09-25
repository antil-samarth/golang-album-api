package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// TestGetAlbums tests the GET /albums endpoint
func TestGetAlbums(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", getAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", w.Code)
	}

	var response []album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if len(response) != len(albums) {
		t.Errorf("Expected %d albums; got %d", len(albums), len(response))
	}
}

// TestPostAlbums tests the POST /albums endpoint
func TestPostAlbums(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", postAlbums)

	newAlbum := album{ID: "4", Title: "Test Album", Artist: "Test Artist", Price: 19.99}
	jsonValue, _ := json.Marshal(newAlbum)

	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status Created; got %v", w.Code)
	}

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if !reflect.DeepEqual(newAlbum, response) {
		t.Errorf("Expected %v; got %v", newAlbum, response)
	}
}

// TestGetAlbumByID tests the GET /albums/:id endpoint for an existing album
func TestGetAlbumByID(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", getAlbumsByID)

	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", w.Code)
	}

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if !reflect.DeepEqual(albums[0], response) {
		t.Errorf("Expected %v; got %v", albums[0], response)
	}
}

// TestGetAlbumByIDNotFound tests the GET /albums/:id endpoint for a non-existent album
func TestGetAlbumByIDNotFound(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", getAlbumsByID)

	req, _ := http.NewRequest("GET", "/albums/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status Not Found; got %v", w.Code)
	}
}
