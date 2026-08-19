package categories

type CategoryAdapter interface {
	// Category entity -> other model
	Convert(source *Category) any

	// other model -> Category entity
	Rebuild(source any) (dest *Category, err error)
}
