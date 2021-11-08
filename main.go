package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/asmcos/requests"
	"github.com/tidwall/gjson"
)

func make_spotify(email string, password string) {
	url := "https://spclient.wg.spotify.com/signup/public/v1/account"
	headers := requests.Header{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0",
	}
	data := requests.Datas{
		"birth_day":             "13",
		"birth_month":           "7",
		"birth_year":            "1999",
		"collect_personal_info": "undefined",
		"creation_flow":         "",
		"creation_point":        "https://www.spotify.com/jp/",
		"displayname":           email,
		"gender":                "neutral",
		"iagree":                "1",
		"key":                   "a1e486e2729f46d6bb368d6b2bcda326",
		"platform":              "www",
		"referrer":              "",
		"send-email":            "0",
		"thirdpartyemail":       "0",
		"email":                 email + "@gmail.com",
		"password":              password,
		"password_repeat":       password,
	}

	r, _ := requests.Post(url, headers, data)
	status := gjson.Get(r.Text(), "status")
	//fmt.Println(r.Text())

	if status.String() == "1" {
		login_token := gjson.Get(r.Text(), "login_token")
		info := email + "@gmail.com:" + password + " | " + login_token.String()
		fmt.Println("[+] " + info)
		file, _ := os.OpenFile("ac.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		file.Write(([]byte)(info + "\n"))
		file.Close()
	} else if status.String() == "20" {
		info := gjson.Get(r.Text(), "errors.email")
		fmt.Println(info.String() + email)
	} else if status.String() == "0" {
		fmt.Println("You have reached the maximum attempts allowed. Please try again in an hour.")
		os.Exit(0)
	} else if status.String() == "320" {
		fmt.Println("Het lijkt erop dat je een proxy-dienst gebruikt. Schakel deze diensten uit en probeer het opnieuw. Neem voor meer informatie contact op met de klantenservice.")
		os.Exit(0)
	}

}

//ランダム文字列参考元 https://qiita.com/srtkkou/items/ccbddc881d6f3549baf1
func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func main() {
	for {
		email := RandString(15)
		password := RandString(10)
		make_spotify(email, password)
		time.Sleep(1 * time.Second)
	}
}
