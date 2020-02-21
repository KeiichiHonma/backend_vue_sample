package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestIndexHandler(t *testing.T) {
	r := InitRestApi()
	req := httptest.NewRequest("GET","/",nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec,req)
	assert.Equal(t,http.StatusOK, rec.Code)
	jsonBytes := ([]byte)(rec.Body.String())
	data := new(IndexResponse)
	if err := json.Unmarshal(jsonBytes, data);err != nil{
		t.Error("失敗",rec.Body.String())
		return
	}
	assert.Equal(t, "投稿API",data.Message)
}
