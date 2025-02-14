package handlers

import (
	"net/http"
	"reflect"
	"testing"
)

func TestHandleGet(t *testing.T) {
	type args struct {
		service *service.URLService
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleGet(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
