package HTTP_server_testing_theme

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test",
			want: want{
				code:        200,
				response:    "{\"status\":\"ok\"}",
				contentType: "application/json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			w := httptest.NewRecorder()

			h := http.HandlerFunc(StatusHandler)
			h.ServeHTTP(w, request)
			res := w.Result()

			if statusCode := res.StatusCode; statusCode != tt.want.code {
				t.Errorf("want status code: %d, got: %d", tt.want.code, statusCode)
			}

			defer res.Body.Close()
			respBody, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if respBody := string(respBody); respBody != tt.want.response {
				t.Errorf("want response body: %s, got: %s", tt.want.response, respBody)
			}

			if contentType := res.Header.Get("Content-Type"); contentType != tt.want.contentType {
				t.Errorf("want content-type: %s, got: %s", tt.want.contentType, contentType)
			}
		})
	}
}

func TestUserViewHandler(t *testing.T) {
	type want struct {
		user User
	}

	tests := []struct {
		name        string
		enter       map[string]User
		userID      string
		statusCode  int
		contentType string
		want        want
	}{
		{
			name: "positive test",
			enter: map[string]User{
				"u1": {
					ID:        "u1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				"u2": {
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			userID: "u1",
			want: want{
				user: User{
					ID:        "u1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			statusCode:  200,
			contentType: "application/json",
		},
		{
			name: "error code 404",
			enter: map[string]User{
				"u1": {
					ID:        "u1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				"u2": {
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			userID:      "u3",
			statusCode:  404,
			contentType: "text/plain; charset=utf-8",
		},
		{
			name: "error code 400",
			enter: map[string]User{
				"u1": {
					ID:        "u1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				"u2": {
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			userID:      "",
			statusCode:  400,
			contentType: "text/plain; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := "/users" + "/?user_id=" + tt.userID
			req := httptest.NewRequest(http.MethodGet, target, nil)
			w := httptest.NewRecorder()

			h := http.HandlerFunc(UserViewHandler(tt.enter))
			h.ServeHTTP(w, req)
			resp := w.Result()

			if statusCode := resp.StatusCode; statusCode != tt.statusCode {
				t.Errorf("Want status code: %d, got: %d", tt.statusCode, statusCode)
			}

			defer resp.Body.Close()
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if contentType := resp.Header.Get("Content-Type"); contentType != tt.contentType {
				t.Errorf("Want content-type: %s, got: %s", tt.contentType, contentType)
			}

			var user User
			_ = json.Unmarshal(respBody, &user)

			if user != tt.want.user {
				t.Errorf("Want user: %s, got: %s", tt.want.user, respBody)
			}
		})
	}
}
