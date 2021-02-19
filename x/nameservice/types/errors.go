package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "Name does not exist")
)
