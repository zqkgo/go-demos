package proxy

import (
	"sync"
	"net/http"
	"github.com/micro/go-log"
	"math"
	"errors"
	"strconv"
	"net/url"
	"strings"
	"io/ioutil"
)

type IProxy interface {
	Forward(http.ResponseWriter, *http.Request) error
}

type Server struct {
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	Port   int  `json:"port"`
	conns  int
}

type Proxy struct {
	servers map[string]*Server
	mu      sync.RWMutex
	scheme  string
	host    string
	port    int
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := p.Forward(w, r)
	if err != nil {
		log.Fatal("fail to forward, err: ", err)
	}
}

func (p *Proxy) Forward(w http.ResponseWriter, r *http.Request) error {
	// 选择server
	s, err := p.PickServer()
	if err != nil {
		return err
	}

	// 处理请求代理
	serverHost := s.Scheme + "://" + s.Host + ":" + strconv.Itoa(s.Port)
	serverURL := serverHost + r.URL.Path
	r.Host = serverHost
	r.URL, err = url.Parse(serverURL)
	if err != nil {
		return err
	}
	r.RequestURI = ""

	// TODO: 设置转发header
	// r.Header.Set("Origin", p.URL())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 不接受目标主机的重定向
			// 返回最近一次请求的响应
			return http.ErrUseLastResponse
		},
	}

	// 开始转发
	s.conns++
	resp, err := client.Do(r)
	s.conns--
	if err != nil {
		return p.Forward(w, r)
	}
	defer resp.Body.Close()

	// 处理响应代理
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		w.Header().Set(k, strings.Join(v, ";"))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	w.Write(b)

	return nil
}

func (p *Proxy) PickServer() (*Server, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if len(p.servers) == 0 {
		return nil, errors.New("no servers")
	}
	// 选择连接数最小的server
	var s *Server
	minConn := math.MaxInt64
	for _, server := range p.servers {
		if server.conns < minConn {
			minConn = server.conns
			s = server
		}
	}
	return s, nil
}

func (p *Proxy) URL() string {
	return p.scheme + "://" + p.host + strconv.Itoa(p.port)
}

func (p *Proxy) AddServer(key string, server *Server) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, ok := p.servers[key]
	if ok {
		return errors.New("key has been used")
	}
	p.servers[key] = server
	return nil
}

func (p *Proxy) RemoveServer(key string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, ok := p.servers[key]
	if !ok {
		return errors.New("no server found")
	}
	delete(p.servers, key)
	return nil
}

func NewProxy(scheme, host string, port int) *Proxy {
	return &Proxy{
		servers: make(map[string]*Server),
		scheme: scheme,
		host: host,
		port: port,
	}
}