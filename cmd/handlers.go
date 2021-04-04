package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GET /articles
func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	randomSleep(start, 3000)

	defer app.logger.Success("GET", strconv.FormatInt(time.Since(start).Milliseconds(), 10))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

// POST /articles/new
func (app *application) Create(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	randomSleep(start, 3000)

	defer app.logger.Success("POST", strconv.FormatInt(time.Since(start).Milliseconds(), 10))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

// PUT /articles/update
func (app *application) Update(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	randomSleep(start, 3000)

	defer app.logger.Success("PUT", strconv.FormatInt(time.Since(start).Milliseconds(), 10))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

// DELETE /articles/deleet
func (app *application) Delete(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	randomSleep(start, 3000)

	defer app.logger.Success("DELETE", strconv.FormatInt(time.Since(start).Milliseconds(), 10))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func randomSleep(start time.Time, delay int) {
	rand.Seed(start.UTC().UnixNano())
	n := rand.Intn(delay)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
