package hwt

import (
	"context"
	"errors"
	"net/http"
)

const PSK = "some key"

var reqUserKey = new(int)

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		k := req.Header.Get("sm-auth")
		if k == "" {
			w.Header().Set("www-authenticate", "sm-auth")
			http.Error(w, "missing/invalid key", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), reqUserKey, "valid user")
		h(w, req.WithContext(ctx))
	}
}

func getUser(ctx context.Context) (string, error) {
	u, ok := ctx.Value(reqUserKey).(string)
	if !ok {
		return "", errors.New("user key not found in context")
	}
	return u, nil
}
