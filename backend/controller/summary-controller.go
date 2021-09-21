package controller

import (
	"context"

	"github.com/chingu-voyages/v33-bears-team-13/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SummaryController interface {
	FindAll(ct context.Context) ([]string, error)
	Save(ctx *gin.Context, ct context.Context) error

}

type controller struct {
	service service.SummaryService
}

var validate *validator.Validate

func New(service service.SummaryService) SummaryController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ct context.Context) ([]string, error) {
	return c.service.FindAll(ct)
}

func (c *controller) Save(ctx *gin.Context, ct context.Context) error {
    var summary string
	err := ctx.ShouldBindJSON(&summary)
	if err != nil {
		return err
	}
	c.service.Save(ct, summary)
	return nil
}

