package conn

import (
	amino "github.com/tendermint/go-amino"
	cryptoamino "github.com/kava-labs/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
