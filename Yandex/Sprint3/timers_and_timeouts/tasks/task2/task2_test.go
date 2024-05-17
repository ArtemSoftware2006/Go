package main

import (
	"testing"
)

func TestQuizRunner(t *testing.T) {
	fillAnswers := func(answers []string, answerCh chan string) {
		go func() {
			for i := range answers {
				answerCh <- answers[i]
			}
		}()
	}

	checkAnswers := func(questions []string, answers []string, expectedCorrectAnswers int, answerCh chan string) {
		correctAnswers := QuizRunner(questions, answers, answerCh)
		if correctAnswers != expectedCorrectAnswers {
			t.Errorf("Expected %d correct answers, but got %d", expectedCorrectAnswers, correctAnswers)
		}
	}

	questions := []string{
		"What is 2 + 2 ?",
		"What is 2 * 2 ?",
		"What is 2 - 2 ?",
	}
	faileAnswers := []string{"4", "3", "0"}
	rightAnswers := []string{"4", "4", "0"}

	answerCh := make(chan string)
	defer close(answerCh)
	go fillAnswers(faileAnswers, answerCh)
	checkAnswers(questions, rightAnswers, 2, answerCh)

	go fillAnswers(rightAnswers, answerCh)
	checkAnswers(questions, rightAnswers, 3, answerCh)

	checkAnswers(questions, rightAnswers, 0, answerCh)
}
