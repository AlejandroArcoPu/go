package main

import (
	"fmt"
	"log"

	"example.com/calc_grades"
)

func main() {
	log.SetPrefix("students: ")
	log.SetFlags(0)

	finalNote := 5.0

	fmt.Println(calc_grades.Passed(finalNote))

	examNote := 6.0
	homeworkNote := 3.0

	fmt.Println(calc_grades.FinalNote(examNote, homeworkNote))

	students := []string{"Alex", "Pedro", "Mario"}
	examNotes := []float64{9.2, 8.2, 9.3}
	homeworkNotes := []float64{8.5, 7.3, 9.0}

	fmt.Println(calc_grades.FinalNotes(students, examNotes, homeworkNotes))
}
