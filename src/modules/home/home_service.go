package home

type HomeService struct{}

func NewHomeService() *HomeService {
	return &HomeService{}
}

func (s *HomeService) GetHello() string {
	return "Hello from HomeController!!!"
}

func (s *HomeService) Test() string {
	return "Test"
}
