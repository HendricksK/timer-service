package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	timer "github.com/HendricksK/timer-service/timer"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestTimers(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/timers/50", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, w.Body.String())
	// Need to set as default value
	// Need to build a seed script
	assert.Contains(t, w.Body.String(), "54b686fa-b0d4-4dfa-a312-e90d811c25bd")

}

func TestTimersRead(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/timer/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, w.Body.String())

}

func TestTimersCreate(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/timer", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, w.Body.String())

}

func TestTimersUpdate(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/timer/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, w.Body.String())

}

func TestTimersDelete(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/timer/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, w.Body.String())

}

func TestTimerCrud(t *testing.T) {
	fmt.Println()

	fmt.Println("Testing CRUD against array")
	fmt.Println("TestRead")
	assert.NotNil(t, timer.TestRead())
	fmt.Println("TestReadById")
	assert.NotNil(t, timer.TestReadById("wqdwqdwd878736gefduh"))
	fmt.Println("TestCreate")
	assert.NotNil(t, timer.TestCreate(timer.GetTestTimer()))
	// assert.NotNil(t, timer.TestUpdate("wqdwqdwd878736gefduh"))
	fmt.Println("TestDelete")
	assert.NotNil(t, timer.TestDelete("wqdwqdwd878736gefduh"))
}
