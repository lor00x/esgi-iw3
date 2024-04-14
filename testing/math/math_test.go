package math

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestAdd1(t *testing.T) {

	result := Add(1, 2)

	assert.Equal(t, result, 3)
	//if result != 3 {
	//	t.Errorf("Add(1, 2) = %d; want 3", result)
	//}
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1 plus 1",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
		},
		{
			name: "Test 2 plus 2",
			args: args{
				a: 2,
				b: 2,
			},
			want: 4,
		},
		{
			name: "Test 3 plus 3",
			args: args{
				a: 3,
				b: 3,
			},
			want: 6,
		},
		//{
		//	name: "Test 3 plus 3",
		//	args: args{
		//		a: -1,
		//		b: -1,
		//	},
		//	want: 6,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	fmt.Printf("Loop %d\n", b.N)
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
