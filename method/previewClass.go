package method

import (
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type PreviewClass struct {
}

func (p PreviewClass) getPoints(previewID string, cookies []*http.Cookie) (points map[string]float64, err error) {
	URL := fmt.Sprintf("https://stu.ityxb.com/back/bxg/preview/info?previewId=%s", previewID)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	utils.SetHeader(req)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := global.Client.Do(req)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	points, err = utils.NewParsePreviewMap(body)
	if err != nil {
		return
	}

	return
}

func (p PreviewClass) replayAttack(points map[string]float64, cookies []*http.Cookie) {
	var wg sync.WaitGroup
	urlS := "https://stu.ityxb.com/back/bxg/preview/updateProgress"
	req, err := http.NewRequest("POST", urlS, nil)
	if err != nil {
		return
	}
	utils.SetHeader(req)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	for pointID, videoDuration := range points {
		wg.Add(1)
		go p.generatedPOST(pointID, videoDuration, req, &wg)
		utils.WaitTime()
	}
	wg.Wait()
	fmt.Println("Done")

	return
}

func (p PreviewClass) generatedPOST(pointID string, videoDuration float64, req *http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
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
		log.Println(err)
	}
	fmt.Println("working...")

	return
}

func (p PreviewClass) Watch(cookies []*http.Cookie) (err error) {
	points, err := p.getPoints(global.PreviewID, cookies)
	if err != nil {
		return
	}
	p.replayAttack(points, cookies)

	return
}
