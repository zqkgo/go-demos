package proxy

import (
	"testing"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"fmt"
	"io/ioutil"
	"time"
)

var data = `
[
  {
    "scheme": "http",
    "host":"127.0.0.1",
    "port": 8888,
    "conns": 0
  },
  {
    "scheme": "http",
    "host":"127.0.0.1",
    "port": 8887,
    "conns": 0
  },
  {
    "scheme": "http",
    "host":"127.0.0.1",
    "port": 8886,
    "conns": 0
  },
  {
    "scheme": "http",
    "host":"127.0.0.1",
    "port": 8885,
    "conns": 0
  }
]
`

var proxyPort = 50134

func TestProxy(t *testing.T) {
	// 配置server
	var servers []Server
	err := json.Unmarshal([]byte(data), &servers)
	if err != nil {
		t.Fatal("fail to mock: ", err)
	}
	proxy := NewProxy("http", "127.0.0.1", proxyPort)
	for _, s := range servers {
		ns := &Server{
			Scheme: s.Scheme,
			Host: s.Host,
			Port:s.Port,
		}
		err := proxy.AddServer(uuid.New().String(), ns) // 一开始是 &s，大坑！！！
		if err != nil {
			t.Log("fail to add server, err: ", err)
		}
	}

	// 启动server
	for k, s := range servers {
		go func(k int ,s Server) {
			mux := http.NewServeMux()
			server := http.Server{
				Addr:    s.Host + ":" + strconv.Itoa(s.Port),
				Handler: mux,
			}
			mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
				<-time.After(500 * time.Millisecond) // 模拟使server持有连接一段时间
				w.Write([]byte(fmt.Sprintf("server %d with port %d handling\n", k, s.Port)))
			})
			err := server.ListenAndServe()
			if err != nil {
				t.Error(err)
			}
		}(k, s)
	}

	// 启动代理
	go func() {
		http.ListenAndServe(":"+strconv.Itoa(proxyPort), proxy)
	}()

	// 发起请求
	for i := 0; i < 200; i++ {
		go func() {
			resp, err := http.Get("http://localhost:"+strconv.Itoa(proxyPort)+"/foo")
			if err != nil {
				t.Error("fail to request proxy, err: ", err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error("fail to read server response:, err: ", err)
			}
			t.Log(string(b))
		}()
	}

	<-time.After(10 * time.Second)
}
