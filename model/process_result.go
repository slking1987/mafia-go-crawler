package model

type ProcessResult struct {
	Url     Input
	SubUrls []Input
	Images  []Image
}

type Image struct {
	Url  string
	Desc string
}
