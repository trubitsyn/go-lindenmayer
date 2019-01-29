// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package lindenmayer

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	sys := LSystem{
		Variables: []rune{'F'},
		Constants: []rune{'+', '-'},
		Axiom:     'F',
		Rules: []Rule{
			Rule{In: "F", Out: "F+F-F-F+F"},
		},
	}
	Iterate(&sys, 1, func(i int, s string) {
		if s != "F+F-F-F+F" {
			t.FailNow()
		}
	})
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
	if err := Process(lsystem, operations); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
