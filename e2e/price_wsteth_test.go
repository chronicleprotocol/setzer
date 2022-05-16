package e2e

import (
	"testing"

	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"

	"github.com/stretchr/testify/suite"
)

func TestPriceWSTETH2ESuite(t *testing.T) {
	suite.Run(t, new(PriceWSTETHE2ESuite))
}

type PriceWSTETHE2ESuite struct {
	SmockerAPISuite
}

func (s *PriceWSTETHE2ESuite) TestPrice() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("wsteth").WithSymbol("WSTETH/ETH").WithPrice(1.062334)).
		Add(origin.NewExchange("balancerV2").WithSymbol("STETH/ETH").WithPrice(1.0573)).
		Add(origin.NewExchange("curve").WithSymbol("STETH/ETH").WithPrice(1.044)).
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(2339)).
		Add(origin.NewExchange("ftx").WithSymbol("ETH/USD").WithPrice(2331)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/USD").WithPrice(2339)).
		Add(origin.NewExchange("gemini").WithSymbol("ETH/USD").WithPrice(2340)).
		Add(origin.NewExchange("kraken").WithSymbol("ETH/USD").WithPrice(2338)).
		Add(origin.NewExchange("ethrpc")).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err := callSetzer("price", "wstethusd")
	s.Require().NoError(err)
	s.Require().Equal("2492.7824058521", out)
}
