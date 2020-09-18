package main

import (
	"./lib/baseconversion/base62"
	"./lib/shortlinkconversion"
	"fmt"
)

func main() {
	base62Convertor := base62.Convertor{}
	shortLinkConvertor := shortlinkconversion.InitConvertor(base62Convertor)
	fmt.Println(shortLinkConvertor.Decode("zzz"))
}
