package proxy

import (
	"product-image-service/app/cache"
	"product-image-service/app/s3"
	pb "product-image-service/proto/imagepb"
)

func EliminarImagen(req *pb.DeleteImageRequest) (*pb.DeleteResponse, error) {
	metadata, err := cache.ObtenerMetadata(req.ImageId)
	if err != nil {
		return nil, err
	}

	err = s3.EliminarObjeto(metadata.Filename)
	if err != nil {
		return nil, err
	}

	err = cache.EliminarMetadata(req.ImageId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{Message: "Imagen eliminada"}, nil
}
