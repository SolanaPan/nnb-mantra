package types

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "kyc"
	StoreKey   = ModuleName
)

type KYCStatus int

const (
	KYCUnknown KYCStatus = iota
	KYCPending
	KYCApproved
	KYCRejected
)

type MsgSubmitKYC struct {
	Address sdk.AccAddress
	Info    string // Could be JSON or a reference to off-chain data
}

type MsgApproveKYC struct {
	Address  sdk.AccAddress
	Approver sdk.AccAddress
}

type MsgRejectKYC struct {
	Address  sdk.AccAddress
	Approver sdk.AccAddress
	Reason   string
}

type MsgUpdateKYC struct {
	Address sdk.AccAddress
	Info    string
}

func (msg MsgSubmitKYC) Route() string { return ModuleName }
func (msg MsgSubmitKYC) Type() string  { return "submit_kyc" }
func (msg MsgSubmitKYC) ValidateBasic() error {
	if msg.Address.Empty() {
		return fmt.Errorf("address cannot be empty")
	}
	return nil
}

// ...similarly for other Msg types...

// MsgServer interface for KYC module
type MsgServer interface {
	SubmitKYC(context.Context, *MsgSubmitKYC) (*MsgSubmitKYCResponse, error)
	ApproveKYC(context.Context, *MsgApproveKYC) (*MsgApproveKYCResponse, error)
	RejectKYC(context.Context, *MsgRejectKYC) (*MsgRejectKYCResponse, error)
}

type MsgSubmitKYCResponse struct{}
type MsgApproveKYCResponse struct{}
type MsgRejectKYCResponse struct{}
