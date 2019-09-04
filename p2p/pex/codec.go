package pex

var cdc *amino.Codec = amino.NewCodec()

func init() {
	RegisterPexMessage(cdc)
}
