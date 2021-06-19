package funcs

import (
	"crypto/tls"
	"github.com/robfig/cron"
	"goutils/utils"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func ScheduleMtLogin() {
	c := cron.New()

	// https://cron.qqe2.com/
	spec := os.Getenv("mtlogcorn")
	err := c.AddFunc(spec, mtLogin)
	if err != nil {
		panic(err)
	}
	c.Start()
}

func mtLogin() {
	utils.ConsolePl("mtlogin running")
	payload := make(url.Values)
	payload.Add("username", os.Getenv("mtloguname"))
	payload.Add("password", os.Getenv("mtlogpass"))
	req, err := http.NewRequest(
		http.MethodPost,
		"https://kp.m-team.cc/takelogin.php",
		strings.NewReader(payload.Encode()),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://kp.m-team.cc")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Referer", "gzip, deflate, br")
	req.Header.Set("Accept-Encoding", "https://kp.m-team.cc")

	// fiddler抓包使用
	/*proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}*/
	// 设置跳过不安全的 HTTPS
	tls11Transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			MaxVersion:         tls.VersionTLS11,
			InsecureSkipVerify: true,
		},
		//Proxy: proxy,
	}
	// 全局设置一个 cookie 存储器，不然 302等后续请求无法自动携带之前请求的cookie
	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Transport: tls11Transport,
		Jar:       jar,
	}
	resp, err := client.Do(req)
	/*body, err := ioutil.ReadAll(r.Body)*/
	if err != nil {
		utils.ConsolePl(err)
	}
	var showMap = map[string]string{}
	showMap["status"] = resp.Status
	utils.ConsolePf("mtlogin response_%v", showMap)
}
