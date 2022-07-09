package main

import (
	"encoding/json"
	"fmt"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
)

var hmacSampleSecret []byte = []byte(getEnv("JWT_KEY", "liuhasldhaklsjdhklashdklajshdk,jahsdkjash"))

// Generate the session token
func getToken(username string, id uint) string {
	jsonStr, _ := json.Marshal(map[string]string{"username": username, "id": fmt.Sprint(id)})
	expired := time.Now().Add(time.Hour * 8).Unix()
	token, err := jose.Sign(string(jsonStr), jose.HS256, hmacSampleSecret, jose.Header("expire", expired))
	if err == nil {
		return token
	}

	return ""
}

// Validate the session token
func valid(token string) (valid bool, dat map[string]string) {
	var r map[string]string
	data, headers, err := jose.Decode(token, hmacSampleSecret)
	if err != nil {
		return false, nil
	}
	time64 := int64(headers["expire"].(float64))
	if time64 < time.Now().Unix() {
		return false, nil
	}
	err = json.Unmarshal([]byte(data), &r)
	if err != nil {
		return false, nil
	}
	return true, r
}
