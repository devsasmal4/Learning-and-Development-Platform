package mocks

import (
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MockModule() entity.Module {
	mockId, _ := primitive.ObjectIDFromHex("635fb6560fc5f5c7d38976ec")
	instrctions := [...]string{"Instruction 1", "Instruction 2", "Instruction 3"}
	return entity.Module{
		Id:                   mockId,
		ModuleName:           "Module 1",
		ModuleStudyMaterial:  "https://drive.google.com/drive/folders/15MOf2M5WHOakvBONv_iWdlvOasqcgQXS?usp=sharing",
		ModuleCreatedBy:      "User 1",
		ModuleDepartment:     "Department 1",
		ModuleInstructions:   instrctions[:],
		ModuleDescription:    "Description 1",
		ModuleStatus:         true,
		ModuleCreatedOn:      time.Now(),
		ModuleCompletionDate: time.Now(),
		ModuleQuiz: entity.Quiz{
			QuizName:            "name1",
			QuizPassingMarks:    50,
			QuizDurationMinutes: 90,
			QuizTotalMarks:      100,
			QuizQuestionaires: []entity.Question{
				{
					QuestionText:      "question1",
					AnswerExplanation: "explaination 1 min 10",
					QuestionOptions: []entity.Option{
						{
							OptionString: "option1",
							OptionScore:  23,
							IsCorrect:    true,
						},
					},
				},
			},
		},
	}
}

func MockUser() entity.User {
	mockId, _ := primitive.ObjectIDFromHex("635fb6560fc5f5c7d38976ec")
	return entity.User{
		Id:           mockId,
		UserMail:     "user@coffeebeans.io",
		UserName:     "user",
		UserRole:     "Admin",
		TimeLoggedIn: time.Now().Unix(),
	}
}

func MockQuizResponse() entity.QuizResponse {
	mockId, _ := primitive.ObjectIDFromHex("635fb6560fc5f5c7d38976ec")
	return entity.QuizResponse{
		Id:         mockId,
		UserId:     mockId,
		ModuleId:   mockId,
		StartDate:  time.Time{},
		EndDate:    time.Time{}.Add(3),
		QuizScore:  70,
		QuizResult: true,
	}
}

func MockQuestion() entity.Question {
	mockId, _ := primitive.ObjectIDFromHex("635fb6560fc5f5c7d38976ec")
	return entity.Question{
		Id:                mockId,
		QuizId:            mockId,
		QuestionText:      "Question1",
		AnswerExplanation: "Answer2121",
		QuestionOptions: []entity.Option{
			{
				OptionString: "option1",
				OptionScore:  23,
				IsCorrect:    true,
			},
		},
	}
}

func MockOption() entity.Option {
	mockId, _ := primitive.ObjectIDFromHex("635fb6560fc5f5c7d38976ec")
	return entity.Option{
		Id:           mockId,
		QuestionId:   mockId,
		OptionString: "option1",
		OptionScore:  23,
		IsCorrect:    true,
	}
}

func MockTestDetailsResponse() []response.TestDetailsResponse{
	return []response.TestDetailsResponse{
		{
			Id:         MockModule().Id,
			UserName:   MockUser().UserName,
			ModuleName: MockModule().ModuleName,
			StartDate:  MockQuizResponse().StartDate,
			EndDate:    MockQuizResponse().EndDate,
			TestStatus: "Completed",
			QuizScore:  MockQuizResponse().QuizScore,
			QuizResult: MockQuizResponse().QuizResult,
			Duration:   (MockQuizResponse().EndDate.Sub(MockQuizResponse().StartDate)).String(),
		},
	}
}
