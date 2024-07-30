package entity

type User struct {
	ID      int
	Balance Balance
}

type Balance struct {
	Numbers float64 //трёхзначные числа на балансе
}

type ChangeBalance struct {
	Amount float64
}
