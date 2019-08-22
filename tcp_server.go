package main

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/channel"
	"github.com/SasukeBo/information/models"
)

var connReg = regexp.MustCompile(`^connect:([\w\d]{16})`)
var discReg = regexp.MustCompile(`^disconn:([\w\d]{16})`)
var prodReg = regexp.MustCompile(`^product:([\w\d]{16})`)
var stopReg = regexp.MustCompile(`^shutdown:([\w\d]{16})`)
var dataReg = regexp.MustCompile(`^data;id:([\d]+);(.*)`)

// FetchIDReg 匹配设备ID
var FetchIDReg = regexp.MustCompile(`^device_(\d+)$`)

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

// handleClient 处理连接
func handleClient(conn net.Conn) {
	var topics []string // 存储当前连接注册的 topic

	for {
		// 设置接收数据缓存大小为512
		dataBytes := make([]byte, 512)
		_, err := conn.Read(dataBytes)
		if err != nil {
			break
		}

		dataStr := string(dataBytes)
		logs.Info("receive data from TCP: ", dataStr)
		if ts, ok := dataRouter(conn, dataStr, topics).([]string); ok {
			topics = ts
		}
	}

	defer logs.Info("close connection from", conn.RemoteAddr())
	defer conn.Close()
	defer offlineDevices(topics)
}

// dataRouter 分解数据string，处理元数据
func dataRouter(conn net.Conn, dataStr string, topics []string) interface{} {
	dataAtoms := strings.Split(dataStr, "@")

	for _, dataAtom := range dataAtoms {
		if matches := connReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			topic := handleStatus(conn, matches[1], "connect")
			if v, ok := topic.(string); ok {
				topics = append(topics, v)
			}
		} else if matches = discReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "disconnect")
		} else if matches = stopReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "stop")
		} else if matches = prodReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "producting")
		} else if matches = dataReg.FindStringSubmatch(dataAtom); len(matches) > 2 {
			keyAndValue := strings.Split(matches[2], ":")
			payload := make(map[string]interface{})
			payload[keyAndValue[0]] = keyAndValue[1]
			payload["_TIME_STAMP_"] = time.Now().Local()

			socketMsg := channel.SocketMsg{
				Topic:   fmt.Sprintf("device_%s", matches[1]),
				Payload: payload,
				Channel: "device",
				Event:   "data",
			}

			channel.DeviceChannel.Broadcast(socketMsg)
		}
	}

	return topics
}

// handleStatus 处理设备状态变更
func handleStatus(conn net.Conn, token, action string) interface{} {
	device := models.Device{Token: token}
	if err := device.GetBy("token"); err != nil {
		conn.Write([]byte(err.Error()))
		logs.Error(err.Error())
		return nil
	}

	var status int
	var topic string

	switch action {
	case "connect":
		status = models.DeviceStatus.OnLine
	case "disconnect":
		status = models.DeviceStatus.OffLine
	case "stop":
		status = models.DeviceStatus.Stop
	case "producting":
		status = models.DeviceStatus.Prod
	}

	deviceStatusLog := models.DeviceStatusLog{Status: status, Device: &device}
	if err := deviceStatusLog.Insert(); err != nil {
		logs.Error(err.Error())
	}

	return topic
}

func offlineDevices(topics []string) {
	for _, topic := range topics {
		if match := FetchIDReg.FindStringSubmatch(topic); len(match) > 1 {
			if id, err := strconv.ParseInt(match[1], 10, 0); err == nil {
				deviceStatusLog := models.DeviceStatusLog{Status: models.DeviceStatus.OffLine, Device: &models.Device{ID: int(id)}}
				if err := deviceStatusLog.Insert(); err != nil {
					logs.Error(err.Error())
				}
			}
		}
	}
}
