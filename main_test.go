package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name    string
		text    []byte
		wantint int
		wanterr error
	}{
		{
			name:    "всё правильно",
			text:    []byte("hello"),
			wantint: 5,
			wanterr: nil,
		},
		{
			name:    "выход за границы utf8",
			text:    []byte{0xff, 0xfe, 0xfd},
			wantint: 0,
			wanterr: ErrInvalidUTF8,
		},
	}

	for _, temp := range cases {
		t.Run(temp.name, func(t *testing.T) {
			got, err := GetUTFLength(temp.text)
			if err != temp.wanterr {
				t.Errorf("GetUTFLength(%v) error = %v; want %v", temp.text, err, temp.wanterr)
			}
			if got != temp.wantint {
				t.Errorf("GetUTFLength(%v) = %v; want %d", temp.text, got, temp.wantint)
			}
		})
	}
}
