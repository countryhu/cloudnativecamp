package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

/*
本周作业
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/

func index(w http.ResponseWriter, r *http.Request) {
	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("k:[%s], vv:[%s]\n", k, vv)
			w.Header().Add(k, vv)
		}
	}

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	ver := os.Getenv("VERSION")
	w.Header().Set("VERSION", ver)
	fmt.Printf("VERSION:[%s]\n", ver)

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	currentIp := ClientIP(r)
	fmt.Printf("currentIp:[%s]", currentIp)
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func getCurrentIp(r *http.Request) string {
	r.Header.Get("")
	return "0.0.0"
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatalf("ListenAndServe failed.")
	}
}
