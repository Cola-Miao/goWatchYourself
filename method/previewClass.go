package method

import (
	"encoding/json"
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type PreviewClass struct {
}

func (p PreviewClass) getChapters(previewID string, cookies []*http.Cookie) (chapters map[string]float64, err error) {
	URL := fmt.Sprintf("https://stu.ityxb.com/back/bxg/preview/info?previewId=%s", previewID)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	utils.SetHeader(req)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := global.Client.Do(req)
	res := new(map[string]any)
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}
	chapters, err = utils.ParsePreviewMap(*res)
	if err != nil {
		return
	}
	fmt.Println(chapters)

	return
}

func (p PreviewClass) replayAttack(chapters map[string]float64, cookies []*http.Cookie) {
	var wg sync.WaitGroup
	urlS := "https://stu.ityxb.com/back/bxg/preview/updateProgress"
	req, err := http.NewRequest("POST", urlS, nil)
	if err != nil {
		fmt.Println(err)
	}
	utils.SetHeader(req)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	for pointID, videoDuration := range chapters {
		wg.Add(1)
		go p.generatedPOST(pointID, videoDuration, req, &wg)
	}
	wg.Wait()
	fmt.Println("Done")

	return
}

func (p PreviewClass) generatedPOST(pointID string, videoDuration float64, req *http.Request, wg *sync.WaitGroup) {
	utils.RandomSleep()
	data := url.Values{
		"previewId":       {global.PreviewID},
		"pointId":         {pointID},
		"watchedDuration": {fmt.Sprintf("%.0f", videoDuration)},
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

func (p PreviewClass) Watch(cookies []*http.Cookie) {
	chapters, err := p.getChapters(global.PreviewID, cookies)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.replayAttack(chapters, cookies)
}
