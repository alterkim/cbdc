package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgPayment defines the payment message
type MsgPayment struct {
	Price  sdk.Coins      `json:"price"`
	Payer  sdk.AccAddress `json:"payer`
	Seller sdk.AccAddress `json:"seller`
}

func NewMsgPayment(price sdk.Coins, payer sdk.AccAddress, seller sdk.AccAddress) MsgPayment {
	return MsgPayment{
		Price:  price,
		Payer:  payer,
		Seller: seller,
	}
}

// Route should return the name of the module
func (msg MsgPayment) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPayment) Type() string { return "payment" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPayment) ValidateBasic() error {
	if msg.Payer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Payer.String())
	}
	if msg.Seller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Seller.String())
	}
	if !msg.Price.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Payer}
}
