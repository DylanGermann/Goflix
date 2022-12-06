package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testStore struct {
	movieId int64
	movies  []*Movie
}

func (t testStore) Open() error {
	return nil
}

func (t testStore) Close() error {
	return nil
}

func (t testStore) GetMovies() ([]*Movie, error) {
	return t.movies, nil
}

func (t testStore) GetMovieById(id int64) (*Movie, error) {
	for _, m := range t.movies {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, nil
}

func (t testStore) CreateMovie(m *Movie) error {
	t.movieId++
	m.ID = t.movieId
	t.movies = append(t.movies, m)
	return nil
}

func (t testStore) FindUser(username string, password string) (bool, error) {
	return true, nil
}

func TestMovieCreateUnit(t *testing.T) {
	srv := newServer()
	srv.store = &testStore{}

	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "https://google.com",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	assert.Nil(t, err)

	r := httptest.NewRequest("POST", "/api/movies/", &buf)
	w := httptest.NewRecorder()

	srv.handleMovieCreate()(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMovieCreateIntegration(t *testing.T) {
	srv := newServer()
	srv.store = &testStore{}

	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "https://google.com",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	assert.Nil(t, err)
	r := httptest.NewRequest("POST", "/api/movies/", &buf)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAzNDAxODksImlhdCI6MTY3MDMzNjU4OSwidXNlcm5hbWUiOiJnb2xhbmcifQ.lHG4VWKAONP6AD1t6SqkwTk7KTMjzFNTa6aTGf9Bvs0"
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	w := httptest.NewRecorder()

	srv.serveHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
