package gserver

import (
	"context"
	ghandlers "github.com/nglmq/ozon-test/internal/app/grpc/handlers"
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"

	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func StartGRPCServer(urlService *service.URLService) {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	shortenerServer := &ghandlers.ShortenerServer{
		Service: urlService,
	}

	proto.RegisterShortenerServer(grpcServer, shortenerServer)

	reflection.Register(grpcServer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("gRPC on port 5051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	log.Println("Shutdown gRPC Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcServer.GracefulStop()

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("gRPC Server exiting")
}
