package db

import (
	"fmt"
	"net/http"
	"os"
)

func GetFromCache(path string) ([]byte, string, error) {
	content, err := os.ReadFile(CDNConfig.CacheDir + path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, "", nil
		} else {
			return nil, "", fmt.Errorf("failed to read file from cache")
		}
	}

	mimeType := http.DetectContentType(content)

	return content, mimeType, nil
}
