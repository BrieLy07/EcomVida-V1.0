package redis

import (
	"encoding/json"
	"fmt"
)

type EstadoPedido struct {
	ID     string `json:"id"`
	Estado string `json:"estado"`
}

func PublicarEstado(id string, estado string) error {
	dato := EstadoPedido{
		ID:     id,
		Estado: estado,
	}
	jsonData, err := json.Marshal(dato)
	if err != nil {
		return err
	}
	return Cliente.Publish(Ctx, "pedidos", jsonData).Err()
}

func SuscribirseAEstados(canal string) *redis.PubSub {
	return Cliente.Subscribe(Ctx, canal)
}

func ObtenerEstadoActual(id string) (string, error) {
	return Cliente.Get(Ctx, fmt.Sprintf("pedido:%s", id)).Result()
}

func GuardarEstado(id string, estado string) error {
	return Cliente.Set(Ctx, fmt.Sprintf("pedido:%s", id), estado, 0).Err()
}
