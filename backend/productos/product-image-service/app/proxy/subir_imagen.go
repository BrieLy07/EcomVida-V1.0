package proxy

import (
	"fmt"
	"time"

	"product-image-service/app/cache"
	"product-image-service/app/s3"
	pb "product-image-service/proto/imagepb"
)

func SubirImagen(req *pb.UploadImageRequest) (*pb.ImageResponse, error) {
	id := fmt.Sprintf("img-%d", time.Now().UnixNano())
	filename := id + "-" + req.Filename

	url, err := s3.SubirAObjeto(filename, req.ImageData)
	if err != nil {
		return nil, err
	}

	err = cache.GuardarMetadata(id, req.ProductId, filename, url)
	if err != nil {
		return nil, err
	}

	return &pb.ImageResponse{ImageId: id, Url: url}, nil
}
