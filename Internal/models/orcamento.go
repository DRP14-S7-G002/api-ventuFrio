package models

type Orcamento struct {
	ID              int     `json:"id"`
	ValorTotal      float64 `json:"valor_total"`
	PrazoEntrega    string  `json:"prazo_entrega"`
	AdminUserName   string  `json:"admin_user_name"`
	StatusOrcamento string  `json:"status_orcamento"`
	PedidoID        int     `json:"pedido_id"`
}
