package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/SolanaPan/nnb/v1/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
}

func NewKeeper(storeKey storetypes.StoreKey) Keeper {
	return Keeper{storeKey: storeKey}
}

func (k Keeper) SetKYCStatus(ctx sdk.Context, addr sdk.AccAddress, status types.KYCStatus) {
	store := ctx.KVStore(k.storeKey)
	store.Set(addr.Bytes(), []byte{byte(status)})
}

func (k Keeper) GetKYCStatus(ctx sdk.Context, addr sdk.AccAddress) types.KYCStatus {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(addr.Bytes())
	if len(bz) == 0 {
		return types.KYCUnknown
	}
	return types.KYCStatus(bz[0])
}
