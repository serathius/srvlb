package main

import (
    "fmt"
	"log"
	"net/http"

    "github.com/elazarl/goproxy"
	"github.com/peterbourgon/srvproxy/proxy"
	"github.com/peterbourgon/srvproxy/retry"
)

func main() {
	var rt http.RoundTripper
	rt = http.DefaultTransport
	rt = proxy.Proxy(proxy.Next(rt))
	rt = retry.Retry(retry.Next(rt))

	t := &http.Transport{}
	t.RegisterProtocol("dnssrv", rt)

    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    proxy.Tr = t

    proxy.NonproxyHandler = http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
		    if r.Host == "" {
			    fmt.Fprintln(w, "Cannot handle requests without Host header, e.g., HTTP 1.0")
			    return
		    }
		    r.URL.Scheme = "dnssrv"
		    r.URL.Host = r.Host
		    proxy.ServeHTTP(w, r)
	})

    proxy.OnRequest().DoFunc(
        func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
            r.URL.Scheme = "dnssrv"
            return r,nil
    })
    log.Fatal(http.ListenAndServe(":1080", proxy))
}

