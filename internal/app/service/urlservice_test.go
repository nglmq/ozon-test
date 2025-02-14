package service

import (
	"context"
	"reflect"
	"testing"
)

func TestNewURLService(t *testing.T) {
	type args struct {
		repository URLRepository
	}
	tests := []struct {
		name string
		args args
		want *URLService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewURLService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLService_GetOriginalURL(t *testing.T) {
	type fields struct {
		repository URLRepository
	}
	type args struct {
		ctx   context.Context
		short string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &URLService{
				repository: tt.fields.repository,
			}
			got, err := s.GetOriginalURL(tt.args.ctx, tt.args.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOriginalURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLService_ShortenURL(t *testing.T) {
	type fields struct {
		repository URLRepository
	}
	type args struct {
		ctx      context.Context
		original string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.URLResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &URLService{
				repository: tt.fields.repository,
			}
			got, err := s.ShortenURL(tt.args.ctx, tt.args.original)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortenURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortenURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
