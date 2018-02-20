package helper

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Parameter struct {
	Params       map[string]string
	Limit        int
	Page         int
	Order        string
	TotalRecords int
	TotalPages   int64
}

const (
	defaultLimit = "25"
	defaultOrder = "asc"
	defaultPage  = "1"
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
	self.Order = c.DefaultQuery("order", defaultOrder)

	page, err := validate(c.DefaultQuery("page", defaultPage))
	if err != nil {
		return err
	}
	self.Page = int(math.Max(1, float64(page)))

	self.Params = make(map[string]string)
	c.Request.ParseForm()
	for key, _ := range c.Request.Form {
		if !StringInSlice(key, []string{"limit", "page", "order", "total_pages", "total_records"}) {
			self.Params[key] = c.Query(key)
		}
	}

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
