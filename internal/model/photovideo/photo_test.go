package photovideo

import (
	"reflect"
	"testing"
)

func TestNewPhoto(t *testing.T) {
	type args struct {
		id   uint64
		name string
	}
	tests := []struct {
		name string
		args args
		want *Photo
	}{
		{"Kanon", args{1, "Kanon"}, NewPhoto(1, "Kanon")},
		{"Empty", args{}, &Photo{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPhoto(tt.args.id, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoto_String(t *testing.T) {
	type fields struct {
		Id   uint64
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Kanon", fields{1, "Kanon"}, "ID: 1, Name: Kanon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Photo{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Photo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
