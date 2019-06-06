package model

type ModelInterface interface {
	Find(id int) Model
}
