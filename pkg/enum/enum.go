package enum

//----------------------------------------------------
// CoinType
//----------------------------------------------------

//CoinType Bitcoin種別(CayenneWalletで取引するcoinの種別)
type CoinType string

// coin_type
const (
	BTC CoinType = "btc"
	BCH CoinType = "bch"
	//ETH CoinType = "eth"
)

//CoinTypeValue coin_typeの値
var CoinTypeValue = map[CoinType]uint8{
	BTC: 1,
	BCH: 2,
	//ETH: 3,
}

func (c CoinType) String() string {
	return string(c)
}

// ValidateBitcoinType BitcoinTypeのバリデーションを行う
func ValidateBitcoinType(val string) bool {
	if _, ok := CoinTypeValue[CoinType(val)]; ok {
		return true
	}
	return false
}

//----------------------------------------------------
// BTCVersion
//----------------------------------------------------

//BTCVersion 実行環境
type BTCVersion int

// environment
const (
	BTCVer16 BTCVersion = 160000
	BTCVer17 BTCVersion = 170000
	BTCVer18 BTCVersion = 180000
	BTCVer19 BTCVersion = 190000
)

func (v BTCVersion) Int() int {
	return int(v)
}

//----------------------------------------------------
// NetworkType
//----------------------------------------------------

//NetworkType ネットワーク種別
type NetworkType string

// network type
const (
	NetworkTypeMainNet    NetworkType = "mainnet"
	NetworkTypeTestNet3   NetworkType = "testnet3"
	NetworkTypeRegTestNet NetworkType = "regtest"
)

//----------------------------------------------------
// AddressType
//----------------------------------------------------

//AddressType address種別
type AddressType string

// address type
const (
	AddressTypeLegacy     AddressType = "legacy"
	AddressTypeP2shSegwit AddressType = "p2sh-segwit"
	AddressTypeBech32     AddressType = "bech32"
)

func (a AddressType) String() string {
	return string(a)
}

//AddressTypeValue address_typeの値
var AddressTypeValue = map[AddressType]uint8{
	AddressTypeLegacy:     0,
	AddressTypeP2shSegwit: 1,
	AddressTypeBech32:     2,
}

//----------------------------------------------------
// KeyStatus
//----------------------------------------------------

//KeyStatus Key生成進捗ステータス
type KeyStatus string

// key_status
const (
	KeyStatusGenerated            KeyStatus = "generated"              //hd_walletによってkeyが生成された
	KeyStatusImportprivkey        KeyStatus = "importprivkey"          //importprivkeyが実行された
	KeyStatusPubkeyExported       KeyStatus = "pubkey_exported"        //pubkeyがexportされた(receipt/payment)
	KeyStatusMultiAddressImported KeyStatus = "multi_address_imported" //multiaddがimportされた(receipt/payment)
	KeyStatusAddressExported      KeyStatus = "address_exported"       //addressがexportされた
)

func (k KeyStatus) String() string {
	return string(k)
}

//KeyStatusValue key_statusの値
var KeyStatusValue = map[KeyStatus]uint8{
	KeyStatusGenerated:            0,
	KeyStatusImportprivkey:        1,
	KeyStatusPubkeyExported:       2,
	KeyStatusMultiAddressImported: 3,
	KeyStatusAddressExported:      4,
}

// ValidateKeyStatus KeyStatusのバリデーションを行う
func ValidateKeyStatus(val string) bool {
	if _, ok := KeyStatusValue[KeyStatus(val)]; ok {
		return true
	}
	return false
}

//----------------------------------------------------
// TxType
//----------------------------------------------------

//TxType トランザクション種別
type TxType string

// tx_type
const (
	TxTypeUnsigned    TxType = "unsigned"
	TxTypeUnsigned2nd TxType = "unsigned2nd"
	TxTypeSigned      TxType = "signed"
	TxTypeSent        TxType = "sent"
	TxTypeDone        TxType = "done"
	TxTypeNotified    TxType = "notified"
	TxTypeCancel      TxType = "canceled"
)

//TxTypeValue tx_typeの値
var TxTypeValue = map[TxType]uint8{
	TxTypeUnsigned:    1,
	TxTypeUnsigned2nd: 2,
	TxTypeSigned:      3,
	TxTypeSent:        4,
	TxTypeDone:        5,
	TxTypeNotified:    6,
	TxTypeCancel:      7,
}

// ValidateTxType TxTypeのバリデーションを行う
func ValidateTxType(val string) bool {
	if _, ok := TxTypeValue[TxType(val)]; ok {
		return true
	}
	return false
}

// Search SliceのtxTypes内にtが含まれているかチェックする
func (t TxType) Search(txTypes []TxType) bool {
	for _, v := range txTypes {
		if v == t {
			return true
		}
	}
	return false
}
