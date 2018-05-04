package client

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func hash(pubKey, priKey, ts string) string {
	s := fmt.Sprintf("%s%s%s", ts, priKey, pubKey)
	hash := md5.New()
	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}

func timeStamp() string {
	ts := strconv.Itoa(time.Now().Nanosecond())
	return ts
}

func buildRoute(s ...string) string {
	result := baseURI
	for _, st := range s {
		result = result + st
	}

	return fmt.Sprintf("%s?", result)
}

func addParameters(publicKey, privateKey string) string {
	ts := timeStamp()
	hash := hash(publicKey, privateKey, ts)
	params := url.Values{}
	params.Set("apikey", publicKey)
	params.Set("hash", hash)
	params.Set("ts", ts)

	return params.Encode()
}
