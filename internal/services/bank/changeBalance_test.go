package bank

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/services/bank/mocks"
	"testing"
)

func TestBank_ChangeBalance(t *testing.T) {
	type mockBehavior func(m *bank.MockBankReposI, user entity.User)
	testTable := []struct {
		name string
		input entity.User
		mockBehavior func(m *bank.MockBankReposI, user entity.User)
		expectedResponseBody entity.Balance
		expectedErrorResponse error
	}{
		{
			name: "success",
			input: entity.User{ID: 0},
			mockBehavior: func(m *bank.MockBankReposI, user entity.User) {
				m.EXPECT().
			}
		},
	}

}
