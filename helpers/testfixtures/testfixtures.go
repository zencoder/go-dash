package testfixtures

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Load test fixture from path relative to fixtures directory
func LoadFixture(path string) (js string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("LoadFixture Error. ioutil.ReadFile. path = %s, Err = %s", path, err.Error()))
	}
	return string(f)
}

func CompareFixture(t *testing.T, fixturePath string, actualContent string) {
	expectedContent := LoadFixture(fixturePath)
	if os.Getenv("GENERATE_FIXTURES") != "" {
		ioutil.WriteFile(fixturePath, []byte(actualContent), os.ModePerm)
		fmt.Println("Wrote fixture: " + fixturePath)
	} else {
		require.Equal(t, expectedContent, actualContent)
	}
}
