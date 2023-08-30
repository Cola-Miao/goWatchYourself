package method

import (
	"encoding/json"
	"errors"
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/models"
	"goWatchYourself/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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

func GetVideos(id int) (videos map[float64]float64, err error) {
	URL := fmt.Sprintf("https://stu.ityxb.com/back/bxg_anon/courseGraduation/getDetails?objectId=%d", id)
	resp, err := global.Client.Get(URL)
	res := new(map[string]any)
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}
	videos, err = utils.ParseMap(*res)
	if err != nil {
		return
	}

	return
}

func ReplayAttack(videos map[float64]float64, cookies []*http.Cookie) {
	var wg sync.WaitGroup
	urlS := "https://stu.ityxb.com/back/bxg_anon/courseGraduation/record"
	req, err := http.NewRequest("POST", urlS, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	for videoID, duration := range videos {
		wg.Add(1)
		go generatedPOST(videoID, duration, req, &wg)
	}
	wg.Wait()
	fmt.Println("Done")

	return
}

func generatedPOST(videoID, duration float64, req *http.Request, wg *sync.WaitGroup) {
	utils.RandomSleep()
	data := url.Values{
		"duration": {fmt.Sprintf("%.0f", duration)},
		"courseId": {strconv.Itoa(global.CourseID)},
		"videoId":  {fmt.Sprintf("%.0f", videoID)},
	}
	body := strings.NewReader(data.Encode())
	rc := io.NopCloser(body)
	req.Body = rc
	_, err := global.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("working...")
	(*wg).Done()

	return
}
