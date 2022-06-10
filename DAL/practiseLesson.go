package DAL

type PracticeLesson struct {
	Id        int `db:"Id"`
	LessonId  int `db:"LessonId"`
	StudentId int `db:"StudentId"`
	CarId     int `db:"CarId"`
}

func (pl *PracticeLesson) GetId() int {
	return pl.Id
}
