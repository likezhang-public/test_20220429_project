package main

import (
	"context"
	"log"
	"time"

	pb "github.com/skema-repo/likezhang-public/test2/grpc-go/Module1/Pack1"
	"github.com/skema-dev/skema-go/grpcmux"
)

func main() {
    // create conn according to grpc.yaml
	conn, err := grpcmux.GetConn()
	client := pb.NewTest1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

    // heathcheckReply
	heathcheckReply, err := client.Heathcheck (ctx, &pb.HealthcheckRequest {
		// Msg: "world",
	})
	HandleReply("heathcheckReply", heathcheckReply, err)
	// log.Printf("Greeting: %s", heathcheckReply.GetMsg())

    // helloworldReply
	helloworldReply, err := client.Helloworld (ctx, &pb.HelloRequest {
		// Msg: "world",
	})
	HandleReply("helloworldReply", helloworldReply, err)
	// log.Printf("Greeting: %s", helloworldReply.GetMsg())
}

// HandleReply handle the reply
func HandleReply(replyName string, reply interface{}, err error) {
	if err != nil {
		log.Fatalf("%s could not greet: %v", replyName, err)
	}
	log.Printf(replyName+" from server: %v\n", reply)
}