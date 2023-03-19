package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var jsonFile *string = nil

func init() {
	// jsonFile = flag.String("env", "dev", "a string")
	// flag.Parse()
}
func LoadConfig() map[string]interface{} {
	file, err := os.Open("local.json")
	//file, err := os.Open(*jsonFile + ".json")
	if err != nil {
		return nil
	}
	defer file.Close()
	var envVar map[string]interface{}
	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &envVar)
	return envVar

}
