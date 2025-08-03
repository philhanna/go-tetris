package tetris

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLocation_String(t *testing.T) {
	type fields struct {
		Row int
		Col int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"happy", fields{1, 2}, "(1,2)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := Location{
				Row: tt.fields.Row,
				Col: tt.fields.Col,
			}
			assert.Equal(t, tt.want, loc.String())
		})
	}
}
