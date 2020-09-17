package main

import (
	"./lib/baseconversion/base62"
	"./lib/shortlinkconversion"
	"fmt"
)

func main() {
	base62Convertor := base62.Base62Convertor{}
	shortLinkConvertor := shortlinkconversion.InitLinkConvertor(base62Convertor)
	fmt.Println(shortLinkConvertor.Decode("zzz"))
}
