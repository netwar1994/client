package qr

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	baseUrl string
}

func NewService(url string) *Service {
	return &Service{baseUrl: url}
}

func (s *Service) Encode(ctx context.Context, content string, size int) ([]byte, error) {
	client := &http.Client{}
	reqUrl := s.baseUrl + fmt.Sprintf("?data=%s&size=%dx%d", content, size, size)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		reqUrl,
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

func (s *Service) Download(data []byte) error {
	err := ioutil.WriteFile("qrcode"+".png", data, 0666)
	if err != nil {
		log.Println(err)
	}
	return nil
}
