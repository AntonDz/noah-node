package store

import (
	"github.com/noah-blockchain/noah-node/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
