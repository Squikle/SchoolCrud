package DAL

type Person struct {
	Id       int     `db:"Id"`
	FistName *string `db:"FistName"`
	LastName *string `db:"LastName"`
	Phone    *string `db:"Phone"`
	Email    *string `db:"Email"`
	Sex      *string `db:"Sex"`
}

func (p *Person) GetId() int {
	return p.Id
}
