package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mainHandle(t *testing.T) {
	type args struct {
		url                string
		expectedStatusCode int
		expectedCount      int
		expectedBody       string
	}

	args1 := args{
		url:                "/cafe?count=2&city=moscow",
		expectedStatusCode: http.StatusOK,
		expectedCount:      2,
		expectedBody:       "Мир кофе,Сладкоежка",
	}

	args2 := args{
		url:                "/cafe?count=10&city=omsk",
		expectedStatusCode: http.StatusBadRequest,
		expectedCount:      0,
		expectedBody:       "wrong city value",
	}

	args3 := args{
		url:                "/cafe?count=10&city=moscow",
		expectedStatusCode: http.StatusOK,
		expectedCount:      4,
		expectedBody:       "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should return 2 cafes, staus code 200",
			args: args1,
		},
		{
			name: "Should return 'wrong city value', status code 400",
			args: args2,
		},
		{
			name: "Should return 4 cafes, status code 200",
			args: args3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := getResponse(tt.args.url)
			require.Equal(t, tt.args.expectedStatusCode, responseRecorder.Code,
				"wrong status code, want %d, got %d", tt.args.expectedStatusCode, responseRecorder.Code)

			require.Equal(t, tt.args.expectedBody, responseRecorder.Body.String(),
				"wrong body, want %s, got %s", tt.args.expectedBody, responseRecorder.Body.String())
		})
	}
}

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}
