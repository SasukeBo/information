package proto

import (
	"context"
	"fmt"
	"github.com/SasukeBo/information/channel"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
)

type devicePushServer struct {
	UnimplementedDevicePushServer
}

// PushStatus 当接收到客户端发送来的设备状态更新，做出响应
func (d *devicePushServer) PushStatus(ctx context.Context, req *Status) (*Response, error) {
	message := channel.Message{IDName: "id", IDValue: int(req.DeviceId)}
	channel.Publish("deviceStatus", message)
	return &Response{Ok: true}, nil
}

func (d *devicePushServer) PushProduct(ctx context.Context, req *Product) (*Response, error) {
	message := channel.Message{IDName: "productInsID", IDValue: int(req.InstanceId)}
	channel.Publish("productIns", message)
	return &Response{Ok: true}, nil
}

// Run start gRPC server
func Run() {
	env := beego.AppConfig.String
	port := env("grpcport")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logs.Error(err)
	}

	logs.Info("gRPC server listen on port: %s", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterDevicePushServer(grpcServer, new(devicePushServer))
	grpcServer.Serve(lis)
}
