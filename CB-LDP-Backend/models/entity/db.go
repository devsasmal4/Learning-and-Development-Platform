package entity

import (
	"cb-ldp-backend/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Module struct {
	Id                   primitive.ObjectID `bson:"_id" json:"_id"`
	ModuleName           string             `bson:"module_name" json:"module_name" binding:"required,max=100"`
	ModuleStudyMaterial  string             `bson:"module_study_material,omitempty" json:"module_study_material,omitempty" binding:"required,max=200,startswith=https://drive.google.com/"`
	ModuleCreatedBy      string             `bson:"module_created_by,omitempty" json:"module_created_by" binding:"required"`
	ModuleDepartment     string             `bson:"module_department,omitempty" json:"module_department" binding:"required"`
	ModuleInstructions   []string           `bson:"module_instructions,omitempty" json:"module_instructions" binding:"required,min=1,max=20,unique"`
	ModuleDescription    string             `bson:"module_description,omitempty" json:"module_description" binding:"required"`
	ModuleStatus         bool               `bson:"module_status,omitempty" json:"module_status"`
	ModuleCreatedOn      time.Time          `bson:"module_created_on" json:"module_created_on" binding:"required"`
	ModuleCompletionDate time.Time          `bson:"module_completion_date" json:"module_completion_date"`
	ModuleQuiz           Quiz               `bson:"module_quiz" json:"module_quiz" binding:"dive,required"`
}

type Quiz struct {
	Id                  primitive.ObjectID `bson:"_id" json:"_id"`
	ModuleId            primitive.ObjectID `bson:"module_id" json:"module_id"`
	QuizName            string             `bson:"quiz_name" json:"quiz_name,omitempty" binding:"max=100"`
	QuizPassingMarks    int                `bson:"quiz_passing_marks" json:"quiz_passing_marks" binding:"min=1,max=1000"`
	QuizDurationMinutes int                `bson:"quiz_duration_minutes" json:"quiz_duration_minutes" binding:"min=-1" default:"-1"`
	QuizTotalMarks      int                `bson:"quiz_total_marks" json:"quiz_total_marks" binding:"required,min=1,max=1000"`
	QuizQuestionaires   []Question         `bson:"quiz_questionaires" json:"quiz_questionaires" binding:"dive,required,min=1,max=50,unique"`
}

type Question struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id"`
	QuizId            primitive.ObjectID `bson:"quiz_id" json:"quiz_id"`
	QuestionText      string             `bson:"question_text" json:"question_text"  binding:"required,max=200"`
	AnswerExplanation string             `bson:"answer_explanation" json:"answer_explanation,omitempty" binding:"min=10,max=1000"`
	QuestionOptions   []Option           `bson:"question_options" json:"question_options" binding:"dive,required,min=2,max=10,unique"`
}

type Option struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	QuestionId   primitive.ObjectID `bson:"question_id" json:"question_id"`
	OptionString string             `bson:"option_string" json:"option_string" binding:"min=1,max=250"`
	OptionScore  int                `bson:"option_score" json:"option_score,omitempty" binding:"min=-10"`
	IsCorrect    bool               `bson:"is_correct" json:"is_correct,omitempty"`
}

type QuizResponse struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	UserId     primitive.ObjectID `bson:"user_id" json:"user_id"`
	ModuleId   primitive.ObjectID `bson:"module_id" json:"module_id" binding:"required"`
	StartDate  time.Time          `bson:"start_date" json:"start_date"`
	EndDate    time.Time          `bson:"end_date" json:"end_date"`
	QuizScore  int                `bson:"quiz_score" json:"quiz_score" binding:"required"`
	QuizResult bool               `bson:"quiz_result" json:"quiz_result"`
}

type User struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	UserMail     string             `bson:"user_mail" json:"user_mail" binding:"min=1,max=50,email,endswith=coffeebeans.io"`
	UserName     string             `bson:"user_name" json:"user_name" binding:"min=1,max=50"`
	UserRole     constants.UserRole `bson:"user_role" json:"user_role"`
	EmployeeId   string             `bson:"employee_id" json:"employee_id"`
	ZohoId       int64              `bson:"zoho_id" json:"zoho_id"`
	TimeLoggedIn int64              `bson:"time_logged_in" json:"time_logged_in"`
}
