package handlers

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/handlers/mocks"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_PutMoneyIn(t *testing.T) {
	type mockBehavior func(m *bank.MockHeadHandler)
	testTable := []struct {
		name                 string
		userID               int
		body                 amount
		bankUseCase          mockBehavior
		expectedStatusCode   int
		expectedResponseBody any
	}{
		{
			name: "success put money",
			bankUseCase: func(m *bank.MockHeadHandler) {
				m.EXPECT().ChangeBalance(0, entity.ChangeBalance{Amount: 100}, "put").Return(entity.Balance{Numbers: 100}, nil)
			},
			userID: 0,
			body: amount{
				OperationType: "put",
				TotalChange:   100,
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: entity.Balance{Numbers: 100},
		},
		{
			name: "success take money",
			bankUseCase: func(m *bank.MockHeadHandler) {
				m.EXPECT().ChangeBalance(0, entity.ChangeBalance{Amount: 100}, "take").Return(entity.Balance{Numbers: 0}, nil)
			},
			userID: 0,
			body: amount{
				OperationType: "take",
				TotalChange:   100,
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: entity.Balance{Numbers: 0},
		},
		{
			name: "wrong operation type", //ожидаем тейк, но приходит пут. Хендлер должен отдать нам бедРеквест, но вылетает ошибка
			bankUseCase: func(m *bank.MockHeadHandler) {
				m.EXPECT().ChangeBalance(0, entity.ChangeBalance{Amount: 100}, "take").Return(entity.Balance{Numbers: 0}, http.StatusBadRequest)
			},
			userID: 0,
			body: amount{
				OperationType: "put",
				TotalChange:   100,
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: entity.Balance{Numbers: 0},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newMock := bank.NewMockHeadHandler(ctrl)
			tt.bankUseCase(newMock)

			h := New(newMock)

			requestBody, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatal(err)
				return
			}

			//Test Server
			r := chi.NewRouter()
			r.Put("/{UserID}", h.PutMoneyIn)

			//Test Request
			w := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodPut, "/"+fmt.Sprintf("%d", tt.userID), strings.NewReader(string(requestBody)))
			if err != nil {
				t.Fatal(err)
				return
			}

			r.ServeHTTP(w, req)
			actualResponseBody := entity.Balance{}
			json.Unmarshal(w.Body.Bytes(), &actualResponseBody)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, actualResponseBody)
		})
	}
}
