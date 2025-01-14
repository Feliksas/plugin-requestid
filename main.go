package plugin_requestid

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

const defaultHeader = "X-Request-Id"

type Config struct {
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var b [16]byte
		_, err := rand.Read(b[:])
		if err != nil {
			log.Fatal(err)
		}
		id := fmt.Sprintf("%x", b)

		r.Header.Set(config.HeaderName, id)

		next.ServeHTTP(rw, r)
	}), nil
}
