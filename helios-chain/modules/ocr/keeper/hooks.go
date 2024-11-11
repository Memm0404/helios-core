package keeper

import (
	"helios-core/helios-chain/modules/ocr/types"
)

type OcrHooks interface {
	SetHooks(h types.OcrHooks)
}

// Set the hooks
func (k *Keeper) SetHooks(h types.OcrHooks) {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = h
}
