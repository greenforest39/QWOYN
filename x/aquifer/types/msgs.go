package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPutAllocationToken{}

func NewMsgPutAllocationToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgPutAllocationToken {
	return &MsgPutAllocationToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgPutAllocationToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgPutAllocationToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgPutAllocationToken{}

func NewMsgBuyAllocationToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgPutAllocationToken {
	return &MsgPutAllocationToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgBuyAllocationToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgBuyAllocationToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgSetDepositEndTime{}

func NewMsgSetDepositEndTime(sender sdk.AccAddress, endTime uint64) *MsgSetDepositEndTime {
	return &MsgSetDepositEndTime{
		Sender:  sender.String(),
		EndTime: endTime,
	}
}

func (msg MsgSetDepositEndTime) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgSetDepositEndTime) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgInitICA{}

func NewMsgInitICA(sender sdk.AccAddress, endTime uint64) *MsgInitICA {
	return &MsgInitICA{
		Sender: sender.String(),
	}
}

func (msg MsgInitICA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgInitICA) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgExecAddLiquidity{}

func NewMsgExecAddLiquidity(sender sdk.AccAddress, endTime uint64) *MsgExecAddLiquidity {
	return &MsgExecAddLiquidity{
		Sender: sender.String(),
	}
}

func (msg MsgExecAddLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgExecAddLiquidity) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
