package DAL

type Payments struct {
	Id        int     `db:"Id"`
	StudentId int     `db:"StudentId"`
	SchoolId  int     `db:"SchoolId"`
	Amount    float64 `db:"Amount"`
}

func (p *Payments) GetId() int {
	return p.Id
}
