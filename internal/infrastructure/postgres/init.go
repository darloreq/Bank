package postgres

import (
	"context"
	"coolBank/internal/entity"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func New() (*db, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:admin@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	return &db{conn}, err
}

type db struct {
	conn *pgx.Conn
}

func (d *db) ShowBalance(userID int) (entity.Balance, error) {
	//queryRow запрос одной строки, query запрос нескольких
	var balance entity.Balance

	err := d.conn.QueryRow(context.Background(), "SELECT user_balance FROM user_balance WHERE user_id = $1", userID).Scan(&balance.Numbers)
	if err != nil {
		fmt.Println("error in query balance", err)
		return balance, err
	}

	return balance, nil
}

func (d *db) PutMoneyInDB(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {
	var newBalance = entity.Balance{Numbers: amountPut.Amount}

	err := d.conn.QueryRow(context.Background(), "INSERT INTO user_balance (user_balance) WHERE user_id = $1 VALUES ($2) RETURNING user_balance", userID, newBalance.Numbers).Scan(&newBalance) //возвращаю значение из ДБ, потому что соответствующий метод должен отдать в БЛ новый баланс из ДБ
	if err != nil {
		fmt.Println("error in query balance", err)
		return newBalance, err
	}
	return newBalance, nil
}

// тебе не нужны 2 метода для изменения баланса, попробуй написать всё через changebalance
func (d *db) TakeMoneyInDB(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {

}

func (d *db) MakeUser(user entity.CreateUser) (entity.User, error) {
	var newBalance = entity.Balance{Numbers: 0}
	var newUser entity.User

	err := d.conn.QueryRow(context.Background(), "INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&newUser.ID)
	if err != nil {
		fmt.Println("error in query balance", err)
		return newUser, err
	}

	_, err = d.conn.Exec(context.Background(), "INSERT INTO user_balance (user_balance, user_id) VALUES ($1, $2)", newBalance.Numbers, newUser.ID)
	if err != nil {
		fmt.Println("error in query balance", err)
		return newUser, err
	}

	newUser.Balance = newBalance
	newUser.Name = user.Name

	return newUser, err
}
