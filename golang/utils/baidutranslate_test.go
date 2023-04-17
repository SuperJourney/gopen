package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
) // Import the "assert" package

func TestTranslate(t *testing.T) {
	type args struct {
		message  string
		fromLang string
		toLang   string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "Translation Success",
			args: args{
				message:  "apple",
				fromLang: "en",
				toLang:   "zh",
			},
			want:      "苹果",
			assertion: assert.NoError,
		},
		{
			name: "Translation Success",
			args: args{
				message:  "苹果 智能节气是对方 啊是对方",
				fromLang: "zh",
				toLang:   "en",
			},
			want:      "Apple's intelligent solar term is the other party, it's the other party",
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Translate(tt.args.message, tt.args.fromLang, tt.args.toLang)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
