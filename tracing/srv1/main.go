package main

import (
	"net/http"
	"log"
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	httpclient "github.com/openzipkin/zipkin-go/middleware/http"
)

var tracer *zipkin.Tracer

func main() {
	// host使用docker container name
	reporter := httpreporter.NewReporter("http://host.docker.internal:9411/api/v2/spans")
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint("srv1", "host.docker.internal:65512")
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
		zipkinhttp.SpanName("serve srv1"),
	)

	f := &Foo{}
	http.Handle("/trace", serverMiddleware(f))

	http.ListenAndServe(":65512", nil)
}

type Foo struct {}

func (f *Foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clientTags := map[string]string{
		"client": "testClient",
	}

	transportTags := map[string]string{
		"conf.timeout": "default",
	}

	client, err := httpclient.NewClient(
		tracer,
		httpclient.WithClient(&http.Client{}),
		httpclient.ClientTrace(true),
		httpclient.ClientTags(clientTags),
		httpclient.TransportOptions(httpclient.TransportTags(transportTags)),
	)
	if err != nil {
		log.Fatalf("unable to create http client: %+v", err)
	}

	// 向srv2发起请求
	req, _ := http.NewRequest("GET", "http://host.docker.internal:65513/test", nil)
	req = req.WithContext(r.Context())
	res, err := client.DoWithAppSpan(req, "get srv2")
	if err != nil {
		log.Fatalf("unable to execute client request: %+v", err)
	}
	res.Body.Close()

	// 向srv3发起请求
	req, _ = http.NewRequest("GET", "http://host.docker.internal:65514/test", nil)
	req = req.WithContext(r.Context())
	//req.URL.Host = "host.docker.internal:65514"
	res, err = client.DoWithAppSpan(req, "get srv3")
	if err != nil {
		log.Fatalf("unable to execute client request: %+v", err)
	}
	res.Body.Close()

	w.Write([]byte("response form srv1"))
}