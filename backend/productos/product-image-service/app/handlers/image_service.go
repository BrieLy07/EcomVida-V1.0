package handlers

import (
	"context"

	"product-image-service/app/proxy"
	pb "product-image-service/proto/imagepb"
)

type ImageServiceServer struct {
	pb.UnimplementedImageServiceServer
}

func (s *ImageServiceServer) UploadImage(ctx context.Context, req *pb.UploadImageRequest) (*pb.ImageResponse, error) {
	return proxy.SubirImagen(req)
}

func (s *ImageServiceServer) GetImage(ctx context.Context, req *pb.GetImageRequest) (*pb.ImageData, error) {
	return proxy.ObtenerImagen(req)
}

func (s *ImageServiceServer) GetImagesByProduct(ctx context.Context, req *pb.ProductImagesRequest) (*pb.ProductImagesResponse, error) {
	return proxy.ObtenerImagenesPorProducto(req)
}

func (s *ImageServiceServer) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.DeleteResponse, error) {
	return proxy.EliminarImagen(req)
}
