package handler

import (
	"assignment/internal/controller"
	"assignment/internal/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_CommonFriends(t *testing.T) {
	testCases := []struct {
		Name            string
		request         []string
		expectedCtrl    model.FriendshipInfo
		expectedCtrlErr error
		expectedRespond string
		expectedStatus  int
	}{
		{
			Name:    "Success",
			request: []string{"user1@example.com", "user2@example.com"},
			expectedCtrl: model.FriendshipInfo{
				List:    []string{"friend_in_common_1@example.com", "friend_in_common_1@example.com"},
				Amounts: 2,
			},
			expectedRespond: "{\n    \"List\": [\n        \"friend_in_common_1@example.com\",\n        \"friend_in_common_1@example.com\"\n    ],\n    \"Amounts\": 2\n}",
			expectedCtrlErr: nil,
			expectedStatus:  200,
		},
		{
			Name:            "Please insert at least two different emails",
			request:         []string{"user1@example.com", "user1@example.com"},
			expectedCtrl:    model.FriendshipInfo{},
			expectedCtrlErr: nil,
			expectedRespond: "{\"error\":\"Please insert two different emails\"}",
			expectedStatus:  400,
		},
		{
			Name:            "Internal server error",
			request:         []string{"user1@example.com", "user2@example.com"},
			expectedCtrl:    model.FriendshipInfo{},
			expectedCtrlErr: errors.New("Internal server error"),
			expectedRespond: "{\"error\":\"Internal Server Error\"}",
			expectedStatus:  500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create new Request
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/friends/common?email1=%s&email2=%s", tc.request[0], tc.request[1]), nil)

			// Set up a record to the response from handler
			res := httptest.NewRecorder()

			// Setup and defined mock behavior
			ctrl := new(controller.MockController)
			ctrl.On("CommonFriends", req.Context(), tc.request).
				Return(tc.expectedCtrl, tc.expectedCtrlErr)

			// Setup instance to use mock file in test
			instance := New(ctrl)
			handler := instance.CommonFriends()

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
