package tx

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/action"
)

type Storager interface {
	CreateFilePath(actionType action.ActionType, txType TxType, txID int64) string
	GetFileNameType(filePath string) (*FileName, error)
	ValidateFilePath(filePath string, txTypes []TxType) (action.ActionType, TxType, int64, error)
	ReadFile(path string) (string, error)
	WriteFile(path, hexTx string) (string, error)
}

// FileRepository to store transaction info as csv file
type FileRepository struct {
	filePath string
	logger   *zap.Logger
}

type FileName struct {
	ActionType  action.ActionType
	TxType      TxType
	TxReceiptID int64
}

// NewFileRepository
func NewFileRepository(filePath string, logger *zap.Logger) *FileRepository {
	return &FileRepository{
		filePath: filePath,
		logger:   logger,
	}
}

// about file structure
// e.g. ./data/tx/receipt/receipt_8_unsigned_1534744535097796209
//  - ./data/tx/ : file path
//  - receipt/   : actionType
//  - receipt_8_unsigned_1534744535097796209 : {actionType}_{txReceiptID}_{txType}_{timestamp}

// CreateFilePath create file path for transaction file
func (r *FileRepository) CreateFilePath(actionType action.ActionType, txType TxType, txID int64) string {

	// ./data/tx/receipt/receipt_8_unsigned_1534744535097796209
	baseDir := fmt.Sprintf("%s%s/", r.filePath, actionType.String())
	return fmt.Sprintf("%s%s_%d_%s_", baseDir, actionType.String(), txID, txType)
}

// GetFileNameType returns as FileName type
func (r *FileRepository) GetFileNameType(filePath string) (*FileName, error) {
	// just file path or full path
	//./data/tx/receipt/receipt_8_unsigned_1534744535097796209
	tmp := strings.Split(filePath, "/")
	fileName := tmp[len(tmp)-1]

	//receipt_5_unsigned_1534466246366489473
	//s[0]: actionType
	//s[1]: txReceiptID
	//s[2]: txType
	//s[3]: timestamp
	s := strings.Split(fileName, "_")
	if len(s) != 4 {
		return nil, errors.Errorf("invalid file path: %s", fileName)
	}

	fileNameType := FileName{}

	//Action
	if !action.ValidateActionType(s[0]) {
		return nil, errors.Errorf("invalid file path: %s", fileName)
	}
	fileNameType.ActionType = action.ActionType(s[0])

	//receiptID
	var err error
	fileNameType.TxReceiptID, err = strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return nil, errors.Errorf("invalid file path: %s", fileName)
	}

	//txType
	if !ValidateTxType(s[2]) {
		return nil, errors.Errorf("error: invalid file: %s", fileName)
	}
	fileNameType.TxType = TxType(s[2])

	return &fileNameType, nil
}

// ValidateFilePath validate file path which could be full path
func (r *FileRepository) ValidateFilePath(filePath string, expectedTxTypes []TxType) (action.ActionType, TxType, int64, error) {
	fileType, err := r.GetFileNameType(filePath)
	if err != nil {
		return "", "", 0, err
	}
	//txType
	if !(fileType.TxType).Search(expectedTxTypes) {
		return "", "", 0, errors.Errorf("txType is invalid: %s", fileType.TxType)
	}
	return fileType.ActionType, fileType.TxType, fileType.TxReceiptID, nil
}

// ReadFile read file
func (r *FileRepository) ReadFile(path string) (string, error) {
	ret, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Errorf("fail to call ioutil.ReadFile(%s) error: %v", path, err)
	}

	return string(ret), nil
}

// WriteFile write file
func (r *FileRepository) WriteFile(path, hexTx string) (string, error) {
	ts := strconv.FormatInt(time.Now().UnixNano(), 10)

	//crate directory if not exisiting
	tmp1 := strings.Split(path, "/")
	tmp2 := tmp1[0 : len(tmp1)-1] //cut filename
	dir := strings.Join(tmp2, "/")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	byteTx := []byte(hexTx)
	fileName := path + ts
	err := ioutil.WriteFile(fileName, byteTx, 0644)
	if err != nil {
		return "", errors.Errorf("fail to call ioutil.WriteFile(%s) error: %s", fileName, err)
	}

	return fileName, nil
}
