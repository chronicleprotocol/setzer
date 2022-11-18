package e2e

import (
	"strings"
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
	s.Require().Equal(toPrecision("1645.821567", 6), toPrecision(out, 6))
}

func (s *PriceRETHSuite) TestCircuit() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/USD").WithPrice(1601.5438929600)).
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(1601.9400000000)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/USD").WithPrice(1601.2500000000)).
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

// toPrecision reduces precision of a string to the given number of decimal places.
// It does not round the number.
func toPrecision(s string, p int) string {
	if i := strings.Index(s, "."); i >= 0 {
		if len(s) <= i+p+1 {
			return s + strings.Repeat("0", i+p+1-len(s))
		}
		return s[:i+p+1]
	}
	return s
}
