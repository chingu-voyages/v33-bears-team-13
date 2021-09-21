package service

import (
	"context"

	"github.com/chingu-voyages/v33-bears-team-13/backend/repository"
)

type SummaryService interface {
	Save(context.Context, string) error
	FindAll(ct context.Context) ([]string, error)
}


type summaryService struct {
	repository repository.RemoteRecord
}

func New(summaryRepository repository.RemoteRecord) SummaryService {
	return &summaryService{
		repository: summaryRepository,
	}
}

func (service *summaryService) Save(ct context.Context, summary string) error {
	service.repository.AddSummary(ct, summary)
	return nil
}

func (service *summaryService) FindAll(ct context.Context) ([]string, error) {
	return service.repository.Read(ct)
}

