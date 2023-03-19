package response

import "cb-ldp-backend/models/entity"

type QuizQuestions struct {
	ModuleName        string            `json:"module_name"`
	QuizQuestionaires []entity.Question `json:"quiz_questionaires"`
	QuizName          string            `json:"quiz_name,omitempty"`
}

type AnswerResponse struct {
	ScoreEarned       int    `json:"score_earned"`
	AnswerExplanation string `json:"answer_explanation,omitempty"`
	CorrectAnswer     string `json:"correct_answer"`
}

type QuestionJsonResponse struct {
	TotalScore int                `json:"total_marks"`
	Questions  []QuestionResponse `json:"questions" validate:"dive,required"`
}

type QuestionResponse struct {
	QuestionText      string           `json:"question_text" validate:"required,max=200"`
	AnswerExplanation string           `json:"answer_explanation" validate:"min=10,max=1000"`
	QuestionOptions   []OptionResponse `json:"question_options" validate:"dive,required"`
}

type OptionResponse struct {
	OptionString string `json:"option_string" validate:"min=1,max=250"`
	OptionScore  int    `json:"option_score" validate:"min=-10"`
	IsCorrect    bool   `json:"is_correct"`
}
