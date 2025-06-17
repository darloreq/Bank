package handlers

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/handlers/mocks"
	"testing"
)

func TestHandler_PutMoneyIn(t *testing.T) {
	type mockBehavior func(m *bank.MockHeadHandler, user entity.User)
	testTable := []struct {
		name     string
		inputBody string
		user     entity.User
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedResponseBody any
	}{
		{
			name: "success",
			inputBody: `{}`,
			user: entity.User{
				Name: "<Test>",
				ID: 0,
				Balance: entity.Balance{Numbers: 0},
			},
			mockBehavior: func(m *bank.MockHeadHandler, user entity.User) {
				m.EXPECT(),
			}
		},
	}
}
