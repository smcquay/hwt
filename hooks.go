package hwt

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"
)

var reqStartKey = new(int)

type Timer func(path string, dur time.Duration)
type Statuser func(path, status string)

func NewMetricsHooks(timer Timer, status Statuser) *twirp.ServerHooks {
	hs := &twirp.ServerHooks{}

	hs.RequestReceived = func(ctx context.Context) (context.Context, error) {
		return markReqStart(ctx), nil
	}

	hs.ResponseSent = func(ctx context.Context) {
		name, ok := twirp.MethodName(ctx)
		if !ok {
			// XXX (sm) : something else?
			panic("missing name")
		}
		start, ok := getReqStart(ctx)
		if !ok {
			// XXX (sm) : something else?
			panic("missing start")
		}
		dur := time.Now().Sub(start)
		timer(name, dur)

		sc, ok := twirp.StatusCode(ctx)
		if !ok {
			panic("missing code")
		}
		status(name, sc)
	}

	return hs
}

func markReqStart(ctx context.Context) context.Context {
	return context.WithValue(ctx, reqStartKey, time.Now())
}

func getReqStart(ctx context.Context) (time.Time, bool) {
	t, ok := ctx.Value(reqStartKey).(time.Time)
	return t, ok
}
