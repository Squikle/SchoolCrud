package DAL

type Car struct {
	Id     int     `db:"Id"`
	Brand  *string `db:"Brand"`
	Model  *string `db:"Model"`
	Number *string `db:"Number"`
	Year   int     `db:"Year"`
	Color  *string `db:"Color"`
}

func (c *Car) GetId() int {
	return c.Id
}
