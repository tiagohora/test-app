package xrateservice

import (
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain/xratemodel"
)

type ExchangeRateService interface {
	DoStuff(m xratemodel.ExchangeRate) error
}

type exchangeRateService struct {
	r xratedomain.RepositoryInterface
}

func New(repo xratedomain.RepositoryInterface) ExchangeRateService {
	return &exchangeRateService{
		r: repo,
	}
}


func (s *exchangeRateService) DoStuff(m xratemodel.ExchangeRate) error {
	return s.r.GetFromDatabase()
}