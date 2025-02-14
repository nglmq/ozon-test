package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandlePost(t *testing.T) {
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
			assert.Equalf(t, tt.want, HandlePost(tt.args.service), "HandlePost(%v)", tt.args.service)
		})
	}
}
