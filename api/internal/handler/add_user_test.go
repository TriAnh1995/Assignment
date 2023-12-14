package handler

import (
	"assignment/internal/controller"
	"assignment/internal/model"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_AddUsers(t *testing.T) {

	testCases := []struct {
		Name            string
		requestinput    string
		request         model.User
		expectedCtrl    error
		expectedRespond string
		expectedStatus  int
	}{

		{
			Name:            "Success",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "Alice", Email: "alice@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"message\":\"Add user successfully!\"}",
			expectedStatus:  200,
		},
		{
			Name:            "Name cannot be blank",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "", Email: "alice@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Name cannot be blank\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Email cannot be blank",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "Alice", Email: ""},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Email cannot be blank\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Failed to get your information",
			requestinput:    `"name": "%s","email":"%s"`,
			request:         model.User{Name: "Alice", Email: "alice@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Failed to get your information\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Name Invalid",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "alice", Email: "alice@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Name Invalid\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Email Invalid",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "Alice", Email: "alice%example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Email invalid\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Internal server error",
			requestinput:    `{"name": "%s","email":"%s"}`,
			request:         model.User{Name: "Alice", Email: "alice@example.com"},
			expectedCtrl:    errors.New("Internal server error"),
			expectedRespond: "{\"error\":\"Internal Server Error\"}",
			expectedStatus:  500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// Create new Request
			reqBody := []byte(fmt.Sprintf(tc.requestinput, tc.request.Name, tc.request.Email))

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))

			req.Header.Set("Content-Type", "application/json")

			// Set up a record to the response from handler
			res := httptest.NewRecorder()

			// Setup and defined mock behavior
			ctrl := new(controller.MockController)

			ctrl.On("AddUsers", req.Context(), tc.request).
				Return(tc.expectedCtrl)

			// Setup instance to use mock file in test
			instance := New(ctrl)

			handler := instance.AddUsers()

			// Create context for test, and pass Request for it
			c, _ := gin.CreateTestContext(res)
			c.Request = req

			// Run function in handler with request context and mock file
			handler(c)

			// Review the result
			require.Equal(t, tc.expectedRespond, res.Body.String())
			require.Equal(t, tc.expectedStatus, res.Code)
		})
	}
}
