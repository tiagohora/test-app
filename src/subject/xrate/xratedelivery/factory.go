package xratedelivery

import (
	"github.com/tiagohora/test-app/src/subject/xrate/xratedata"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedomain/xrateservice"
)

var consumer ExchangeRateConsumer

func Factory() ExchangeRateConsumer {
	if consumer == nil {
		service := xrateservice.New(xratedata.New())
		consumer = New(service)
	}

	return consumer
}
