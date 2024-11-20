package postgres

import (
	"context"
	"coolBank/internal/entity"
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
	//TODO implement me
	panic("implement me")
}

func (d *db) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) MakeUser() entity.User {
	//TODO implement me
	panic("implement me")
}
