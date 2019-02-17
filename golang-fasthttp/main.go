package main

import (
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	fasthttp.ListenAndServe("0.0.0.0:8000", handler)
}

var routes = map[string]func(ctx *fasthttp.RequestCtx){
	"/hello":       hello,
	"/file":        file,
	"/sleep/10ms":  sleep10ms,
	"/sleep/100ms": sleep100ms,
	"/sleep/1s":    sleep1s,
	"/sleep/5s":    sleep5s,
}

func handler(ctx *fasthttp.RequestCtx) {
	routes[string(ctx.Path())](ctx)
}

func hello(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("hello"))
}

func file(ctx *fasthttp.RequestCtx) {
	fasthttp.ServeFile(ctx, "file.txt")
}

func sleep10ms(ctx *fasthttp.RequestCtx) {
	time.Sleep(10 * time.Millisecond)
	ctx.Write([]byte("ok"))
}

func sleep100ms(ctx *fasthttp.RequestCtx) {
	time.Sleep(100 * time.Millisecond)
	ctx.Write([]byte("ok"))
}

func sleep1s(ctx *fasthttp.RequestCtx) {
	time.Sleep(time.Second)
	ctx.Write([]byte("ok"))
}

func sleep5s(ctx *fasthttp.RequestCtx) {
	time.Sleep(5 * time.Second)
	ctx.Write([]byte("ok"))
}
