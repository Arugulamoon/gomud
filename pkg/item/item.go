package item

type Item struct {
	Id string
}

func New(id string) *Item {
	return &Item{
		Id: id,
	}
}
