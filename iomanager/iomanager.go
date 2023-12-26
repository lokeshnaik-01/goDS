package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJson(data interface{}) error
	WriteResult(data interface{}) error
}

