package handlers

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/handlers/mocks"
	"errors"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ShowBalance(t *testing.T) {
	type mockBehavior func(m *bank.MockHeadHandler, user entity.User)
	testTable := []struct {
		name                 string
		inputBody            string
		user                 entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "success",
			inputBody: `{"UserID":"0"}`,
			user: entity.User{
				Name:    "<Test>",
				ID:      0,
				Balance: entity.Balance{Numbers: 0},
			},
			mockBehavior: func(m *bank.MockHeadHandler, user entity.User) {
				m.EXPECT().ShowBalance(gomock.Any()).Return(entity.Balance{Numbers: user.Balance.Numbers}, errors.New("mock error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"Balance":"0"}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newMock := bank.NewMockHeadHandler(ctrl)
			tt.mockBehavior(newMock, tt.user)

			h := New(newMock)

			//Test Server
			r := chi.NewRouter()
			r.Get("/{UserID}", h.ShowBalance)

			//Test Request
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/"+tt.inputBody, nil)
			if err != nil {
				log.Fatal(err)
			}
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}

}
