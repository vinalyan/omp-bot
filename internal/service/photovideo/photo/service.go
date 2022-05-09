package photo

import "github.com/ozonmp/omp-bot/internal/model/photovideo"

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
