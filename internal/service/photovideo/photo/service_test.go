package photo

import (
	"reflect"
	"testing"

	"github.com/ozonmp/omp-bot/internal/model/photovideo"
)

func TestDummyPhotoService_Describe(t *testing.T) {
	type fields struct {
		allEntities []photovideo.Photo
	}
	type args struct {
		photoID uint64
	}

	photo_1 := photovideo.Photo{Id: 1, Name: "one"}
	photo_2 := photovideo.Photo{Id: 2, Name: "two"}

	photos := []photovideo.Photo{photo_1, photo_2}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *photovideo.Photo
		wantErr bool
	}{
		{"Фото 1", fields{photos}, args{0}, &photo_1, false},
		{"Фото 2", fields{photos}, args{1}, &photo_2, false},
		{"Ошибка", fields{photos}, args{4}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DummyPhotoService{
				allEntities: tt.fields.allEntities,
			}
			got, err := d.Describe(tt.args.photoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyPhotoService.Describe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DummyPhotoService.Describe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDummyPhotoService_List(t *testing.T) {
	type fields struct {
		allEntities []photovideo.Photo
	}
	type args struct {
		cursor uint64
		limit  uint64
	}

	photo_1 := photovideo.Photo{Id: 1, Name: "one"}
	photo_2 := photovideo.Photo{Id: 2, Name: "two"}

	photos := []photovideo.Photo{photo_1, photo_2}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []photovideo.Photo
		wantErr bool
	}{
		{"Фото 1", fields{photos}, args{0, 1}, []photovideo.Photo{photo_1}, false},
		{"Фото 1 и 2", fields{photos}, args{0, 12}, []photovideo.Photo{photo_1, photo_2}, false},
		{"Фото 2", fields{photos}, args{1, 12}, []photovideo.Photo{photo_2}, false},
		{"Ошибка. Курсор больше индекса", fields{photos}, args{8, 1}, nil, true},
		{"Ошибка. Пустой спискок товаров", fields{}, args{8, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DummyPhotoService{
				allEntities: tt.fields.allEntities,
			}
			got, err := d.List(tt.args.cursor, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyPhotoService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DummyPhotoService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
