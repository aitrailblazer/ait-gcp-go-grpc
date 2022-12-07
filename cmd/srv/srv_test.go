//
// Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

// create tests for the server
// Path: api/server/main_test.go

package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
)

//	func TestServer(t *testing.T) {
//		req, err := http.NewRequest("GET", "/v1/", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//		rr := httptest.NewRecorder()
//		handler := http.HandlerFunc(v1Handler)
//		handler.ServeHTTP(rr, req)
//		if status := rr.Code; status != http.StatusOK {
//			t.Errorf("handler returned wrong status code: got %v want %v",
//				status, http.StatusOK)
//		}
//		expected := "{status: 'running'}"
//		if rr.Body.String() != expected {
//			t.Errorf("handler returned unexpected body: got %v want %v",
//				rr.Body.String(), expected)
//		}
//	}
//
//	func Test_v1Handler(t *testing.T) {
//		tests := map[string]struct {
//			method     string
//			path       string
//			wantCode   int
//			wantResult string
//		}{
//			"simple":   {method: http.MethodGet, path: "/v1/", wantCode: http.StatusOK, wantResult: "{status: 'running'}"},
//			"error 01": {method: http.MethodPost, path: "/v1/", wantCode: http.StatusOK, wantResult: "{status: 'running error - expected http.MethodGet'}"},
//			"error 02": {method: http.MethodGet, path: "/v1/1", wantCode: http.StatusOK, wantResult: "{status: 'running error - expected /v1/'}"},
//		}
//		for name, tc := range tests {
//			t.Run(name, func(t *testing.T) {
//				req, err := httptest.NewRequest(tc.method, tc.path, nil)
//				if err != nil {
//					t.Fatal(err)
//				}
//				rr := httptest.NewRecorder()
//				handler := http.HandlerFunc(v1Handler)
//				handler.ServeHTTP(rr, req)
//				if status := rr.Code; status != tc.wantCode {
//					t.Errorf("handler returned wrong status code: got %v want %v",
//						status, tc.wantCode)
//				}
//				if rr.Body.String() != tc.wantResult {
//					t.Errorf("handler returned unexpected body: got %v want %v",
//						rr.Body.String(), tc.wantResult)
//				}
//			})
//		}
//	}

func Fuzz_v1Handler(f *testing.F) {
	m := map[string]bool{
		http.MethodConnect: true,
		http.MethodDelete:  true,
		http.MethodGet:     true,
		http.MethodOptions: true,
		http.MethodPatch:   true,
		http.MethodPost:    true,
		http.MethodPut:     true,
		http.MethodTrace:   true,
	}

	e := echo.New()

	testCases := []struct {
		name         string
		method       string
		path         string
		expectResult string
		expectStatus int

		wantError bool
	}{
		{name: "simple", method: http.MethodGet, path: "/v1/", expectResult: "{status: \"running\"}", expectStatus: http.StatusOK, wantError: false},
		{name: "error 01", method: http.MethodGet, path: "/v1/1/", expectResult: "{\"status\":\"Not Found\"}", expectStatus: http.StatusNotFound, wantError: false},
		{name: "error 02", method: http.MethodPost, path: "/v1/", expectResult: "{\"status\":\"Not Found\"}", expectStatus: http.StatusNotFound, wantError: false},
		{name: "error 03", method: http.MethodPost, path: "/", expectResult: "{\"status\":\"Not Found\"}", expectStatus: http.StatusNotFound, wantError: false},
	}
	for _, test := range testCases {
		f.Log(test.name, test.method, test.path)
		f.Add(test.name, test.method, test.path, test.expectResult, test.expectStatus, test.wantError)
	}
	f.Fuzz(func(t *testing.T, name, method, path, expectResult string, expectStatus int, wantError bool) {
		t.Log("method ", method, "path ", path)
		// check if path is valid
		_, err := url.ParseRequestURI(path)
		if err == nil {
			// check if method is allowed
			_, ok := m[method]
			if ok {
				c, _ := request(method, path, e)
				if (c != http.StatusOK) && (c != http.StatusNotFound) {
					t.Errorf("method %s path %s handler returned wrong status code: got %v expect %v",
						method, path, c, expectStatus)
				}
			}
		}
	})
}

func request(method, path string, e *echo.Echo) (int, string) {
	e = routerSetup(e)
	log.Println("request:method ", method)
	log.Println("request:path ", path)

	req := httptest.NewRequest(method, path, nil)
	req.Host = "localhost"
	rec := httptest.NewRecorder()
	e.Debug = true
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

// func Test_v1Handler(t *testing.T) {
// 	e := echo.New()
// 	testCases := []struct {
// 		name         string
// 		method       string
// 		path         string
// 		expectResult string
// 		expectStatus int
// 	}{
// 		{name: "simple", method: http.MethodGet, path: "/v1/", expectResult: "{status: \"running\"}", expectStatus: http.StatusOK},
// 		{name: "error 01", method: http.MethodGet, path: "/v1/1/", expectResult: "{status: \"running error - expected /v1/\"}", expectStatus: http.StatusBadRequest},
// 		// "error 02": {method: http.MethodPost, path: "/v1/", wantCode: http.StatusOK, wantResult: "{status: 'running error - expected http.MethodGet'}", wantErr: true},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Log("tc.method ", tc.method)
// 			t.Log("tc.path ", tc.path)

// 			c, b := request(tc.method, tc.path, e)
// 			if c != tc.expectStatus {
// 				t.Errorf("%s handler returned wrong status code: got %v expect %v",
// 					tc.name, c, tc.expectStatus)
// 			}
// 			if b != tc.expectResult {
// 				t.Errorf("%s handler returned unexpected body: got %v expect %v",
// 					tc.name, b, tc.expectResult)
// 			}
// 		})
// 	}
// }
