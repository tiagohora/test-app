package xratedelivery

import (
	"context"
	"fmt"
	"github.com/tiagohora/test-app/src/infra/dummy"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain/xratemodel"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain/xrateservice"
)

type ExchangeRateConsumer interface {
	Handle(c context.Context, r dummy.ResponseWriter)
}

type exchangeRateConsumer struct {
	s xrateservice.ExchangeRateService
}

func New(service xrateservice.ExchangeRateService) ExchangeRateConsumer {
	return &exchangeRateConsumer{s: service}

}

func (co *exchangeRateConsumer) Handle(c context.Context, r dummy.ResponseWriter) {
	model := xratemodel.ExchangeRate{}
	if err := co.s.DoStuff(model); err != nil {
		fmt.Printf("Consumer error: %v", err)
	}

	r.WriteString("Hello World")

}