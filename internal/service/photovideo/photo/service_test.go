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

func TestDummyPhotoService_Create(t *testing.T) {
	type fields struct {
		allEntities []photovideo.Photo
	}
	type args struct {
		photo photovideo.Photo
	}

	photo_1 := photovideo.Photo{Id: 1, Name: "one"}
	photo_2 := photovideo.Photo{Name: "two"}
	photo_nil_name := photovideo.Photo{}

	photos := []photovideo.Photo{photo_1}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{"Фото 1", fields{}, args{photo_2}, 0, false},
		{"Фото 2", fields{photos}, args{photo_2}, 1, false},
		{"Ошибка. Нет имени", fields{photos}, args{photo_nil_name}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DummyPhotoService{
				allEntities: tt.fields.allEntities,
			}
			got, err := d.Create(tt.args.photo)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyPhotoService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DummyPhotoService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDummyPhotoService_Update(t *testing.T) {
	type fields struct {
		allEntities []photovideo.Photo
	}
	type args struct {
		photoID uint64
		photo   photovideo.Photo
	}

	photo_1 := photovideo.Photo{Id: 1, Name: "one"}
	photo_2 := photovideo.Photo{Id: 2, Name: "two"}
	//photo_nil_name := photovideo.Photo{}

	photos := []photovideo.Photo{photo_1}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Фото 1", fields{photos}, args{0, photo_2}, false},
		{"Ошибка. Пустое значение", fields{}, args{0, photo_2}, true},
	}

	//TODO тут добавить проверку на изменение
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DummyPhotoService{
				allEntities: tt.fields.allEntities,
			}
			if err := d.Update(tt.args.photoID, tt.args.photo); (err != nil) != tt.wantErr {
				t.Errorf("DummyPhotoService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDummyPhotoService_Remove(t *testing.T) {
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
		want    bool
		wantErr bool
	}{
		{"Фото 1", fields{photos}, args{0}, true, false},
		{"Ошибка. Пустое значение", fields{}, args{0}, false, true},
		{"Ошибка. Большой аргумент", fields{photos}, args{12}, false, true},
	}
	for _, tt := range tests {
		//count := len(tt.fields.allEntities)
		t.Run(tt.name, func(t *testing.T) {
			d := &DummyPhotoService{
				allEntities: tt.fields.allEntities,
			}
			count := d.Len()
			got, err := d.Remove(tt.args.photoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyPhotoService.Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DummyPhotoService.Remove() = %v, want %v", got, tt.want)
			}
			if d.Len() != count-1 && err == nil {
				t.Errorf("Длинна не уменьшилась было %d, стало %d", count, d.Len())
			}
		})
	}
}
