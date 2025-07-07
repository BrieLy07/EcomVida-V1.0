package proxy

import (
	"product-image-service/app/cache"
	"product-image-service/app/s3"
	pb "product-image-service/proto/imagepb"
)

func ObtenerImagen(req *pb.GetImageRequest) (*pb.ImageData, error) {
	metadata, err := cache.ObtenerMetadata(req.ImageId)
	if err != nil {
		return nil, err
	}

	data, err := s3.ObtenerObjeto(metadata.Filename)
	if err != nil {
		return nil, err
	}

	return &pb.ImageData{
		Filename: metadata.Filename,
		Data:     data,
	}, nil
}
