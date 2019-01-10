// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package lindenmayer

import (
	"testing"
)

func TestExample(t *testing.T) {
	sys := LSystem {
		Variables: []rune{'F'},
		Constants: []rune{'+', '-'},
		Axiom: 'F',
		Rules: []Rule{
			Rule{In: "F", Out: "F+F−F−F+F"},
		},
	}
	Iterate(sys, 1, func (s string) {
		if (s != "F+F−F−F+F") {
			t.Fail()
		}
	})
}
