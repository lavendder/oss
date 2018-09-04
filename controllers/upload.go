package controllers

import (
	"net/http"
	"io/ioutil"
	"mime"
	"os"
	."oss_dfs/library/log"
	"regexp"
	"crypto/sha1"
	"oss_dfs/models"
	"io"
	"fmt"
	"time"
	"math/rand"
	"hash/crc32"
	"path/filepath"
	"strconv"
	"strings"
)

type UploadController struct {
	BaseController
}

const (
	UPLOAD_PATH    	  = "../tmp"
	MIN_FILE_SIZE     = 1       // bytes
	MAX_FILE_SIZE     = 26214400 // bytes
	IMAGE_EXTS        = "(jpg|gif|p?jpeg|(x-)?png)"
	ACCEPT_FILE_EXTS  = IMAGE_EXTS
)

type FileInfo struct {
	Url          string `json:"url,omitempty"`
	Key		     string `json:"key"`
	Name         string `json:"name"`
	Ext          string `json:"ext"`
	Secret       int    `json:"secret"`
	MimeType     string `json:"mimeType"`
	Size         int64  `json:"size,omitempty"`
	Error        string `json:"error,omitempty"`
}

func (fi *FileInfo) ValidateExt() (valid bool) {
	acceptFileTypes := regexp.MustCompile(ACCEPT_FILE_EXTS)
	if acceptFileTypes.MatchString(fi.Ext) {
		return true
	}
	fi.Error = "invalid_file_type"
	return false
}

func (fi *FileInfo) DispatchHostId() int {
	hostId := GenerateRangeNum(1, 3)
	return hostId
}

func (fi *FileInfo) DispatchFilePath() string {
	var path string
	fileKey := GetFileSha1(fi.Name)
	ieee := crc32.NewIEEE()
	io.WriteString(ieee, fileKey)
	dispatchKey := int(ieee.Sum32())
	for dispatchKey > 0 {
		subDir := dispatchKey % 100
		path = strconv.Itoa(subDir) + "/" + path
		dispatchKey = int(dispatchKey /100)
	}
	return path
}

func (fi *FileInfo) ValidateMimeType() (valid bool) {
	mimeTypeMap := map[string]string{
		"application/msword":"doc",
		"application/octet-stream":"",
		"application/pdf":"pdf",
		"application/vnd.ms-excel":"xls",
		"application/CDFV2-corrupt":"xls",
		"application/vnd.ms-publisher":"pub",
		"application/vnd.ms-powerpoint":"ppt",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document":"docx",
		"application/vnd.rn-realmedia":"rmvb",
		"application/x-msdownload":"exe",
		"image/bmp":"bmp",
		"image/x-ms-bmp":"bmp",
		"image/gif":"gif",
		"image/jpeg":"jpg",
		"image/png":"png",
		"text/plain":"txt",
		"text/xml":"xml",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":"xlsx",
		"application/x-rar":"rar",
		"application/zip":"zip",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation":"pptx",
		"application/vnd.ms-works":"wps",
		"application/vnd.ms-office":"office",
		"audio/x-wav":"wav",
		"audio/mpeg":"mp3",
		"video/mp4":"mp4",
	}
	if _, ok := mimeTypeMap[fi.MimeType]; ok {
		return true
	} else {
		fi.Error = "invalid_file_mime_type"
		return false
	}
}

func (c *UploadController) Upload() {
	bucket := c.GetString(":bucket")
	b, err := models.GetByBucket(bucket)
	//校验bucket
	if err != nil {
		c.ErrorData("",  http.StatusBadRequest, "bucket_not_allowed")
		return
	}
	//校验referer
	if b.IsReferer > 0 {
		acceptReferer := regexp.MustCompile(b.RefererValue)
		referer := c.Ctx.Input.Header("referer")
		if !acceptReferer.MatchString(referer) {
			c.ErrorData("",  http.StatusBadRequest, "invalid_referer")
			return
		}
	}
	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		c.ErrorData("",  http.StatusBadRequest, "invalid_file")
		return
	}
	if fileHeader.Size < MIN_FILE_SIZE || fileHeader.Size > MAX_FILE_SIZE {
		c.ErrorData("",  http.StatusBadRequest, "file_size_error")
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.ErrorData("",  http.StatusBadRequest, "invalid_file")
		return
	}
	//获取文件mimetype和后缀
	mimeType := http.DetectContentType(fileBytes)
	fileEndings, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		c.ErrorData("",  http.StatusInternalServerError, "cant_read_file_mime_type")
		return
	}

	fi := &FileInfo{
		Name: fileHeader.Filename,
		Ext: strings.TrimLeft(fileEndings[0], "."),
		MimeType:mimeType,
		Secret:b.Secret,
	}
	if !fi.ValidateExt() {
		c.ErrorData("",  http.StatusBadRequest, fi.Error)
		return
	}
	if !fi.ValidateMimeType() {
		c.ErrorData("",  http.StatusBadRequest, fi.Error)
		return
	}
	//生成存储路径
	fileKey := GetFileSha1(fileHeader.Filename)
	hostId := fi.DispatchHostId()
	storagePath := filepath.Join(UPLOAD_PATH, strconv.Itoa(hostId) + "/" + fi.DispatchFilePath())
	OssLogger.Info("upload_file_name_path", fileKey, storagePath)

	//存储源文件
	var storage Storage
	var storageInfo StorageInfo
	factory := new(StorageFactory)
	storage = factory.storageIns("local")

	storageInfo.fileBytes = fileBytes
	storageInfo.storageDir = storagePath
	storageInfo.fileKey = fileKey
	storageInfo.fromFile = "file"
	_, err = storage.SaveFile(storageInfo)
	if err != nil {
		c.ErrorData("",  http.StatusInternalServerError, "write_file_failed")
		return
	}
	fi.Key = fileKey
	c.SuccessData(fi)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSha1(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		OssLogger.Info("file_read_error")
	}
	defer file.Close()
	h := sha1.New()
	io.Copy(h, file)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}
