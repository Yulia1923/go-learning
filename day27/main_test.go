package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserSuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/users/1",
		nil,
	)
	//request：请求:模拟客户端发送 GET /users/1，但目前还没有交给 API 处理。
	recorder := httptest.NewRecorder()
	//recorder：响应记录器，准备好接收并记录 API 返回的状态码、响应头和响应体。
	router.ServeHTTP(recorder, request) //让 Gin 路由对象处理这个模拟请求，并把产生的响应写入 recorder
	//检查状态码
	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"期望状态码 %d, 实际得到 %d",
			http.StatusOK,
			recorder.Code,
		)
	}
	//检查响应体
	expectedBody := `{"id":1,"name":"KimYoonsik"}`

	if recorder.Body.String() != expectedBody {
		t.Errorf(
			"期待响应%s, 实际得到 %s",
			expectedBody,
			recorder.Body.String(),
		)
	}
}
func TestGetUserInvalidID(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/users/abc",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"期望状态码 %d, 实际得到 %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}
func TestGetUserNotFound(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/users/3",
		nil,
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf(
			"期望状态码 %d, 实际得到 %d",
			http.StatusNotFound,
			recorder.Code,
		)
	}
}
func TestGetUserZeroID(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/users/0",
		nil,
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"期望状态码 %d, 实际得到 %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}
