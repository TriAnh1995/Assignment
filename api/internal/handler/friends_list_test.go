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

func TestHandler_FriendList(t *testing.T) {
	testCases := []struct {
		Name            string
		requestInput    string
		request         string
		expectedCtrl    model.FriendshipInfo
		expectedCtrlErr error
		expectedRespond string
		expectedStatus  int
	}{
		{
			Name:         "Success",
			requestInput: `{"email": "%s"}`,
			request:      "user@example.com",
			expectedCtrl: model.FriendshipInfo{
				List:    []string{"friend_number_1@example.com", "friend_number_2@example.com"},
				Amounts: 2,
			},
			expectedCtrlErr: nil,
			expectedRespond: "{\"message\":\"Success!\"}{\n    \"List\": [\n        \"friend_number_1@example.com\",\n        \"friend_number_2@example.com\"\n    ],\n    \"Amounts\": 2\n}",
			expectedStatus:  200,
		},
		{
			Name:         "Failed to get your information",
			requestInput: `"email": "%s"`,
			request:      "user@example.com",
			expectedCtrl: model.FriendshipInfo{
				List:    nil,
				Amounts: 0,
			},
			expectedCtrlErr: nil,
			expectedRespond: "{\"error\":\"Failed to get your information\"}",
			expectedStatus:  400,
		},
		{
			Name:         "Email invalid",
			requestInput: `{"email": "%s"}`,
			request:      "userExampleCom",
			expectedCtrl: model.FriendshipInfo{
				List:    nil,
				Amounts: 0,
			},
			expectedCtrlErr: nil,
			expectedRespond: "{\"error\":\"Invalid Email Format\"}",
			expectedStatus:  400,
		},
		{
			Name:         "Internal server error",
			requestInput: `{"email": "%s"}`,
			request:      "user@example.com",
			expectedCtrl: model.FriendshipInfo{
				List:    nil,
				Amounts: 0,
			},
			expectedCtrlErr: errors.New("Internal server error"),
			expectedRespond: "{\"error\":\"Internal Server Error\"}",
			expectedStatus:  500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new request
			reqBody := []byte(fmt.Sprintf(tc.requestInput, tc.request))
			req := httptest.NewRequest(http.MethodGet, "/friends/list", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			// Set up a recorder to capture the response from the handler
			res := httptest.NewRecorder()

			// Set up and define mock behavior
			ctrl := new(controller.MockController)

			ctrl.On("FriendsList", req.Context(), tc.request).
				Return(tc.expectedCtrl, tc.expectedCtrlErr)

			// Create an instance of the handler with the mock controller
			instance := New(ctrl)
			handler := instance.FriendsList()

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
