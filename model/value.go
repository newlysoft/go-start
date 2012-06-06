package model

type Value interface {
	String() string
	SetString(str string) error
	IsEmpty() bool
	Required(metaData *MetaData) bool
	Validator
}
