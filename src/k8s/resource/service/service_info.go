package service

//get Service information

type SvcInfo struct {
	Name string
	Selflink string
}

func (s *SvcInfo) GetName() string{
	return s.Name
}

func (s *SvcInfo) GetSelfLink() string{
	return s.Selflink
}

