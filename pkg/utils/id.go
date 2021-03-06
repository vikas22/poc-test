package utils

import (
	"math/rand"
	"time"
)

const (
	base                string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	maxRandomIntCeil    int    = 9999999999999
	firstJan2014EpochTs int64  = 1388534400 * 1000 * 1000 * 1000
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewID() string {
	nanosec := time.Now().UTC().UnixNano()
	random := int64(rand.Intn(maxRandomIntCeil))
	base64Rand := base62(random)
	if len(base64Rand) > 4 {
		base64Rand = base64Rand[len(base64Rand)-4:]
	}
	b62 := base62(nanosec - firstJan2014EpochTs)
	id := b62 + base64Rand
	return id
}

func RandomInt(ceil int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(ceil)
}

func base62(num int64) string {
	index := base
	encoded := ""
	for {
		encoded = string(index[num%62]) + encoded
		num = num / 62
		if num == 0 {
			break
		}
	}
	return encoded
}
