package processor

import (
	"mafia-go-crawler/model"
	"mafia-go-crawler/common/log"
	"fmt"
	"github.com/opesun/goquery"
	"strings"
	"mafia-go-crawler/common"
)

type ImageProcessor struct {
	name string
}

func constructImage() MafiaProcessor {
	return &ImageProcessor{name: common.PROCESSOR_IMAGE}
}

func (p *ImageProcessor) Name() string {
	return p.name
}

func (p *ImageProcessor) Do(input *model.CrawlResult) (*model.ProcessResult, error) {
	log.Debug(fmt.Sprintf("[Processor][%s]%s", p.Name(), input.Url.Url))

	htmlNodes, err := goquery.ParseString(input.Page)
	if err != nil {
		return nil, err
	}
	selNodes := htmlNodes.Find("image")
	log.Debug(fmt.Sprintf("[Processor][%s]sel nodes:%d", p.Name(), len(selNodes)))
	var images []model.Image
	for i := 0; i < len(selNodes); i ++ {
		id := selNodes.Eq(i).Attr("id")
		src := strings.TrimSpace(selNodes.Eq(i).Attr("src"))
		// check
		if src == "" {
			continue
		}

		tempImage := model.Image{Url: src, Desc: id}
		images = append(images, tempImage)
		log.Debug(fmt.Sprintf("[Processor][%s]trans data, url:%s, desc:%s", p.Name(), src, id))
	}

	return &model.ProcessResult{Url: input.Url, Images: images}, nil
}
