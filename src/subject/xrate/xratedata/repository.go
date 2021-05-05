package xratedata

import (
	"fmt"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain"
)

type exchangeRateRepository struct {}

func New() xratedomain.RepositoryInterface {
	return &exchangeRateRepository{}
}

func (r *exchangeRateRepository) GetFromDatabase() error {
	fmt.Println("Getting from database")
	return nil
}