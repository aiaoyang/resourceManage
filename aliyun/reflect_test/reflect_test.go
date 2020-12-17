package reflect_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aiaoyang/resourceManager/config"
)

func Test_reflect(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
		t.Log("helloooooooooo")
	}
}

func Benchmark_reflect(t *testing.B) {
	for i := 0; i < t.N; i++ {
		newClients([]string{""})
		fmt.Println(time.Now())
	}
}

func newClients(region []string) {
	fmt.Println(config.GVC)
}
