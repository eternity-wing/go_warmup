package shortlinkconversion

type linkConverter interface {
	Encode(number int) (string, error)
	Decode(str string) (int, error)
}

type shortLinkConvertor struct {
	convertor linkConverter
}

func InitConvertor(c linkConverter) *shortLinkConvertor {
	return &shortLinkConvertor{
		convertor: c,
	}
}

func (s *shortLinkConvertor) setConvertor(c linkConverter) {
	s.convertor = c
}

func (s *shortLinkConvertor) Encode(number int) (string, error) {
	return s.convertor.Encode(number)
}

func (s *shortLinkConvertor) Decode(str string) (int, error) {
	return s.convertor.Decode(str)
}
