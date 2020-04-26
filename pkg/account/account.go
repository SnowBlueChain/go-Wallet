package account

// AccountType utilization purpose of address
type AccountType string

// account_type
const (
	AccountTypeClient        AccountType = "client"    // for client receipt address
	AccountTypeReceipt       AccountType = "receipt"   // for receipt action
	AccountTypePayment       AccountType = "payment"   // for payment action
	AccountTypeFee           AccountType = "fee"       // for transaction fee
	AccountTypeStored        AccountType = "stored"    // storage for huge amount of coin
	AccountTypeAnonymous     AccountType = "anonymous" // payment user
	AccountTypeAuthorization AccountType = "auth"      // authorization for multi address
	AccountTypeTest          AccountType = "test"      // for unittest
)

// String converter
func (a AccountType) String() string {
	return string(a)
}

// Is compare with params
func (a AccountType) Is(v string) bool {
	return a.String() == v
}

// Allow return true if acnt is in list
func Allow(acnt string, accountList []AccountType) bool {
	if !ValidateAccountType(acnt) {
		return false
	}
	for _, v := range accountList {
		if acnt == v.String() {
			return true
		}
	}
	return false
}

// NotAllow return true if acnt is not in list
func NotAllow(acnt string, accountList []AccountType) bool {
	if !ValidateAccountType(acnt) {
		return false
	}
	for _, v := range accountList {
		if acnt == v.String() {
			return false
		}
	}
	return true
}

// AccountTypeMap account_type mapper
var AccountTypeMap = map[string]AccountType{
	"client":        AccountTypeClient,
	"receipt":       AccountTypeReceipt,
	"payment":       AccountTypePayment,
	"fee":           AccountTypeFee,
	"stored":        AccountTypeStored,
	"authorization": AccountTypeAuthorization,
}

// ValidateAccountType validate AccountType
func ValidateAccountType(v string) bool {
	if _, ok := AccountTypeMap[v]; ok {
		return true
	}
	return false
}

// AccountTypeValue account_type value
var AccountTypeValue = map[AccountType]uint32{
	AccountTypeClient:        0,
	AccountTypeReceipt:       1,
	AccountTypePayment:       2,
	AccountTypeFee:           3,
	AccountTypeStored:        4,
	AccountTypeAuthorization: 5,
}

// AccountTypeMultisig true: account type is for multisig address
var AccountTypeMultisig = map[AccountType]bool{
	AccountTypeClient:        false,
	AccountTypeReceipt:       true,
	AccountTypePayment:       true,
	AccountTypeFee:           true,
	AccountTypeStored:        true,
	AccountTypeAuthorization: false,
}
