package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const typeMsgAddMarket = "add_market"

var _ sdk.Msg = &MsgAddMarket{}

// NewMsgAddMarket creates the new input for adding an market to blockchain
func NewMsgAddMarket(creator string, ticket string) *MsgAddMarket {
	return &MsgAddMarket{
		Creator: creator,
		Ticket:  ticket,
	}
}

// Route return the message route for slashing
func (msg *MsgAddMarket) Route() string {
	return RouterKey
}

// Type returns the msg add market type
func (msg *MsgAddMarket) Type() string {
	return typeMsgAddMarket
}

// GetSigners return the creators address
func (msg *MsgAddMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes return the marshalled bytes of the msg
func (msg *MsgAddMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the input creation market
func (msg *MsgAddMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Ticket == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ticket param")
	}
	return nil
}

// typeMsgUpdateMarket is the market name of update market
const typeMsgUpdateMarket = "update_market"

var _ sdk.Msg = &MsgUpdateMarket{}

// NewMsgUpdateMarket accepts the params to create new update body
func NewMsgUpdateMarket(creator, ticket string) *MsgUpdateMarket {
	return &MsgUpdateMarket{
		Creator: creator,
		Ticket:  ticket,
	}
}

// Route return the message route for slashing
func (msg *MsgUpdateMarket) Route() string {
	return RouterKey
}

// Type return the update market type
func (msg *MsgUpdateMarket) Type() string {
	return typeMsgUpdateMarket
}

// GetSigners return the creators address
func (msg *MsgUpdateMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes return the marshalled bytes of the msg
func (msg *MsgUpdateMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the input update market
func (msg *MsgUpdateMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Ticket == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ticket param")
	}

	return nil
}
