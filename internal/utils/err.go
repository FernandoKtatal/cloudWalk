package utils

import "errors"

var (
	FileNotFound   = errors.New("file not found")
	ParseKillLine  = errors.New("error parsing kill line")
	PlayerNotFound = errors.New("player not found")
)
