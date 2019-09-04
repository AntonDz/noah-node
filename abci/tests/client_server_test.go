package tests

import (
	client2 "github.com/noah-blockchain/noah-node/abci/client"
	"github.com/noah-blockchain/noah-node/abci/example/kvstore"
	server3 "github.com/noah-blockchain/noah-node/abci/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientServerNoAddrPrefix(t *testing.T) {
	addr := "localhost:26658"
	transport := "socket"
	app := kvstore.NewKVStoreApplication()

	server, err := server3.NewServer(addr, transport, app)
	assert.NoError(t, err, "expected no error on NewServer")
	err = server.Start()
	assert.NoError(t, err, "expected no error on server.Start")

	client, err := client2.NewClient(addr, transport, true)
	assert.NoError(t, err, "expected no error on NewClient")
	err = client.Start()
	assert.NoError(t, err, "expected no error on client.Start")
}
