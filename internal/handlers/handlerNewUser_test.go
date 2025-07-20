package handlers

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/handlers/mocks"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_NewUser(t *testing.T) {
	type mockBehavior func(m *bank.MockHeadHandler, user entity.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "success",
			inputBody: `{"name": "Test"}`,
			inputUser: entity.User{
				Name:    "Test",
				Balance: entity.Balance{0},
			},
			mockBehavior: func(m *bank.MockHeadHandler, user entity.User) {
				m.EXPECT().CreateUser(gomock.Any()).Return(user)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"ID":0,"Balance":{"Numbers":0},"Name":"Test"}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newMock := bank.NewMockHeadHandler(ctrl)
			tt.mockBehavior(newMock, tt.inputUser)

			h := New(newMock)

			//Test Server
			r := chi.NewRouter()
			r.Post("/{UserID}", h.NewUser)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/signIn", strings.NewReader(tt.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
