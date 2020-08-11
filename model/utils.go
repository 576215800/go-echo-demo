package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gitstliu/go-id-worker"
)

var currWoker *idworker.IdWorker

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}
func init() {
	currWoker = &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
}
func GetNewID() (int64, error) {
	return currWoker.NextId()
}
