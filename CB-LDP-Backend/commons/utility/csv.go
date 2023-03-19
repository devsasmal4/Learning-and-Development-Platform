package utility

import (
	"cb-ldp-backend/config"
	"cb-ldp-backend/models/response"
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/twinj/uuid"
)

var envVar = config.LoadConfig()

func GenerateCsv(data []response.TestDetailsResponse) (string, error) {
	dirPath := envVar["dir_path"].(string) + "downloads/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", err
	}

	fileName := dirPath + uuid.NewV4().String() + ".csv"
	csvFile, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer csvFile.Close()
	column := []string{
		"UserName", "ModuleName", "StartDate", "EndDate", "QuizStatus", "QuizScore", "QuizResult", "Duration",
	}
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Write(column)

	for _, csvData := range data {
		row := []string{}
		row = append(row, csvData.UserName)
		row = append(row, csvData.ModuleName)
		row = append(row, csvData.StartDate.String())
		row = append(row, csvData.EndDate.String())
		row = append(row, csvData.TestStatus)
		row = append(row, strconv.Itoa(int(csvData.QuizScore)))
		if csvData.QuizResult {
			row = append(row, "Pass")
		} else {
			row = append(row, "Fail")
		}
		row = append(row, csvData.Duration)
		csvWriter.Write(row)
	}
	csvWriter.Flush()
	return fileName, nil
}

func ConvertCsvToJson(fileName string) (interface{}, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvData, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	csvData = csvData[1:]
	var questions []response.QuestionResponse

	var totalScore int
	for _, data := range csvData {
		var question response.QuestionResponse
		questionCount, _ := strconv.Atoi(data[0])
		if questionCount < 1 || questionCount > 50 {
			return nil, errors.New("Should have 1 to 50 questions")
		}
		question.QuestionText = data[1]
		question.AnswerExplanation = data[2]
		optionCount, _ := strconv.Atoi(data[3])
		if optionCount < 2 || optionCount > 10 {
			return nil, errors.New("Should have 2 to 10 options")
		}
		var option response.OptionResponse
		for i := 0; i < optionCount; i++ {
			option.OptionString = data[(2*i)+4]
			option.OptionScore, _ = strconv.Atoi(data[(2*i)+5])
			if i == 0 {
				totalScore += option.OptionScore
				option.IsCorrect = true
			} else {
				option.IsCorrect = false
			}
			question.QuestionOptions = append(question.QuestionOptions, option)
		}
		questions = append(questions, question)
	}

	questionJsonResponse := response.QuestionJsonResponse{TotalScore: totalScore, Questions: questions}
	var validResponse response.ValidResponse = questionJsonResponse
	if err := validResponse.ValidateStruct(); err != nil {
		return nil, err
	}
	return questionJsonResponse, nil
}
