package DAL

type TheoryLesson struct {
	Id       int `db:"Id"`
	LessonId int `db:"LessonId"`
	GroupId  int `db:"GroupId"`
}

func (tl *TheoryLesson) GetId() int {
	return tl.Id
}
