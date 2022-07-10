package utils

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	clientOnce sync.Once
)

func GetProxyClient() *http.Client {
	clientOnce.Do(func() {
		proxyHost := "127.0.0.1:10086"

		baseDialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		var dialContext DialContext
		dialSocksProxy, err := proxy.SOCKS5("tcp", proxyHost, nil, baseDialer)
		if err != nil {
			panic(errors.Wrap(err, "Error creating SOCKS5 proxy"))
		}
		if contextDialer, ok := dialSocksProxy.(proxy.ContextDialer); ok {
			dialContext = contextDialer.DialContext
		} else {
			panic(errors.New("Failed type assertion to DialContext"))
		}
		httpClient = newClient(dialContext)
	})
	return httpClient
}

func newClient(dialContext DialContext) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}

type DialContext func(ctx context.Context, network, address string) (net.Conn, error)

func GetHtmlContent(url string) string {
	rsp, err := doGetRequest(url)
	if err != nil {
		logrus.Error("GetHtmlContent error", "url", url, "err", err)
		return ""
	}
	body, _ := ioutil.ReadAll(rsp.Body)
	return string(body)
}

func DownloadFile(url, filepath string) {
	out, err := os.Create(filepath)
	if err != nil {
		logrus.Error("create error", "file", filepath, "err", err)
		return
	}
	defer out.Close()
	rsp, err := doGetRequest(url)
	if err != nil {
		logrus.Error("DownloadFile error", "url", url, "err", err)
		return
	}
	defer rsp.Body.Close()
	// Writer the body to file
	_, err = io.Copy(out, rsp.Body)
	if err != nil {
		logrus.Error("copy steam error", "url", url, "err", err)
		return
	}
}

func doGetRequest(url string) (*http.Response, error) {
	hc := GetProxyClient()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", getRandomUA())
	req.Header.Add("referer", url)
	return hc.Do(req)
}
