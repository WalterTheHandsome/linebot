package lib

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadCredential(path string) (*ServerCredential, error) {
	ret := new(ServerCredential)
	data, err := ioutil.ReadFile(path)
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
