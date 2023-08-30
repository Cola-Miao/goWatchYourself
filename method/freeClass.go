package method

import (
	"encoding/json"
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type FreeClass struct {
}

func (f FreeClass) getVideos(id int) (videos map[float64]float64, err error) {
	URL := fmt.Sprintf("https://stu.ityxb.com/back/bxg_anon/courseGraduation/getDetails?objectId=%d", id)
	resp, err := global.Client.Get(URL)
	res := new(map[string]any)
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}
	videos, err = utils.ParseFreeMap(*res)
	if err != nil {
		return
	}

	return
}

func (f FreeClass) replayAttack(videos map[float64]float64, cookies []*http.Cookie) {
	var wg sync.WaitGroup
	urlS := "https://stu.ityxb.com/back/bxg_anon/courseGraduation/record"
	req, err := http.NewRequest("POST", urlS, nil)
	if err != nil {
		fmt.Println(err)
	}
	utils.SetHeader(req)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	for videoID, duration := range videos {
		wg.Add(1)
		go f.generatedPOST(videoID, duration, req, &wg)
	}
	wg.Wait()
	fmt.Println("Done")

	return
}

func (f FreeClass) generatedPOST(videoID, duration float64, req *http.Request, wg *sync.WaitGroup) {
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

func (f FreeClass) Watch(cookies []*http.Cookie) {
	videos, err := f.getVideos(global.CourseID)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.replayAttack(videos, cookies)
}
