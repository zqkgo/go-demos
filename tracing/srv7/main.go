package main

import (
	"net/http"
	"log"
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"time"
)

var tracer *zipkin.Tracer

func main() {
	reporter := httpreporter.NewReporter("http://host.docker.internal:9411/api/v2/spans")
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint("srv7", "host.docker.internal:65518")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	tracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// 使用middleware自动处理server端投递span
	serverMiddleware := zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.TagResponseSize(true),
		zipkinhttp.SpanName("serve srv7"),
	)

	f := &Foo{}
	http.Handle("/test", serverMiddleware(f))
	http.ListenAndServe(":65518", nil)
}

type Foo struct {}

func (f *Foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	<- time.After(3 * time.Second)

	w.Write([]byte("response from srv7"))
}