package kyc

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgWhitelist{}, "kyc/whitelist", nil)
	cdc.RegisterConcrete(MsgRevokeWhitelist{}, "kyc/revokeWhitelist", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
	codec.RegisterCrypto(msgCdc)
}
