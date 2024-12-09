package entity

type User struct {
	ID      int
	Balance Balance
	Name    string
}

type Balance struct {
	Numbers float64
}

type ChangeBalance struct {
	Amount float64
}

type CreateUser struct {
	Name string `json:"name"`
}
