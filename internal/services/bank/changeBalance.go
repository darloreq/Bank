package bank

import "coolBank/internal/entity"

func (b *Bank) ChangeBalance(userID int, amount entity.ChangeBalance, operationType string) (entity.Balance, error) {

	user := entity.User{ID: userID}

	switch operationType {
	case "put":
		balance, err := b.PutMoneyIn(user, amount)
		if err != nil {
			return entity.Balance{}, err
		}
		return balance, nil

	case "take":
		balance, err := b.TakeMoneyFrom(user, amount)
		if err != nil {
			return entity.Balance{}, err
		}
		return balance, nil

	default:
		return entity.Balance{}, WrongOperationType
	}

}
