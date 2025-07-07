package memento

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartState struct {
	UsuarioID string     `json:"usuario_id"`
	Items     []CartItem `json:"items"`
}

type CartMemento struct {
	State CartState
}

func NewCartMemento(state CartState) *CartMemento {
	return &CartMemento{State: state}
}

func (m *CartMemento) GetSavedState() CartState {
	return m.State
}
