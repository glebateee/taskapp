package web_service

import (
	"fmt"
	"os"
)

func (s *WebService) GetMainPage() ([]byte, error) {
	htmlFilePath := os.Getenv("PROJECT_ROOT")
	if htmlFilePath == "" {
		htmlFilePath = "./public/index.html"
	}
	html, err := s.webRepository.GetFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("get file from repository: %w", err)
	}
	return html, nil
}
