package state_test

import (
	"os"
	"testing"

	"github.com/noah-blockchain/noah-node/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
