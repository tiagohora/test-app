//+build wireinject

package xratedelivery

import (
	"github.com/google/wire"
	"github.com/tiagohora/test-app/src/subject/xrate/data"
	"github.com/tiagohora/test-app/src/subject/xrate/domain/service"
)

func Factory() ExchangeRateConsumer {
	 wire.Build(New, service.New, data.New)
	 return &exchangeRateConsumer{}
}
