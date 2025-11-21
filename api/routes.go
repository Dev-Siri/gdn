package api

import (
	"net/http"

	"github.com/Dev-Siri/gdn/constants"
	"github.com/Dev-Siri/gdn/db"
	"github.com/Dev-Siri/gdn/logging"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap/zapcore"
)

func fileSystemHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	if path == "/" {
		if db.CDNConfig.Log {
			go logging.Log(zapcore.InfoLevel, "Successfully served file", fasthttp.StatusOK, "HIT", ctx)
		}

		ctx.SetContentType("text/html")
		ctx.WriteString(constants.IndexHTML)
		return
	}

	file, mimeType, exists, err := db.ReadAsset(path)

	if err != nil {
		if db.CDNConfig.Log {
			go logging.Log(zapcore.ErrorLevel, "Failed to serve asset.", fasthttp.StatusInternalServerError, "HIT", ctx)
		}

		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if !exists {
		fileFromOrigin, existsOnOrigin, err := requestServer(path)

		if err != nil {
			go logging.Log(zapcore.ErrorLevel, "Failed to request origin server.", fasthttp.StatusInternalServerError, "MISS", ctx)
			ctx.Error("Failed to request origin server.", fasthttp.StatusInternalServerError)
		}

		if !existsOnOrigin && db.CDNConfig.Log {
			go logging.Log(zapcore.ErrorLevel, "Asset not found.", fasthttp.StatusNotFound, "MISS", ctx)
			ctx.Error("Asset not found.", fasthttp.StatusNotFound)
			return
		}

		go func() {
			if err := db.WriteAsset(path, fileFromOrigin); err != nil {
				logging.Log(zapcore.ErrorLevel, "Failed to save remote file to cache", fasthttp.StatusInternalServerError, "MISS", ctx)
			}
		}()

		mimeType := http.DetectContentType(fileFromOrigin)

		ctx.SetContentType(mimeType)
		ctx.Write(fileFromOrigin)
		return
	}

	if db.CDNConfig.Log {
		go logging.Log(zapcore.InfoLevel, "Successfully served file", fasthttp.StatusOK, "HIT", ctx)
	}

	ctx.SetContentType(mimeType)
	ctx.Write(file)
}

func RegisterRoutes(router *router.Router) {
	router.NotFound = fileSystemHandler
}
