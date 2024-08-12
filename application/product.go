package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disable"
	ENABLED  = "enable"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == DISABLED {
		return false, nil
	}
	return true, nil
}
