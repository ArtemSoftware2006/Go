package main

import "time"

func QuizRunner(questions, answers []string, answerCh chan string) int {
	counter := 0
	for i := 0; i < len(questions); i++ {
		select {
		case <-time.After(time.Second):
			continue
		case res := <-answerCh:
			if res == answers[i] {
				counter++
			}
		}
	}
	return counter
}
