package app

type Person struct {
	Id        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
}
