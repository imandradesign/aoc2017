package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const dataDirEnvVar = "AOC_DATA_DIR"

func loadInputData(day string) ([]string, error) {
	dataDir, ok := os.LookupEnv(dataDirEnvVar)
	if !ok {
		return []string{}, fmt.Errorf("%s not set in your env", dataDirEnvVar)
	}

	loadPath := path.Join(dataDir, fmt.Sprintf("%s.txt", day))
	_, err := os.Stat(loadPath)
	if err != nil {
		return []string{}, err
	}

	bytes, err := ioutil.ReadFile(loadPath)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(bytes), "\n"), nil
}
