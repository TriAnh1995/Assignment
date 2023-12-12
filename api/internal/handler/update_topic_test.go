package handler

import (
	"assignment/internal/controller"
	"assignment/internal/model"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_UpdateTopic(t *testing.T) {
	testCases := []struct {
		Name            string
		requestInput    string
		request         model.UpdateInfo
		expectedCtrl    []string
		expectedCtrlErr error
		expectedRespond string
		expectedStatus  int
	}{
		{
			Name:            "Success",
			requestInput:    `{"sender": "%s","text":"%s"}`,
			request:         model.UpdateInfo{Sender: "alice@example.com", Text: "Hello World! kate@example.com"},
			expectedCtrl:    []string{"kate@example.com", "bob@example.com", "charlie@example.com"},
			expectedCtrlErr: nil,
			expectedRespond: "{\"message\":\"Success: true\"}[\n    \"kate@example.com\",\n    \"bob@example.com\",\n    \"charlie@example.com\",\n    \"kate@example.com\"\n]",
			expectedStatus:  200,
		},
		{
			Name:            "Failed to get your information",
			requestInput:    `"sender": "%s","text":"%s"`,
			request:         model.UpdateInfo{Sender: "alice@example.com", Text: "Hello World! kate@example.com"},
			expectedCtrl:    nil,
			expectedCtrlErr: nil,
			expectedRespond: "{\"error\":\"Failed to get your information\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Invalid Email Format",
			requestInput:    `{"sender": "%s","text":"%s"}`,
			request:         model.UpdateInfo{Sender: "aliceExample.com", Text: "Hello World! kate@example.com"},
			expectedCtrl:    nil,
			expectedCtrlErr: nil,
			expectedRespond: "{\"error\":\"Invalid Email Format\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Internal server error",
			requestInput:    `{"sender": "%s","text":"%s"}`,
			request:         model.UpdateInfo{Sender: "alice@example.com", Text: "Hello World! kate@example.com"},
			expectedCtrl:    nil,
			expectedCtrlErr: controller.ServerError,
			expectedRespond: "{\"error\":\"Internal Server Error\"}",
			expectedStatus:  500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new request
			reqBody := []byte(fmt.Sprintf(tc.requestInput, tc.request.Sender, tc.request.Text))
			req := httptest.NewRequest(http.MethodPost, "/users/update", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			// Set up a recorder to capture the response from the handler
			res := httptest.NewRecorder()

			// Set up and define mock behavior
			ctrl := new(controller.MockController)

			ctrl.On("UpdateTopic", req.Context(), tc.request).
				Return(tc.expectedCtrl, tc.expectedCtrlErr)

			// Create an instance of the handler with the mock controller
			instance := New(ctrl)
			handler := instance.UpdateTopic()

			// Create a context for testing and pass the request
			c, _ := gin.CreateTestContext(res)
			c.Request = req

			// Execute the handler function
			handler(c)

			// Review the results
			require.Equal(t, tc.expectedRespond, res.Body.String())
			require.Equal(t, tc.expectedStatus, res.Code)
		})
	}
}
