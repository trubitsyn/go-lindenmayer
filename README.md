# go-lindenmayer [![Build Status](https://travis-ci.com/trubitsyn/go-lindenmayer.svg?branch=master)](https://travis-ci.com/trubitsyn/go-lindenmayer) [![GoDoc](https://godoc.org/github.com/trubitsyn/go-lindenmayer?status.svg)](https://godoc.org/github.com/trubitsyn/go-lindenmayer) [![Go Report Card](https://goreportcard.com/badge/github.com/trubitsyn/go-lindenmayer)](https://goreportcard.com/report/github.com/trubitsyn/go-lindenmayer)
L-systems library for Go.

**Currently under development.**

## Installation
`go get github.com/trubitsyn/go-lindenmayer`

## Usage
<pre>
package main

import (
	"github.com/trubitsyn/go-lindenmayer"
	"fmt"
)

func main() {
	sys := lindenmayer.LSystem{
		Variables: []rune{'F'},
		Constants: []rune{'+', '-'},
		Axiom:     'F',
		Rules: []Rule{
			Rule{In: "F", Out: "F+F−F−F+F"},
		},
	}
	result := lindenmayer.Iterate(&sys, 5, func(i int, s string) {
		fmt.Println(i, s)
	})
	fmt.Println(result)
	
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
    if err := lindenmayer.Process(result, operations); err != nil {
    	fmt.Println(err)
    }
}
</pre>

## Testing
```
go get -t github.com/trubitsyn/go-lindenmayer
go test github.com/trubitsyn/go-lindenmayer
```

## LICENSE
```
Copyright 2019 Nikola Trubitsyn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
