package utils

import (
	"os"
)

func ShowResolve() (string, error) {
	path := "/etc/resolv.conf"

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
