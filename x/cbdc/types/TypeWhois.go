package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinCBDC = sdk.Coins{sdk.NewInt64Coin("cbdctoken", 1)}

type Whois struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Value   string         `json:"value" yaml:"value"`
	Price   sdk.Coins      `json:"price" yaml:"price"`
}

func NewWhois() Whois {
	return Whois{
		Price: MinCBDC,
	}
}
