package processor

import (
	"mafia-go-crawler/common"
	"mafia-go-crawler/model"
	"fmt"
	"github.com/opesun/goquery"
	"strings"
	"mafia-go-crawler/common/util"
	"mafia-go-crawler/common/log"
)

type UrlProcessor struct {
	name string
}

func constructUrl() MafiaProcessor {
	return &UrlProcessor{name: common.PROCESSOR_URL}
}

func (p *UrlProcessor) Name() string {
	return p.name
}

func (p *UrlProcessor) Do(input *model.CrawlResult) (*model.ProcessResult, error) {
	log.Debug(fmt.Sprintf("[Processor][%s]%s", p.Name(), input.Url.Url))

	htmlNodes, err := goquery.ParseString(input.Page)
	if err != nil {
		return nil, err
	}
	aNodes := htmlNodes.Find("a")
	var subUrls []model.Input
	for i := 0; i < len(aNodes); i ++ {
		text := aNodes.Eq(i).Text()
		href := strings.TrimSpace(aNodes.Eq(i).Attr("href"))
		// check
		if text == "" || href == "" {
			continue
		}
		if !util.RegUrl(href) {
			continue
		}

		tempUrl := model.Input{Url: href, Desc: text}
		// TODO terminate condition
		subUrls = append(subUrls, tempUrl)
		log.Debug(fmt.Sprintf("[Processor][%s]trans data, url:%s, text:%s", p.Name(), href, text))
	}

	return &model.ProcessResult{Url: input.Url, SubUrls: subUrls}, nil
}
