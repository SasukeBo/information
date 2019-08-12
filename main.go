package main

import (
	"net"

	_ "github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/channel"
	_ "github.com/SasukeBo/information/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	go RunTCP()
	beego.Run()
}

// RunTCP start up tcp server
func RunTCP() {
	var tcpAddr *net.TCPAddr
	var err error
	var listener *net.TCPListener
	port := beego.AppConfig.String("tcpport")

	if tcpAddr, err = net.ResolveTCPAddr("tcp4", ":"+port); err != nil {
		logs.Error("TCP adderss resolve error:", err)
	}

	if listener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		logs.Error("TCP listen error:", err)
	}

	logs.Info("TCP server Running on http://127.0.0.1:%s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		logs.Info("ok receive connect from", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	defer logs.Info("close connection from", conn.RemoteAddr())

	for {
		receiveData := make([]byte, 512)
		n, err := conn.Read(receiveData)
		if err != nil {
			logs.Error("receive error:", err)
			break
		}

		if n == 0 {
			break
		}

		event := channel.Event{Topic: "fakeData", Content: string(receiveData)}
		logs.Info("event:", event)
		event.Send()
	}
}
