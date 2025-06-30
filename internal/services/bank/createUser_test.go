package bank

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/services/bank/mocks"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBank_CreateUser(t *testing.T) {
	type mockBehavior func(m *bank.MockBankReposI, user entity.CreateUser)
	testTable := []struct {
		name         string
		input        entity.CreateUser
		mockBehavior mockBehavior
		expected     entity.User
	}{
		{
			name: "success",
			input: entity.CreateUser{
				Name: "test",
			},
			mockBehavior: func(m *bank.MockBankReposI, user entity.CreateUser) {
				m.EXPECT().MakeUser(gomock.Any()).Return(entity.User{}, nil)
			},
			expected: entity.User{
				Name:    "test",
				ID:      0,
				Balance: entity.Balance{Numbers: 0},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newBank := bank.NewMockBankReposI(ctrl)
			tt.mockBehavior(newBank, tt.input)

			b := New(newBank)

			var TestUser entity.CreateUser

			err := json.Unmarshal([]byte(tt.input.Name), &TestUser)
			if err != nil {
				log.Fatal()
			}

			NewTestUser := b.CreateUser(TestUser) //TODO Не съедает имя

			assert.Equal(t, tt.expected, NewTestUser)

		})
	}
}
