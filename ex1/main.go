package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func computeGrade(gotPoints, maxPoints float32) (float32, error) {
	if gotPoints > maxPoints {
		return 0.0, errors.New("The got points can't be over the max points.")
	}
	if gotPoints < 0 {
		return 0.0, errors.New("You can't have negative points.")
	}
	return gotPoints/maxPoints*5 + 1, nil
}

type Exam struct {
	gotPoints float32
	maxPoints float32
}

func getRandomExamPoints() float32 {
	return rand.Float32()*250 - 100
}

func main() {
	// main task
	computeGrade(0, 6)    // 1
	computeGrade(9.5, 18) // 3.6388888
	computeGrade(10, 10)  // 6

	// extra task
	var randomExams []Exam
	for i := 0; i < 20; i++ {
		randomExams = append(randomExams, Exam{getRandomExamPoints(), getRandomExamPoints()})
	}
	for _, randomExam := range randomExams {
		grade, err := computeGrade(randomExam.gotPoints, randomExam.maxPoints)
		if err != nil {
			fmt.Fprintf(os.Stderr, "compute grade of gotpoints %v, maxpoints %v: %v\n", randomExam.gotPoints, randomExam.maxPoints, err)
		} else {
			fmt.Printf("If a student gets %v out of %v points, he or she will receive a %v.\n", randomExam.gotPoints, randomExam.maxPoints, grade)
		}
	}
}
