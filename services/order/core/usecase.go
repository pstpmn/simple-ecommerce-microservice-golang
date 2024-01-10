package orderCore

type (
	u struct {
		repo   IOrderRepo
		helper IHelper
	}
)

// CreateOrder implements IOrderUseCase.
func (*u) CreateOrder(order Order) (*OrderProfile, error) {
	panic("unimplemented")
}

// CreateOrderDetails implements IOrderUseCase.
func (*u) CreateOrderDetails(order []OrderDetail) (*OrderProfile, error) {
	panic("unimplemented")
}

// CancelOrder implements IOrderUseCase.
func (*u) CancelOrder(customerId string, orderId string) error {
	panic("unimplemented")
}

// GetOrderDetail implements IOrderUseCase.
func (*u) GetOrderDetail(customerId string) {
	panic("unimplemented")
}

func NewUseCase(repo IOrderRepo, helper IHelper) IOrderUseCase {
	return &u{
		repo:   repo,
		helper: helper,
	}
}
