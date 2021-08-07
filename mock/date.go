package mock

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	rand_range_int   = 1000
	rand_range_int64 = 1000000000
	size_8           = 8
	str              = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	format           = "2006-01-02 15:04:05"
	Mock             = MOCK(0)
)

type MOCK int

func (m MOCK) Int() int {
	reset()
	return rand.Intn(rand_range_int)
}

func (m MOCK) Int64() int64 {
	reset()
	return rand.Int63n(rand_range_int64)
}

func (m MOCK) Float64(n int) float64 {
	reset()
	f, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(n)+"f", rand.Float64()*1000), 64)
	return f
}

func (m MOCK) String(size int) string {
	reset()
	tmpStr := ""
	for i := 0; i < size; i++ {
		tmpStr += string(str[rand.Intn(62)])
	}
	return tmpStr
}

func (m MOCK) DefaultString() string {
	reset()
	tmpStr := ""
	for i := 0; i < size_8; i++ {
		tmpStr += string(str[rand.Intn(62)])
	}
	return tmpStr
}

func (m MOCK) DefaultTimestamp() int64 {
	reset()
	bef := time.Now().AddDate(0, -1, 0).Unix()
	randN := time.Now().Unix() - bef
	return time.Unix(rand.Int63n(randN)+bef, 0).Unix()
}

func (m MOCK) DefaultDateTime() string {
	reset()
	bef := time.Now().AddDate(0, -1, 0).Unix()
	randN := time.Now().Unix() - bef
	return time.Unix(rand.Int63n(randN)+bef, 0).Format(format)
}

func reset() {
	rand.Seed(time.Now().UnixNano())
}
