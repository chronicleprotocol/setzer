package e2e

import (
	"testing"

	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"

	"github.com/stretchr/testify/suite"
)

func TestPriceRETHSuite(t *testing.T) {
	suite.Run(t, new(PriceRETHSuite))
}

type PriceRETHSuite struct {
	SmockerAPISuite
}

func (s *PriceRETHSuite) TestPrice() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/USD").WithPrice(1601.5438929600)).
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(1601.9400000000)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/USD").WithPrice(1601.2500000000)).
		Add(origin.NewExchange("ftx").WithSymbol("ETH/USD").WithPrice(1600.6000000000)).
		Add(origin.NewExchange("gemini").WithSymbol("ETH/USD").WithPrice(1602.2900000000)).
		Add(origin.NewExchange("kraken").WithSymbol("ETH/USD").WithPrice(1603.2900000000)).
		Add(origin.NewExchange("uniswap_v3").WithPrice(1595.2011018009)).
		Add(origin.NewExchange("curve").WithSymbol("STETH/ETH").WithPrice(1.044)).
		Add(origin.NewExchange("balancerV2").WithSymbol("STETH/ETH").
			WithCustom("rate", "0x0000000000000000000000000000000000000000000000000EF976AF325D68E80000000000000000000000000000000000000000000000000000000000002A300000000000000000000000000000000000000000000000000000000062D81469").
			WithCustom("price", "0x0000000000000000000000000000000000000000000000000d925d70884a3395")).
		Add(origin.NewExchange("rocketpool")).
		Add(origin.NewExchange("ethrpc")).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err := callSetzer("price", "rethusd")
	s.Require().NoError(err)
	s.Require().Equal(dropLastDigits("1645.4670410364", 1), dropLastDigits(out, 1))
}

func (s *PriceRETHSuite) TestCircuit() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/USD").WithPrice(1601.5438929600)).
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(1601.9400000000)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/USD").WithPrice(1601.2500000000)).
		Add(origin.NewExchange("ftx").WithSymbol("ETH/USD").WithPrice(1600.6000000000)).
		Add(origin.NewExchange("gemini").WithSymbol("ETH/USD").WithPrice(1602.2900000000)).
		Add(origin.NewExchange("kraken").WithSymbol("ETH/USD").WithPrice(1603.2900000000)).
		Add(origin.NewExchange("uniswap_v3").WithPrice(1595.2011018009)).
		Add(origin.NewExchange("curve").WithSymbol("STETH/ETH").WithPrice(1.044)).
		Add(origin.NewExchange("balancerV2").
			WithSymbol("STETH/ETH").
			WithCustom("rate", "0x0000000000000000000000000000000000000000000000000EF976AF325D68E80000000000000000000000000000000000000000000000000000000000002A300000000000000000000000000000000000000000000000000000000062D81469").
			WithCustom("price", "0x0000000000000000000000000000000000000000000000000d925d70884a3395")).
		Add(origin.NewExchange("rocketpool").
			WithCustom("price", "0x0000000000000000000000000000000000000000000000008ac7230489e7ffff")).
		Add(origin.NewExchange("ethrpc")).
		Deploy(s.api)

	s.Require().NoError(err)

	_, _, err = callSetzer("price", "rethusd")
	s.Require().Error(err)
}

func dropLastDigits(s string, n int) string {
	return s[:len(s)-n]
}
