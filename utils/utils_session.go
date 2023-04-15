package utils

import (
	"fmt"

	"gopkg.in/boj/redistore.v1"
)

func RediStore(size int, network string, host string, port string, password string, secretKey []byte) (*redistore.RediStore, error) {
	rdStore, err := redistore.NewRediStore(size, network, fmt.Sprintf("%s:%s", host, port), password, secretKey)
	if err != nil {
		return nil, err
	}

	return rdStore, nil
}
