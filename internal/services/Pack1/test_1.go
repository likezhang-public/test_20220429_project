// Package Pack1 provides Pack1 logic implement
package Pack1

import (
	"context"
	"log"

	pb "github.com/skema-repo/likezhang-public/test2/grpc-go/Module1/Pack1"
)

// Pack1
type Pack1 struct{
    pb.UnimplementedTest1Server
}

// NewPack1 建议使用NewXXX方式来返回实现类，实现类名以小写字母开头
func NewPack1() *Pack1 {
	svr := &Pack1{
		// init custom fileds
	}
	// go cron()
	return svr
}

// func cron() {
// 	time.Sleep(time.Second * 10)

// 	client := pb.NewTest1Client(grpc.DefaultConn)
// 	for range time.NewTicker(time.Second * 30).C {
// 		ctx := context.Background()
// 		ctx = metadata.AppendToOutgoingContext(ctx, "x-grpc-metrics-caller", "cron/client")
// 		startTime := time.Now()
// 		rsp, err := client.SayHello(ctx, &pb.HelloRequest{Msg: "bob"})
// 		logging.Debugf(ctx, "rsp is %v %v, cost is %dms", rsp, err, time.Since(startTime)/1e6)
// 	}
// }

// Heathcheck
func (s *Pack1) Heathcheck(ctx context.Context, req *pb.HealthcheckRequest) (rsp *pb.HealthcheckResponse,err error) {
	// implement business logic here ...
	// ...

	log.Printf("Received from Heathcheck request: %v", req)
	rsp = &pb.HealthcheckResponse{
		// Msg: "Hello " + req.GetMsg(),
	}
	return rsp,err
}

// Helloworld
func (s *Pack1) Helloworld(ctx context.Context, req *pb.HelloRequest) (rsp *pb.HelloReply,err error) {
	// implement business logic here ...
	// ...

	log.Printf("Received from Helloworld request: %v", req)
	rsp = &pb.HelloReply{
		// Msg: "Hello " + req.GetMsg(),
	}
	return rsp,err
}