package photovideo

import "fmt"

type Photo struct {
	Id   uint64
	Name string
}

func NewPhoto(id uint64, name string) *Photo {
	return &Photo{
		Id:   id,
		Name: name,
	}
}

func (a Photo) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", a.Id, a.Name)
}
