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
		Add(origin.NewExchange("balancerV2").WithSymbol("STETH/ETH").
			WithCustom("rate", "0x0000000000000000000000000000000000000000000000000EF976AF325D68E80000000000000000000000000000000000000000000000000000000000002A300000000000000000000000000000000000000000000000000000000062D81469").
			WithCustom("price", "0x0000000000000000000000000000000000000000000000000d925d70884a3395")).
		Add(origin.NewExchange("curve").WithSymbol("STETH/ETH").WithPrice(1.044)).
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(2339)).
		Add(origin.NewExchange("ftx").WithSymbol("ETH/USD").WithPrice(2331)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/USD").WithPrice(2339)).
		Add(origin.NewExchange("gemini").WithSymbol("ETH/USD").WithPrice(2340)).
		Add(origin.NewExchange("kraken").WithSymbol("ETH/USD").WithPrice(2338)).
		Add(origin.NewExchange("uniswap_v3").WithPrice(2338)).
		Add(origin.NewExchange("ethrpc")).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err := callSetzer("price", "wstethusd")
	s.Require().NoError(err)
	s.Require().Equal(dropLastDigits("2480.0612649164", 1), dropLastDigits(out, 1))
}
