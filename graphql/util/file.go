package graphql

import (
	"io/ioutil"
)

func readSchema(fileName string) (string, error) {
	ext := ".graphql"
	b, err := ioutil.ReadFile(fileName + ext)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
