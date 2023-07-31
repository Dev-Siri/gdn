package models

type Config struct {
	OriginServer string `json:"origin_server"`
	Log          bool   `json:"log"`
	CacheDir     string `json:"cache_dir"`
}
