package generator

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const DocBaseURL = "https://cloud.tencent.com"

func GenerateAllClients() {
	d, err := goquery.NewDocument("https://cloud.tencent.com/document/api")
	if err != nil {
		panic(err)
	}

	apis := make([]*APIDocEntries, 0)

	d.Find(".pandect-box a").Each(func(i int, a *goquery.Selection) {
		text := strings.TrimSpace(a.Text())

		if strings.HasPrefix(text, "API") && strings.Contains(text, "概览") {
			if urlOrApiContents, ok := a.Attr("href"); ok {
				parentBox := a.ParentsFiltered(".catalogue-box")

				apiDocEntries := &APIDocEntries{}
				apiDocEntries.ContentsURL = DocBaseURL + urlOrApiContents

				if u, ok := parentBox.Find("a[title=数据结构]").First().Attr("href"); ok {
					apiDocEntries.DataTypesURL = DocBaseURL + u
				}

				if apiDocEntries.DataTypesURL != "" {
					apis = append(apis, apiDocEntries)
				}
			}
		}
	})

	wg := sync.WaitGroup{}
	wg.Add(len(apis))

	jobs := make(chan func())

	for i := 0; i < 10; i++ {
		go func() {
			for runJob := range jobs {
				runJob()
				wg.Add(-1)
			}
		}()
	}

	for i := range apis {
		apiDocEntries := apis[i]
		d := NewAPIDoc(apiDocEntries)

		jobs <- func() {
			d.Scan()
			d.SyncAndWriteFiles()
		}
	}

	wg.Wait()
	close(jobs)
}
