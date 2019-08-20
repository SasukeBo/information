package models

import (
	"fmt"
	"net"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models/channel"
)

var connReg = regexp.MustCompile(`^@connect:([\w\d]{16})`)
var discReg = regexp.MustCompile(`^@disconn:([\w\d]{16})`)
var dataReg = regexp.MustCompile(`^@data;id:([\d]+);(.*)`)

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

	var data string

	for {
		receiveData := make([]byte, 512)
		_, err := conn.Read(receiveData)
		if err != nil {
			return
		}

		data = string(receiveData)
		logs.Info("data: ", data)
		// TODO

		if matches := connReg.FindStringSubmatch(data); len(matches) > 1 {
			// 设备连接
			device := Device{Token: matches[1]}
			if err := device.GetBy("token"); err != nil {
				conn.Write([]byte(err.Error()))
				return
			}

			channel.NewTopic(fmt.Sprintf("device_%d", device.ID))
			defer channel.RmTopic(fmt.Sprintf("device_%d", device.ID))
			// TODO: 设备状态变更
		} else if matches = discReg.FindStringSubmatch(data); len(matches) > 1 {
			// 设备断开连接
			device := Device{Token: matches[1]}
			if err := device.GetBy("token"); err != nil {
				conn.Write([]byte(err.Error()))
				return
			}

			channel.RmTopic(fmt.Sprintf("device_%d", device.ID))
			// TODO: 设备状态变更记录
		} else if matches = dataReg.FindStringSubmatch(data); len(matches) > 2 {
			event := channel.Event{
				Topic:   fmt.Sprintf("device_%s", matches[1]),
				Content: matches[2],
				// TODO: 处理为json字符串
			}

			event.Send()
			logs.Info("event:", event)
		}
	}
}
