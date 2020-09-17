package baseconversion

type BaseConverter interface {
	Encode(number int) (string, error)
	Decode(str string) (int, error)
}
