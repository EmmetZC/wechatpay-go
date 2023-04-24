package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRequiredFieldAllSet(t *testing.T) {
	var i int32 = 0

	tests := []struct {
		name    string
		object  interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "no required fields",
			object: struct {
				A string
				B *string
			}{},
			wantErr: assert.NoError,
		},
		{
			name: "no required nullable fields",
			object: struct {
				A string `required:"true"`
				B *string
			}{},
			wantErr: assert.NoError,
		},
		{
			name: "required nullable fields but all set",
			object: struct {
				A map[int]string `required:"true"`
				B []int          `required:"true"`
				C *int32         `required:"true"`
			}{
				A: make(map[int]string),
				B: []int{1, 2},
				C: &i,
			},
			wantErr: assert.NoError,
		},
		{
			name: "required nullable fields but map is null",
			object: struct {
				A map[int]string `required:"true"`
				B []int          `required:"true"`
				C *int32         `required:"true"`
			}{
				B: []int{1, 2},
				C: &i,
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				if err == nil {
					return assert.Fail(t, "err is null", msgAndArgs)
				}
				if err.Error() != "field `A` is required and must be specified in " {
					return assert.Fail(t, "err message not match", msgAndArgs)
				}
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.wantErr(
					t, CheckRequiredFieldAllSet(tt.object),
					fmt.Sprintf("CheckRequiredFieldAllSet(%v)", tt.object),
				)
			},
		)
	}
}
