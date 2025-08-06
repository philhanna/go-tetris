package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType_String(t *testing.T) {
	tests := []struct {
		name string
		tr   Type
		want string
	}{
		{"I", TET_I, "I"},
		{"J", TET_J, "J"},
		{"L", TET_L, "L"},
		{"O", TET_O, "O"},
		{"S", TET_S, "S"},
		{"T", TET_T, "T"},
		{"Z", TET_Z, "Z"},
		{"BOGUS", 86, "?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := tt.tr.String()
			assert.Equal(t, want, have)
		})
	}
}
