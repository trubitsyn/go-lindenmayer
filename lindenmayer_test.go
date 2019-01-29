// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package lindenmayer_test

import (
	"fmt"
	"github.com/trubitsyn/go-lindenmayer"
	"github.com/trubitsyn/go-lindenmayer/systems"
	"testing"
)

func TestIterate(t *testing.T) {
	sys := systems.SierpinskiTriangle()
	lindenmayer.Iterate(&sys, 1, func(i int, s string) {
		if s != "F-GG+F+GG-F-GG-GG" {
			t.FailNow()
		}
	})
}

func BenchmarkIterate1(b *testing.B) {
	benchmarkIterate(1, b)
}

func BenchmarkIterate2(b *testing.B) {
	benchmarkIterate(2, b)
}

func BenchmarkIterate3(b *testing.B) {
	benchmarkIterate(3, b)
}

func BenchmarkIterate4(b *testing.B) {
	benchmarkIterate(4, b)
}

func BenchmarkIterate5(b *testing.B) {
	benchmarkIterate(5, b)
}

func BenchmarkIterate6(b *testing.B) {
	benchmarkIterate(6, b)
}

func BenchmarkIterate7(b *testing.B) {
	benchmarkIterate(7, b)
}

func BenchmarkIterate8(b *testing.B) {
	benchmarkIterate(8, b)
}

func BenchmarkIterate9(b *testing.B) {
	benchmarkIterate(9, b)
}

func BenchmarkIterate10(b *testing.B) {
	benchmarkIterate(10, b)
}

func BenchmarkIterate11(b *testing.B) {
	benchmarkIterate(11, b)
}

func BenchmarkIterate12(b *testing.B) {
	benchmarkIterate(12, b)
}

func benchmarkIterate(limit int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		sys := systems.KochCurve()
		lindenmayer.Iterate(&sys, limit, func(_ int, _ string) {
		})
	}
}

func TestProcess(t *testing.T) {
	lsystem := "F+F-F-F+F"
	operations := map[rune]func(){
		'F': func() {
			// draw forward
		},
		'+': func() {
			// turn left 90°
		},
		'-': func() {
			// turn right 90°
		},
	}
	if err := lindenmayer.Process(lsystem, operations); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
