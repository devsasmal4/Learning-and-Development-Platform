package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModuleResponse struct {
	Id              primitive.ObjectID `json:"id"`
	ModuleName      string             `json:"module_name,omitempty"`
	ModuleCreatedOn time.Time          `json:"module_created_on"`
}

type ModuleViewResponse struct {
	Id                  primitive.ObjectID `json:"id"`
	ModuleName          string             ` json:"module_name,omitempty"`
	ModuleStudyMaterial string             ` json:"module_study_material,omitempty"`
	ModuleDescription   string             ` json:"module_description,omitempty"`
}

type ModuleInstructionsResponse struct {
	Id                 primitive.ObjectID `json:"id"`
	ModuleName         string             ` json:"module_name,omitempty"`
	ModuleInstructions []string           ` json:"module_instructions,omitempty"`
}

type TestDetailsResponse struct {
	Id         primitive.ObjectID `json:"id"`
	UserName   string             `json:"user_name,omitempty"`
	ModuleName string             `json:"module_name,omitempty"`
	StartDate  time.Time          `json:"start_date,omitempty"`
	EndDate    time.Time          `json:"end_date,omitempty"`
	TestStatus string             `json:"test_status,omitempty"`
	QuizScore  int                `json:"quiz_score,omitempty"`
	QuizResult bool               `json:"quiz_result"`
	Duration   string             `json:"duration"`
}
