package coldrepo_test

import (
	"testing"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/address"
	"github.com/hiromaily/go-bitcoin/pkg/testutil"
)

//NewAccountKeyRepository

// TestTx is test for any data operation
func TestAccountKey(t *testing.T) {
	//boil.DebugMode = true
	akRepo := testutil.NewAccountKeyRepository()

	_, err := akRepo.UpdateAddrStatus(
		account.AccountTypeClient,
		address.AddrStatusAddressExported,
		[]string{"cND89yJJxG6KHxZrR7ZrwqQ3yrFSUjDRrGnHKiHFxNjJYmqUQRBu"})
	if err != nil {
		t.Fatalf("fail to call UpdateAddrStatus() %v", err)
	}
}