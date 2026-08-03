package calc_grades

import "fmt"

func FinalNotes(students []string, examNotes []float64, homeworkNotes []float64) map[string]string {
	finalNotes := make(map[string]string)
	for index, student := range students {
		finalNotes[student] = FinalNote(examNotes[index], homeworkNotes[index])
	}
	return finalNotes
}

func FinalNote(exam, homework float64) string {
	examNote := exam * 0.7
	homeworkNote := homework * 0.3
	finalNote := fmt.Sprintf("%.2f", examNote+homeworkNote)
	return finalNote
}

func Passed(grade float64) bool {
	if grade < 6.0 {
		return false
	}
	return true
}
