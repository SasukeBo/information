package proto

import (
	"context"
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"google.golang.org/grpc"
	"io"
	"net"
)

type devicePushServer struct {
	UnimplementedDevicePushServer
}

// PushStatus 当接收到客户端发送来的设备状态更新，做出响应
func (d *devicePushServer) PushStatus(ctx context.Context, req *Device) (*Response, error) {
	o := orm.NewOrm()
	device := models.Device{ID: int(req.Id)}
	if err := o.Read(&device); err != nil {
		return &Response{Ok: false, Message: fmt.Sprintf("%v", err)}, nil
	}

	logs.Info("device %s status change to %d", device.Name, device.Status)
	return &Response{Ok: true}, nil
}

func (d *devicePushServer) PushProducts(stream DevicePush_PushProductsServer) error {
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&Response{Ok: true})
		}
		if err != nil {
			return err
		}
		logs.Info("device with id %d produce product %d", product.DeviceId, product.InstanceId)
	}
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
