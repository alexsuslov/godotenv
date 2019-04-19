package godotenv_test

import (
	"os"
	"testing"

	"github.com/alexsuslov/godotenv"
)

func noPanic(isPanic bool) {
	if !isPanic {
		return
	}
	if r := recover(); r != nil {
		return
		// fmt.Println("Recovered in f", r)
	}
}
func TestGetPanic(t *testing.T) {
	os.Setenv("TEST_INT", "5")
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		panic bool
	}{
		{
			"GetPanic get variable, no panic",
			args{"TEST_INT"},
			"5",
			false,
		},
		{
			"GetPanic not get variable, panic",
			args{"TEST_INT1"},
			"5",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer noPanic(tt.panic)
			if got := godotenv.GetPanic(tt.args.key); got != tt.want {
				t.Errorf("GetPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}
