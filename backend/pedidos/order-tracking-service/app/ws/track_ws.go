package ws

import (
	"encoding/json"
	"net/http"
	"order-tracking-service/app/redis"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketTracking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "No se pudo abrir el WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	suscripcion := redis.SuscribirseAEstados("pedidos")
	defer suscripcion.Close()

	for {
		msg, err := suscripcion.ReceiveMessage(redis.Ctx)
		if err != nil {
			break
		}

		var estado redis.EstadoPedido
		if err := json.Unmarshal([]byte(msg.Payload), &estado); err != nil {
			continue
		}

		if estado.ID == orderID {
			conn.WriteJSON(estado)
		}
	}
}
