package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelexported "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

// DefaultPacketTimeout is the default packet timeout relative to the current block height
const (
	DefaultPacketTimeout = 1000 // NOTE: in blocks
)

// Keeper defines the IBC transfer keeper
type Keeper struct {
	storeKey          sdk.StoreKey
	cdc               *codec.Codec
	boundedCapability sdk.CapabilityKey

	clientKeeper     types.ClientKeeper
	connectionKeeper types.ConnectionKeeper
	channelKeeper    types.ChannelKeeper
	bankKeeper       types.BankKeeper
	supplyKeeper     types.SupplyKeeper
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, capKey sdk.CapabilityKey,
	channelKeeper types.ChannelKeeper,
	bankKeeper types.BankKeeper, supplyKeeper types.SupplyKeeper,
) Keeper {

	// ensure ibc transfer module account is set
	if addr := supplyKeeper.GetModuleAddress(types.GetModuleAccountName()); addr == nil {
		panic("the IBC transfer module account has not been set")
	}

	return Keeper{
		storeKey:          key,
		cdc:               cdc,
		boundedCapability: capKey,
		channelKeeper:     channelKeeper,
		bankKeeper:        bankKeeper,
		supplyKeeper:      supplyKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s/%s", ibctypes.ModuleName, types.SubModuleName))
}

// GetTransferAccount returns the ICS20 - transfers ModuleAccount
func (k Keeper) GetTransferAccount(ctx sdk.Context) supplyexported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.GetModuleAccountName())
}

func (k Keeper) PacketExecuted(ctx sdk.Context, packet channelexported.PacketI, acknowledgement channelexported.PacketDataI) error {
	return k.channelKeeper.PacketExecuted(ctx, packet, acknowledgement)
}
