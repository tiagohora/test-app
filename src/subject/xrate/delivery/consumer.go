package xratedelivery

import (
	"context"
	"fmt"
	"github.com/tiagohora/test-app/src/infra/dummy"
	"github.com/tiagohora/test-app/src/subject/xrate/domain/model"
	"github.com/tiagohora/test-app/src/subject/xrate/domain/service"
)

type ExchangeRateConsumer interface {
	Handle(c context.Context, r dummy.ResponseWriter)
}

type exchangeRateConsumer struct {
	s service.ExchangeRateService
}

func New(service service.ExchangeRateService) ExchangeRateConsumer {
	return &exchangeRateConsumer{s: service}

}

func (co *exchangeRateConsumer) Handle(c context.Context, r dummy.ResponseWriter) {
	model := model.ExchangeRate{}
	if err := co.s.DoStuff(model); err != nil {
		fmt.Printf("Consumer error: %v", err)
	}

	r.WriteString("Hello World")

}