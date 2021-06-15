package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func GeneratePaginationFromRequest(c *gin.Context) (Pagination, error) {
	//*Initializing default
	limit := 2
	page := 1

	query := c.Request.URL.Query()

	for key, value := range query {
		var err error

		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, err = strconv.Atoi(queryValue)
			if err != nil {
				return Pagination{}, err
			}
		case "page":
			page, err = strconv.Atoi(queryValue)
			if err != nil {
				return Pagination{}, err
			}
		}
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}, nil

}

func CountTotalPage(totalData int, limit *Pagination) int {
	if totalData == 0 {
		return 1
	}
	return int(math.Ceil(float64(totalData) / float64(limit.Limit)))
}
