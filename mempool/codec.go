package mempool

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}
