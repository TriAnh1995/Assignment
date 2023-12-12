package handler

import (
	"assignment/internal/controller"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_BlockUsers(t *testing.T) {
	testCases := []struct {
		Name            string
		requestInput    string
		request         []string
		expectedCtrl    error
		expectedRespond string
		expectedStatus  int
	}{
		{
			Name:            "Success",
			requestInput:    `{"requester": "%s","target":"%s"}`,
			request:         []string{"requester@example.com", "target@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"message\":\"Blocked successfully!\"}",
			expectedStatus:  200,
		},
		{
			Name:            "Failed to get your information",
			requestInput:    `"requester": "%s","target":"%s"`,
			request:         []string{"requester@example.com", "target@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Failed to get your information\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Please insert at least two different emails",
			requestInput:    `{"requester": "%s","target":"%s"}`,
			request:         []string{"requester@example.com", "requester@example.com"},
			expectedCtrl:    nil,
			expectedRespond: "{\"error\":\"Please insert two different emails\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Internal server error",
			requestInput:    `{"requester": "%s","target":"%s"}`,
			request:         []string{"requester@example.com", "target@example.com"},
			expectedCtrl:    errors.New("Internal server error"),
			expectedRespond: "{\"error\":\"Internal Server Error\"}",
			expectedStatus:  500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create new Request
			reqBody := []byte(fmt.Sprintf(tc.requestInput, tc.request[0], tc.request[1]))

			req := httptest.NewRequest(http.MethodPost, "/subscriptions/block", bytes.NewBuffer(reqBody))

			req.Header.Set("Content-Type", "application/json")

			// Set up a record to the response from handler
			res := httptest.NewRecorder()

			// Setup and defined mock behavior
			ctrl := new(controller.MockController)

			ctrl.On("BlockUsers", req.Context(), tc.request).
				Return(tc.expectedCtrl)

			// Setup instance to use mock file in test
			instance := New(ctrl)

			handler := instance.BlockUsers()

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
