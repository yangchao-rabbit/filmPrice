package middleware

import (
	"bytes"
	"encoding/json"
	"filmPrice/config"
	"filmPrice/internal/models"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志中间件

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w CustomResponseWriter) String(l *log.Logger) string {
	var resp models.Response
	body := w.body.Bytes()
	err := json.Unmarshal(body, &resp)
	if err != nil {
		l.Printf(" [Error] 解析请求返回值错误：%v", err)
	}

	if config.Get().Log.Level != "debug" {
		resp.Data = nil
	}
	return resp.String()
}

func Logger(l *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		// 元数据只能拿一次，所以需要重新赋值
		bodyRequest, err := ctx.GetRawData()
		if err != nil {
			l.Printf(" [Error] %v", err)
		}

		// 重新将body数据写入请求
		ctx.Request.Body = io.NopCloser(bytes.NewReader(bodyRequest))

		// 替换返回数据
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		ctx.Next()

		statusCode := ctx.Writer.Status()

		switch {
		case statusCode >= 400 && statusCode <= 499:
			l.Printf("[Error] Status: %v | Method: %v | Path: %v | ClientIP: %v | time: %v | RequestBody: %v | Response: %v | Error: %v",
				statusCode,
				ctx.Request.Method,
				ctx.Request.URL,
				ctx.ClientIP(),
				time.Since(t),
				string(bodyRequest),
				blw.String(l),
				ctx.Errors.String(),
			)
		case statusCode >= 500:
			l.Printf("[Error] Status: %v | Method: %v | Path: %v | ClientIP: %v | time: %v | RequestBody: %v | Response: %v | Error: %v",
				statusCode,
				ctx.Request.Method,
				ctx.Request.URL,
				ctx.ClientIP(),
				time.Since(t),
				string(bodyRequest),
				blw.String(l),
				ctx.Errors.String(),
			)
		default:
			l.Printf("Status: %v | Method: %v | Path: %v | ClientIP: %v | time: %v | RequestBody: %v | Response: %v",
				statusCode,
				ctx.Request.Method,
				ctx.Request.URL,
				ctx.ClientIP(),
				time.Since(t),
				string(bodyRequest),
				blw.String(l),
			)
		}
	}
}
