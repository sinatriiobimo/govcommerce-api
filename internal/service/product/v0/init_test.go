package v0

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tlkm-api/internal/repository/product"
)

func TestNew(t *testing.T) {
	type args struct {
		attr InitAttribute
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "case-1: success create service",
			args: args{
				attr: InitAttribute{
					Repo: RepoAttribute{
						ProductPostgre: &product.RepositoryMock{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					New(tt.args.attr)
				})
				return
			}
			assert.NotNil(t, New(tt.args.attr))
		})
	}

	assert.Panics(t, func() {
		New(InitAttribute{})
	})
}
