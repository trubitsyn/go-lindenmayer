// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package lindenmayer

import (
	"strings"
)

type LSystem struct {
	Variables []rune
	Constants []rune
	Axiom     rune
	Rules     []Rule
	Angle     float64
}

type Rule struct {
	In  string
	Out string
}

func Iterate(lsystem LSystem, max int, f func(string)) {
	c := string(lsystem.Axiom)
	for i := 0; i < max; i++ {
		for _, rule := range lsystem.Rules {
			c = strings.Replace(c, rule.In, rule.Out, -1)
		}
		f(c)
	}
}
