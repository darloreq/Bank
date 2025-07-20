package bank

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/services/bank/mocks"
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
			name: "success create user",
			input: entity.CreateUser{
				Name: "test",
			},
			mockBehavior: func(m *bank.MockBankReposI, user entity.CreateUser) {
				m.EXPECT().MakeUser(entity.CreateUser{Name: "test"}).Return(entity.User{ID: 0, Name: "test", Balance: entity.Balance{Numbers: 0}}, nil)
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

			NewTestUser, err := b.CreateUser(tt.input)
			if err != nil {
				log.Fatal()
			}

			assert.Equal(t, tt.expected, NewTestUser)

		})
	}
}
