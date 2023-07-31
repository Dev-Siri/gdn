package api

import (
	"fmt"

	"github.com/Dev-Siri/gdn/db"
	"github.com/valyala/fasthttp"
)

var httpClient *fasthttp.Client = &fasthttp.Client{}

func requestServer(path string) ([]byte, bool, error) {
	req := fasthttp.AcquireRequest()

	uri := fasthttp.AcquireURI()
	uri.Update(db.CDNConfig.OriginServer + path)

	req.SetRequestURIBytes(uri.FullURI())

	res := fasthttp.AcquireResponse()

	if err := httpClient.Do(req, res); err != nil {
		return nil, false, err
	}

	if res.StatusCode() != fasthttp.StatusOK {
		if res.StatusCode() == fasthttp.StatusNotFound {
			return nil, false, nil
		}

		return nil, false, fmt.Errorf("failed to get original file from server")
	}

	content := res.Body()

	go func() {
		fasthttp.ReleaseURI(uri)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	return content, true, nil
}
