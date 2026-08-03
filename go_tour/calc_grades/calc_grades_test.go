package calc_grades

import "testing"

func TestFinalNote(t *testing.T) {
	examNote := 5.0
	finalNote := 5.0
	result := FinalNote(examNote, finalNote)

	if result != "5.00" {
		t.Errorf("FinalNote(5.0, 5.0) = %v, should be 5.00", result)
	}
}

func TestPassedFail(t *testing.T) {
	note := 5.0
	result := Passed(note)
	if result {
		t.Errorf("Passed(5.0) = %t, shouldn't be passed", result)
	}
}

func TestPassedPassed(t *testing.T) {
	note := 9.0
	result := Passed(note)
	if !result {
		t.Errorf("Passed(9.0) = %t, shouldn't be not passed", result)
	}
}
