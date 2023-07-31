package db

import (
	"os"
)

func InitStorage() error {
	if _, err := os.Stat(CDNConfig.CacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(CDNConfig.CacheDir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func ReadAsset(path string) ([]byte, string, bool, error) {
	asset, mimeType, err := GetFromCache(path)

	if err != nil {
		return nil, "", false, err
	}

	return asset, mimeType, asset != nil, nil
}

func WriteAsset(path string, content []byte) error {
	if err := os.WriteFile(CDNConfig.CacheDir+path, content, 0644); err != nil {
		return err
	}

	return nil
}
