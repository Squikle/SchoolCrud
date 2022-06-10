package DAL

type School struct {
	Id          int     `db:"Id"`
	Address     *string `db:"Address"`
	CoursePrice float64 `db:"CoursePrice"`
}

func (s *School) GetId() int {
	return s.Id
}
