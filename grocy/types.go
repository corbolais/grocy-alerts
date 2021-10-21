package grocy

type CurrentVolatileStockResponse struct {
	DueProduct     []CurrentStockResponse `json:"due_products"`
	OverdueProduct []CurrentStockResponse `json:"overdue_products"`
	ExpiredProduct []CurrentStockResponse `json:"expired_products"`
	MissingProduct []CurrentStockResponse `json:"missing_products"`
}

type CurrentStockResponse struct {
	ProductID              string  `json:"product_id"`
	Amount                 string  `json:"amount"`
	AmountMissing          string  `json:"amount_missing,omitempty"`
	AmountAggregated       string  `json:"amount_aggregated"`
	AmountOpened           string  `json:"amount_opened"`
	AmountOpenedAggregated string  `json:"amount_opened_aggregated"`
	BestBeforeDate         string  `json:"best_before_date"`
	IsAggregatedAmount     string  `json:"is_aggregated_amount"`
	Product                Product `json:"product"`
	Name                   string  `json:"name,omitempty"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SimpleProductData struct {
	DueProduct     []SimpleProduct `json:"due_products"`
	OverdueProduct []SimpleProduct `json:"overdue_products"`
	ExpiredProduct []SimpleProduct `json:"expired_products"`
	MissingProduct []SimpleProduct `json:"missing_products"`
}

type SimpleProduct struct {
	Name           string `json:"name"`
	BestBeforeDate string `json:"best_before_date,omitempty"`
	AmountMissing  string `json:"amount_missing,omitempty"`
}
