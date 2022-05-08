package main

import (
	"Module1/internal/services/Pack1"

	pb "github.com/skema-repo/likezhang-public/test2/grpc-go/Module1/Pack1"
	"github.com/skema-dev/skema-go/grpcmux"
	"github.com/skema-dev/skema-go/logging"
)

func main() {
	srv := grpcmux.NewServer()
	srvImp := Pack1.NewPack1()
	pb.RegisterTest1Server(srv, srvImp)

	ctx, mux, conn := srv.GetGatewayInfo()
	pb.RegisterTest1HandlerClient(ctx, mux, pb.NewTest1Client(conn))

    logging.Infof("Serving gRPC start...")
	if err := srv.Serve(); err != nil {
		logging.Fatalf("Serve error %v", err.Error())
	}
}
