package method

import (
	"encoding/json"
	"errors"
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/models"
	"net/http"
	"net/url"
)

func Login(input *models.Input) (cookies []*http.Cookie, err error) {
	urlValues := url.Values{
		"username":       {input.UserName},
		"password":       {input.Password},
		"automaticLogon": {"false"},
	}
	resp, err := global.Client.PostForm("https://stu.ityxb.com/back/bxg_anon/login", urlValues)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	content := new(map[string]any)
	err = json.NewDecoder(resp.Body).Decode(content)
	if err != nil {
		return
	}
	if (*content)["success"] != true {
		err = errors.New((*content)["errorMessage"].(string))
		return
	} else {
		fmt.Println("Login successful")
	}
	cookies = resp.Cookies()

	return
}
