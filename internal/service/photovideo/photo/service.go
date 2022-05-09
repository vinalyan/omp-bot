package photo

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/photovideo"
)

var allEntities = []photovideo.Photo{
	{Id: 1, Name: "one"},
	{Id: 2, Name: "two"},
	{Id: 3, Name: "three"},
	{Id: 4, Name: "four"},
	{Id: 5, Name: "five"},
}

type PhotoService interface {
	Describe(photoID uint64) (*photovideo.Photo, error)
	List(cursor uint64, limit uint64) ([]photovideo.Photo, error)
	Create(photovideo.Photo) (uint64, error)
	Update(photoID uint64, photo photovideo.Photo) error
	Remove(photoID uint64) (bool, error)
}

type DummyPhotoService struct {
	allEntities []photovideo.Photo
}

func NewDummyPhotoService() *DummyPhotoService {
	return &DummyPhotoService{allEntities: allEntities}
}

func (d *DummyPhotoService) Len() int {
	return len(d.allEntities)
}

func (d *DummyPhotoService) Describe(photoID uint64) (*photovideo.Photo, error) {

	if photoID > uint64(d.Len())-1 {
		return nil, fmt.Errorf("Индекс больше длины массива: %d", d.Len()-1)
	}
	return &d.allEntities[photoID], nil
}

//Возвращает список товаров раазмером limit начина с cursor
func (d *DummyPhotoService) List(cursor uint64, limit uint64) ([]photovideo.Photo, error) {

	//var res []photovideo.Photo
	last := cursor + limit

	if d.Len() == 0 {
		return nil, fmt.Errorf("Сипоск товаров пустой")
	}

	if cursor > uint64(d.Len())-1 {
		return nil, fmt.Errorf("Индекс больше длины массива: %d", d.Len()-1)
	}

	if last > uint64(d.Len()) {
		return d.allEntities[cursor:], nil
	} else {
		return d.allEntities[cursor:last], nil
	}
}
