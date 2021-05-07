package service

import (
	"github.com/tiagohora/test-app/src/subject/xrate/domain"
	"github.com/tiagohora/test-app/src/subject/xrate/domain/model"
)

type ExchangeRateService interface {
	DoStuff(m model.ExchangeRate) error
}

type exchangeRateService struct {
	r domain.RepositoryInterface
}

func New(repo domain.RepositoryInterface) ExchangeRateService {
	return &exchangeRateService{
		r: repo,
	}
}


func (s *exchangeRateService) DoStuff(m model.ExchangeRate) error {
	return s.r.GetFromDatabase()
}