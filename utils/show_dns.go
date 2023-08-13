package utils

import (
	"io/ioutil"
)

func ShowResolve() (string, error) {
	path := "/etc/resolv.conf"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
