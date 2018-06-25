package gqutils

import (
	"math/rand"
	"time"
	"path/filepath"
	"os"
	"log"
	"crypto/md5"
	"encoding/hex"
)

// 生成随机字符串
func RandString(length int) string {
	str := "01234abcdefghijklm56789nopqrstuABCDEFGHIJKvwxyzLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

// 获取当前的时间戳
func NowTime() string {
	//获取当前的时间戳
	var times = time.Now().Unix()

	var lctime = time.Unix(times, 0)

	timelayout := "2006-01-02 03:04:05"

	var ltime = lctime.Format(timelayout)

	return ltime
}

func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

func Time() int64 {
	stamp := time.Now().Unix()
	return stamp
}

var (
	rand_seed int64
)

func RandSetSeed() {
	if rand_seed == 0 {
		rand_seed = time.Now().UnixNano()
		rand.Seed(rand_seed)
	}

}

func RandomInt(min int, max int) int {
	RandSetSeed()

	rangenum := max - min

	rangenum += 1
	temp := rand.Intn(rangenum)

	result := int(temp) + min

	return result
}
