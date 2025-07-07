package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"product-image-service/app/handlers"
	pb "product-image-service/proto/imagepb"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("No se pudo escuchar en el puerto %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterImageServiceServer(s, &handlers.ImageServiceServer{})

	fmt.Println("Servidor gRPC corriendo en el puerto", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
