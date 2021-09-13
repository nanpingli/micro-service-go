package main

import (
	"bufio"
	"net"
)

// Go 网络编程 - TCP 简单用例

func TCP() {
	ln, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Prlog.Printf("accept error: %v\n", err)
			continue
		}
		//开始goroutine监听连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	//// 读写缓冲区
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
	}
	wr.WriteString("hello ")
	wr.Write(line)
	wr.Flush() // 一次性syscall
}
