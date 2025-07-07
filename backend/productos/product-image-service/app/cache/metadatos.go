package cache

import (
	"encoding/json"
	"product-image-service/proto/imagepb"
)

type Metadata struct {
	ImageId   string `json:"image_id"`
	ProductId string `json:"product_id"`
	Filename  string `json:"filename"`
	Url       string `json:"url"`
}

func GuardarMetadata(imageId, productId, filename, url string) error {
	meta := Metadata{
		ImageId:   imageId,
		ProductId: productId,
		Filename:  filename,
		Url:       url,
	}
	data, _ := json.Marshal(meta)
	return Cliente.Set(Ctx, imageId, data, 0).Err()
}

func ObtenerMetadata(imageId string) (*Metadata, error) {
	val, err := Cliente.Get(Ctx, imageId).Result()
	if err != nil {
		return nil, err
	}
	var meta Metadata
	json.Unmarshal([]byte(val), &meta)
	return &meta, nil
}

func EliminarMetadata(imageId string) error {
	return Cliente.Del(Ctx, imageId).Err()
}

func ListarImagenesDeProducto(productId string) ([]*imagepb.ImageResponse, error) {
	keys, err := Cliente.Keys(Ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	var resultados []*imagepb.ImageResponse
	for _, key := range keys {
		val, err := Cliente.Get(Ctx, key).Result()
		if err != nil {
			continue
		}
		var meta Metadata
		json.Unmarshal([]byte(val), &meta)
		if meta.ProductId == productId {
			resultados = append(resultados, &imagepb.ImageResponse{
				ImageId: meta.ImageId,
				Url:     meta.Url,
			})
		}
	}
	return resultados, nil
}
