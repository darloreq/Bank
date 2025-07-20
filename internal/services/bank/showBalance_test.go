package bank

import (
	"coolBank/internal/entity"
	bank "coolBank/internal/services/bank/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBank_ShowBalance(t *testing.T) {
	type mockBehavior func(m *bank.MockBankReposI, user entity.User)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		input        entity.User
		expected     entity.Balance
		expectedErr  error
	}{
		{
			name:  "success",
			input: entity.User{ID: 0},
			mockBehavior: func(m *bank.MockBankReposI, user entity.User) {
				m.EXPECT().ShowBalance(user.ID).Return(entity.Balance{Numbers: 0}, nil)
			},
			expected:    entity.Balance{0},
			expectedErr: nil,
		},
		{
			name:  "error case",
			input: entity.User{ID: 1},
			mockBehavior: func(m *bank.MockBankReposI, user entity.User) {
				m.EXPECT().ShowBalance(user.ID).Return(entity.Balance{}, errors.New("not found"))
			},
			expected:    entity.Balance{},
			expectedErr: errors.New("not found"),
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			newBank := bank.NewMockBankReposI(ctrl)
			tt.mockBehavior(newBank, tt.input)

			b := New(newBank)

			TestBalance, err := b.ShowBalance(tt.input)
			if tt.expectedErr != nil {
				assert.Error(t, err)                              // Проверяем что ошибка есть
				assert.EqualError(t, err, tt.expectedErr.Error()) // Проверяем текст ошибки
			} else {
				assert.NoError(t, err) // Проверяем что ошибки нет
			}

			assert.Equal(t, tt.expected, TestBalance)

		})
	}

}
