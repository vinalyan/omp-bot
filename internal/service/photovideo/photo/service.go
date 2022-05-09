package photo

import "github.com/ozonmp/omp-bot/internal/model/photovideo"

type PhotoService interface {
	Describe(photoID uint64) (*photovideo.Photo, error)
	List(cursor uint64, limit uint64) ([]photovideo.Photo, error)
	Create(photovideo.Photo) (uint64, error)
	Update(photoID uint64, photo photovideo.Photo) error
	Remove(photoID uint64) (bool, error)
}

type DummyPhotoService struct{}

func NewDummyPhotoService() *DummyPhotoService {
	return &DummyPhotoService{}
}
