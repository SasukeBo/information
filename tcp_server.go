package main

import (
	"fmt"
	"net"
	"regexp"
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
	deviceIDs := []int{}

	for {
		// 设置接收数据缓存大小为512
		dataBytes := make([]byte, 512)
		n, err := conn.Read(dataBytes)
		if err != nil {
			break
		}

		dataStr := string(dataBytes[:n])
		dataRouter(conn, dataStr, &deviceIDs)
	}

	defer logs.Info("close connection from", conn.RemoteAddr())
	defer conn.Close()
	defer offlineDevices(deviceIDs)
}

// dataRouter 分解数据string，处理元数据
func dataRouter(conn net.Conn, dataStr string, ids *[]int) {
	dataAtoms := strings.Split(dataStr, "@")

	for _, dataAtom := range dataAtoms {
		if matches := connReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "connect", ids)
		} else if matches = discReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "disconnect", ids)
		} else if matches = stopReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "stop", ids)
		} else if matches = prodReg.FindStringSubmatch(dataAtom); len(matches) > 1 {
			handleStatus(conn, matches[1], "producting", ids)
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
}

// handleStatus 处理设备状态变更
func handleStatus(conn net.Conn, token, action string, ids *[]int) {
	device := models.Device{Token: token}
	if err := device.GetBy("token"); err != nil {
		conn.Write([]byte(err.Error()))
		logs.Error(err.Error())
		return
	}

	var status int

	switch action {
	case "connect":
		status = models.DeviceStatus.OnLine
		*ids = append(*ids, device.ID)
	case "disconnect":
		status = models.DeviceStatus.OffLine
		deviceIDs := *ids
		for index, id := range deviceIDs {
			if id == device.ID {
				*ids = append(deviceIDs[:index], deviceIDs[index+1:]...)
			}
		}
	case "stop":
		status = models.DeviceStatus.Stop
	case "producting":
		status = models.DeviceStatus.Prod
	}

	deviceStatusLog := models.DeviceStatusLog{Status: status, Device: &device}
	if err := deviceStatusLog.Insert(); err != nil {
		logs.Error(err.Error())
	}
}

func offlineDevices(ids []int) {
	for _, id := range ids {
		deviceStatusLog := models.DeviceStatusLog{Status: models.DeviceStatus.OffLine, Device: &models.Device{ID: id}}
		deviceStatusLog.Insert()
	}
}
