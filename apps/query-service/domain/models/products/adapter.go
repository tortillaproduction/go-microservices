package products

type ProductAdapter interface {
	// Product entity -> other model
	Convert(source *Product) any

	// other model -> Product entity
	Rebuild(source any) (dest *Product, err error)
}
