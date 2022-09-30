package application

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
