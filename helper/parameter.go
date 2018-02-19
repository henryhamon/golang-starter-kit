package helper

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

const (
	defaultLimit = "25"
)

func NewParameter(c *gin.Context) (*Parameter, error) {
	parameter := &Parameter{}

	if err := parameter.initialize(c); err != nil {
		return nil, err
	}

	return parameter, nil
}

func (self *Parameter) initialize(c *gin.Context) error {

	limit, err := validate(c.DefaultQuery("limit", defaultLimit))
	if err != nil {
		return err
	}

	self.Limit = int(math.Max(1, math.Min(10000, float64(limit))))

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
