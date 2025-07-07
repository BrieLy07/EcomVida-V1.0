package proxy

import (
	"product-image-service/app/cache"
	pb "product-image-service/proto/imagepb"
)

func ObtenerImagenesPorProducto(req *pb.ProductImagesRequest) (*pb.ProductImagesResponse, error) {
	imagenes, err := cache.ListarImagenesDeProducto(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.ProductImagesResponse{Images: imagenes}, nil
}
