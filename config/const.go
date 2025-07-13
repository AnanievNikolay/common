package config

import "errors"

var (
	ErrorConfigPathIsEmpty = errors.New("config path is empty")
	ErrorConfigIsNotExist  = errors.New("config file is not exist")
)

const (
	errTemplate = "failed to parse config file %s"
)
