package helper

import "strconv"

type Parameter struct {
	Params       map[string]string
	Sort         string
	Limit        int
	Page         int
	LastID       int
	Order        string
	IsLastID     bool
	TotalRecords int
	TotalPages   int64
}

func NewParameter() (*Parameter, error) {
	parameter := &Parameter{}

	if err := parameter.initialize(); err != nil {
		return nil, err
	}

	return parameter, nil
}

func (self *Parameter) initialize() error {
	self.Limit = 25
	self.Page = 1
	self.Order = "asc"

	return nil
}

func validate(s string) (int, error) {
	if s == "" {
		return -1, nil
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}
