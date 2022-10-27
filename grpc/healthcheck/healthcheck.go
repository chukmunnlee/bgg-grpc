package healthcheck

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/chukmunnlee/bgg-grpc/grpc/server"
	bgg "github.com/chukmunnlee/bgg-grpc/grpc/server"
	"google.golang.org/grpc"
)

type HealthService struct {
	UnimplementedHealthServer

	Server *grpc.Server
	bggClient bgg.BGGServiceClient
	conn *grpc.ClientConn
}

func New(intf string, targetPort int) (*HealthService, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", intf, targetPort), grpc.WithInsecure())
	if nil != err {
		return nil, err
	}
	c := server.NewBGGServiceClient(conn)
	return &HealthService{ bggClient: c , conn: conn }, nil
}

func (h *HealthService) Start(intf string, port int) error {
	hs := grpc.NewServer()
	h.Server = hs
	RegisterHealthServer(hs, h)

	connStr := fmt.Sprintf("%s:%d", intf, port)
	lis, err := net.Listen("tcp", connStr)
	if nil != err {
		return fmt.Errorf("Cannot create healthcheck listener: %v\n", err)
	}

	log.Printf("Starting healthcheck on interface %s", connStr)
	go func() {
		if err := hs.Serve(lis); nil != err {
			log.Panicf("Cannot start service: %v", err)
		}
	}()

	return nil
}

func (h *HealthService) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {

	log.Printf("Check")

	bggReq := &bgg.FindBoardgameByIdRequest{ GameId: 1 }

	if _, err := h.bggClient.FindBoardgameById(ctx, bggReq); nil != err {
		log.Printf("Check failed: %v", err)
		return &HealthCheckResponse{ Status: HealthCheckResponse_NOT_SERVING }, fmt.Errorf("%v", err)
	}

	return &HealthCheckResponse{ Status: HealthCheckResponse_SERVING }, nil
}

func (h *HealthService) Close() {
	h.conn.Close()
}
