package data

import (
	"fmt"
	"github.com/tiagohora/test-app/src/subject/xrate/domain"
)

type exchangeRateRepository struct {}

func New() domain.RepositoryInterface {
	return &exchangeRateRepository{}
}

func (r *exchangeRateRepository) GetFromDatabase() error {
	fmt.Println("Getting from database")
	return nil
}