// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package systems

import . "github.com/trubitsyn/go-lindenmayer"

func Algae() Lsystem {
	return Lsystem{
		Variables: []rune{'A', 'B'},
		Constants: []rune{},
		Axiom:     "A",
		Rules: []Rule{
			{"A", "AB"},
			{"B", "A"},
		},
	}
}

func FractalTree() Lsystem {
	return Lsystem{
		Variables: []rune{'0', '1'},
		Constants: []rune{'[', ']'},
		Axiom:     "0",
		Rules: []Rule{
			{"1", "11"},
			{"0", "1[0]0"},
		},
	}
}

func CantorSet() Lsystem {
	return Lsystem{
		Variables: []rune{'A', 'B'},
		Constants: []rune{},
		Axiom:     "A",
		Rules: []Rule{
			{"A", "ABA"},
			{"B", "BBB"},
		},
	}
}

func KochCurve() Lsystem {
	return Lsystem{
		Variables: []rune{'F'},
		Constants: []rune{'+', '-'},
		Axiom:     "F",
		Rules: []Rule{
			{"F", "F+F-F-F+F"},
		},
	}
}

func SierpinskiTriangle() Lsystem {
	return Lsystem{
		Variables: []rune{'F', 'G'},
		Constants: []rune{'+', '-'},
		Axiom:     "F-G-G",
		Rules: []Rule{
			{"F", "F-G+F+G-F"},
			{"G", "GG"},
		},
	}
}

func SierpinskiArrowheadCurve() Lsystem {
	return Lsystem{
		Variables: []rune{'X', 'Y'},
		Constants: []rune{'F', '+', '-'},
		Axiom:     "XF",
		Rules: []Rule{
			{"X", "YF+XF+Y"},
			{"Y", "XF-YF-X"},
		},
	}
}

func DragonCurve() Lsystem {
	return Lsystem{
		Variables: []rune{'X', 'Y'},
		Constants: []rune{'F', '+', '-'},
		Axiom:     "FX",
		Rules: []Rule{
			{"X", "X+YF+"},
			{"Y", "-FX-Y"},
		},
	}
}

func FractalPlant() Lsystem {
	return Lsystem{
		Variables: []rune{'X', 'F'},
		Constants: []rune{'+', '-', '[', ']'},
		Axiom:     "X",
		Rules: []Rule{
			{"X", "F+[[X]-X]-F[-FX]+X"},
			{"F", "FF"},
		},
	}
}
