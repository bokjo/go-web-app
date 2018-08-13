package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}

	w.Header().Add("Content-Encoding", "gzip")

	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()

	grw := gzipRespinseWriter{
		ResponseWriter: w,
		Writer:         gzipwriter,
	}

	gm.Next.ServeHTTP(grw, r)
}

type gzipRespinseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipRespinseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
