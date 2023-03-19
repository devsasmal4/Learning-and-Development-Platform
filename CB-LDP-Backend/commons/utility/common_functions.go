package utility

import (
	"cb-ldp-backend/models/response"
	"math/rand"
	"strings"
	"time"
)

func GetDurationAndStatus(startTime time.Time, endTime time.Time) (string, string) {
	duration := (endTime.Sub(startTime) / time.Minute).String()
	status := "Completed"
	if endTime.IsZero() {
		status = "Ongoing"
		duration = "0"
		if startTime.IsZero() {
			status = "Not yet started"
		}
	}
	duration = strings.TrimSuffix(duration, "ns")
	duration = strings.TrimSuffix(duration, "s")
	return duration, status
}

func Contains(slice []string, exp string) bool {
	for _, a := range slice {
		if strings.EqualFold(a, exp) {
			return true
		}
	}
	return false
}

func Shuffle(testDetails response.QuizQuestions) {
	// shuffle options for each question
	rand.Seed(time.Now().UnixNano())
	for _, q := range testDetails.QuizQuestionaires {
		qo := q.QuestionOptions
		rand.Shuffle(len(qo), func(i, j int) { qo[i], qo[j] = qo[j], qo[i] })
	}

	//shuffle questions for each quiz
	questions := testDetails.QuizQuestionaires
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
