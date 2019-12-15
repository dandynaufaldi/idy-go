package model

type IDProvider interface {
	Generate() string
}
