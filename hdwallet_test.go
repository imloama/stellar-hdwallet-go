package hdwallet

import "testing"

func TestFromMnemonic(t *testing.T) {
	mnemonic := "illness spike retreat truth genius clock brain pass fit cave bargain toe"
	FromMnemonic(mnemonic)
}
