package jwt

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"log"
	"encoding/json"
	"strings"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Resp struct {
	Token string `json:"token"`
}

var jwtSecret = "BsLubCFdYfoYHzDUZnIpjPuOXSWlUaPX"

func Jwt() {
	mux := http.NewServeMux()
	mux.HandleFunc("/jwt/auth", func(w http.ResponseWriter, r *http.Request) {
		// 执行身份验证操作

		// 生成 jwt token
		u := User{
			Name:   "Will",
			Age:    27,
			Gender: "Male",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": u,
		})
		tokenStr, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			log.Fatal("fail to make signed token, err :", err)
		}
		resp := Resp{tokenStr}
		b, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("fail to marshal, err:", err)
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(b)
	})
	mux.HandleFunc("/jwt/visit", jwtMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("you are accessing authorized resource"))
	}))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func jwtMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取Scheme+Token
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			w.Write([]byte("authorize failed: no token found"))
			return
		}

		// 去掉Scheme
		exploded := strings.Split(tokenStr, " ")
		if len(exploded) != 2 {
			w.Write([]byte("authorize failed: invalid token"))
			return
		}

		realToken := strings.Trim(exploded[1], " \"")
		// 验证签名
		token, err := jwt.Parse(realToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method, method:", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			w.Write([]byte("authorize failed: fail to parse token with err ->" + fmt.Sprintf("%s", err)))
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.Write([]byte("authorize failed: fail to parse claims with err ->" + fmt.Sprintf("%s", err)))
			return
		}
		// 提取claims中的非敏感数据
		u := claims["user"].(map[string]interface{})
		user := User{
			Name:   u["name"].(string),
			Age:    int(u["age"].(float64)),
			Gender: u["gender"].(string),
		}

		// 保存user数据
		fmt.Println(user)

		handler.ServeHTTP(w, r)
	})
}
