package homework5

import (
	"fmt"
)

func ConductExam() {
	students := PrepareStudents() // Вызов PrepareStudents прямо здесь

	for i := range students {
		student := &students[i]
		fmt.Printf("%s пришел на экзамен.\n", student.Name)
		score := 50

		if student.Attendance {
			score += 25
		}

		if student.PreparedNotes {
			score += 25
		}

		switch {
		case score > 75:
			student.Grade = "отлично"
		case score > 50:
			student.Grade = "хорошо"
		default:
			student.Grade = "удовлетворительно"
		}

		fmt.Printf("%s получил оценку '%s'.\n", student.Name, student.Grade)
	}
}
