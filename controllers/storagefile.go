package controllers

import (
	"github.com/pkg/sftp"
	"path"
	"os"
	"fmt"
)

type Storage interface {
	SaveFile(info StorageInfo) (ret bool, err error)
}


type StorageInfo struct {
	storageDir string
	fileKey string
	fileBytes []byte
	fromFile string
}
type LocalStorage struct {
	storageDir string
	fileKey string
	fileBytes []byte
}

func (storage LocalStorage) SaveFile(info StorageInfo) (ret bool, err error) {
	if exist, _ := PathExists(info.storageDir); !exist {
		err := os.MkdirAll(info.storageDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	f, err := os.OpenFile(info.storageDir + "/" + info.fileKey, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer f.Close()
	f.Write(info.fileBytes)
	return true,nil
}

type SftpStorage struct {
}

func (storage SftpStorage) SaveFile(info StorageInfo) (ret bool, err error){
	var (
		sftpClient *sftp.Client
	)
	sftpClient, err = connect("www", "2V4B6cw9B0", "10.59.72.27", 22)
	if err != nil {
		return false, err
	}
	defer sftpClient.Close()
	var remoteDir = info.storageDir
	var remoteFileName = path.Base(info.fileKey)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		return false, err
	}
	defer dstFile.Close()
	_, writeErr := dstFile.Write(info.fileBytes)
	if writeErr != nil {
		return false, writeErr
	}
	return true, nil
}

type StorageFactory struct {

}

func (factory StorageFactory) storageIns (storageType string) Storage{
	switch storageType {
	case "local":
		return new(LocalStorage)
	case "sftp":
		return new(SftpStorage)
	default:
		return nil
	}
}
