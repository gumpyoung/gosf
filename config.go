package gosf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config map[string]interface{}

func init() {
	App.Config = make(map[string]interface{})
}

// LoadConfig loads a JSON configuration file into the global Config map
func LoadConfig(name string, path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data map[string]interface{}
	json.Unmarshal([]byte(byteValue), &data)

	App.Config[name] = data
}
