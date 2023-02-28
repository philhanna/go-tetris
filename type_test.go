package tetris

import "testing"

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
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
