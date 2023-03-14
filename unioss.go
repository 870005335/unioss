package unioss

import (
	"errors"
	"io"
)

const (
	AliYunConst  = "aliYun"
	TencentConst = "tencent"
	QinNiuConst  = "qiNiu"
)

var currentStorage Storage

type Storage interface {
	GetObjectToFile(objectKey, downloadedFileName string) error
	DeleteObject(objectKey string) error
	PutObject(objectKey string, reader io.Reader) error
	PutObjectFromFile(objectKey, filePath string) error
	IsExists(objectKey string) (bool, error)
	GetObjectUrl(objectKey string) string
}

func NewStorage(ossName string, config Config) error {
	if config.KeyId == "" || config.KeySecret == "" || config.Bucket == "" {
		return errors.New("configuration not correct")
	}
	var err error
	switch ossName {
	case AliYunConst:
		currentStorage, err = newAliYun(config)
	case TencentConst:
		currentStorage, err = newTencent(config)
	case QinNiuConst:
		currentStorage, err = newQiNiu(config)
	default:
		return errors.New("driver not exists")
	}
	if err != nil {
		currentStorage = nil
		return err
	}
	return nil
}

func GetStorage() (Storage, error) {
	if currentStorage != nil {
		return currentStorage, nil
	}
	return nil, errors.New("driver not exists")
}
