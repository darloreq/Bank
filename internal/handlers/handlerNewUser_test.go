package handlers

import (
	"coolBank/internal/entity"
	mock_bank "coolBank/internal/services/bank/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestHandler_NewUser(t *testing.T) {
	type mockBehavior func(m *mock_bank.MockBankReposI, user entity.User)

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
			inputBody: `{"name": "Test"'}`,
			inputUser: entity.User{
				Name: "Test",
				Balance: 0,

			},
			mockBehavior: func(m *mock_bank.MockBankReposI, user entity.User) {
				m.EXPECT().MakeUser(gomock.Any()).Return(user, nil),
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{"name":"Test"}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newUser := mock_bank.NewMockBankReposI(ctrl)
			tt.mockBehavior(newUser, tt.inputUser)
		})
	}
}
