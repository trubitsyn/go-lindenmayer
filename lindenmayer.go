// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package lindenmayer

import (
	"errors"
	"fmt"
	"strings"
)

type Lsystem struct {
	Variables []rune
	Constants []rune
	Axiom     string
	Rules     []Rule
	Angle     float64
}

type Rule struct {
	In  string
	Out string
}

func Iterate(lsystem *Lsystem, limit int, f func(int, string)) string {
	c := lsystem.Axiom
	for i := 0; i < limit; i++ {
		for _, rule := range lsystem.Rules {
			c = strings.Replace(c, rule.In, rule.Out, -1)
		}
		f(i, c)
	}
	return c
}

func Process(lsystem string, operations map[rune]func()) error {
	for _, c := range []byte(lsystem) {
		match := operations[rune(c)]
		if match == nil {
			return errors.New(fmt.Sprintf("could not find matching operation for character %c", c))
		}
		match()
	}
	return nil
}
