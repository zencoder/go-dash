package testfixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadJSONFixture(path string, uo interface{}) (js string, fjs string) {
	// Load in the file and store in a string
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("LoadJSONFixture Error. ioutil.ReadFile. path = %s, Err = %s", path, err.Error()))
	}
	js = string(f)

	// Generate the flat(compacted) JSON
	cb := new(bytes.Buffer)
	err = json.Compact(cb, f)
	if err != nil {
		panic(fmt.Sprintf("LoadJSONFixture Error. json.Compact. path = %s, Err = %s", path, err.Error()))
	}
	fjs = cb.String()

	// Unmarshal the JSON to an object
	err = json.Unmarshal(f, &uo)
	if err != nil {
		panic(fmt.Sprintf("LoadJSONFixture Error. json.Unmarshal. path = %s, Err = %s", path, err.Error()))
	}

	return
}

// Load test fixture from path relative to fixtures directory
func LoadFixture(path string) (js string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("LoadFixture Error. ioutil.ReadFile. path = %s, Err = %s", path, err.Error()))
	}
	js = string(f)
	return
}
