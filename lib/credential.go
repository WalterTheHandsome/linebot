package lib

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadCredential(path string) (*ServerCredential, error) {
	ret := new(ServerCredential)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), ret)
	if err != nil {
		fmt.Println("unmarshall entry error")
		return nil, err
	}
	return ret, nil
}
