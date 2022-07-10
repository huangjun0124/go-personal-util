package _024

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"goutils/utils"
	"io/ioutil"
	"os"
	"strings"
)

func CrawDagaierFlagsPictures() {
	basePath := "E:\\1024\\outs\\zipai"
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		logrus.Error("error reading dir ", " dir ", basePath)
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			path := basePath + "\\" + file.Name()
			dat, _ := os.ReadFile(path)
			var tr TrModel
			json.Unmarshal(dat, &tr)
			tr.Path = path
			downloadImgsForTr(&tr)
		}
	}
	logrus.Info("all process completed.......")
}

func downloadImgsForTr(tr *TrModel) {
	logrus.WithFields(logrus.Fields{
		"title":  tr.Title,
		"images": len(tr.ImageList),
	}).Info("downloading images")
	defer func() {
		_ = os.Remove(tr.Path)
	}()
	if len(tr.ImageList) == 0 {
		return
	}
	for i := 0; i < len(tr.ImageList); i++ {
		file := fmt.Sprintf("E:\\1024\\outs\\zipai\\%s-%03d.jpg", tr.Hash, i)
		if _, err := os.Stat(file); err == nil {
			// path/to/whatever exists
			continue
		} else if errors.Is(err, os.ErrNotExist) {
			//time.Sleep(time.Duration(rand.Int63n(1000)+1000) * time.Millisecond)
			// path/to/whatever does *not* exist
			utils.DownloadFile(tr.ImageList[i], file)
		} else {
			logrus.WithFields(logrus.Fields{
				"file": file,
				"err":  err,
			}).Error("check file error")
		}
	}
}

type TrModel struct {
	Href      string   `json:"Href"`
	Title     string   `json:"Title"`
	Date      string   `json:"Date"`
	Hash      string   `json:"Hash"`
	ImageList []string `json:"ImageList"`
	Path      string   `json:"-"`
}
