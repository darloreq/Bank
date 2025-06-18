package handlers

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/handlers/mocks"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_PutMoneyIn(t *testing.T) {
	type mockBehavior func(m *bank.MockHeadHandler, user entity.User)
	testTable := []struct {
		name                 string
		inputBody            string
		user                 entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody any
	}{
		{
			name:      "success put money",
			inputBody: `{"ID": "0", "operation type": "put", amount: 100}`,
			user: entity.User{
				Name:    "<Test>",
				ID:      0,
				Balance: entity.Balance{Numbers: 0},
			},
			mockBehavior: func(m *bank.MockHeadHandler, user entity.User) {
				m.EXPECT().ChangeBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(entity.Balance{Numbers: user.Balance.Numbers}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: entity.Balance{Numbers: 100},
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
			r.Put("/", h.PutMoneyIn)

			//Test Request
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, "/", strings.NewReader(tt.inputBody))
			if err != nil {
				log.Fatal(err)
			}
			r.ServeHTTP(w, req)
			actualResponseBody := entity.Balance{}
			json.Unmarshal(w.Body.Bytes(), &actualResponseBody)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, actualResponseBody)
		})
	}
}
