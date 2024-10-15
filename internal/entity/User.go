package entity

type User struct {
	ID      int
	Balance Balance
}

type Balance struct {
	Numbers float64
}

type ChangeBalance struct {
	Amount float64
}
