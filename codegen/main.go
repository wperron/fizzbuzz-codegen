package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var (
	out  = flag.String("o", "", "Where to drop the output file")
	pack = flag.String("p", "main", "Name of the package to use in the generated file")
)

func main() {
	flag.Parse()

	// expects args like `3 Fizz 5 Buzz 7 Foo` alternating [factor, display] pairs
	args := flag.Args()
	factors := make([]int, len(args)/2)
	display := make([]string, len(args)/2)
	for i, v := 0, 0; i <= len(args)-2; i += 2 {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatalln(err)
		}

		// Add factor as int to list and related string to display
		factors[v] = n
		display[v] = args[i+1]
		v++
	}

	// For each factor, increment all positions divisible by it in the sequence
	// by 1.
	seq := make([]string, Multiply(factors))
	for i, f := range factors {
		for j := range seq {
			if j%f == 0 {
				seq[j] += display[i]
			}
		}
	}

	handler, err := os.Create(*out)
	if err != nil {
		log.Fatalln(err)
	}

	rawTemplate.Execute(handler, templateParam{
		Package: *pack,
		Seq:     seq,
	})
}

func Multiply(nums []int) (p int) {
	p = 1
	for _, n := range nums {
		p *= n
	}
	return p
}

type templateParam struct {
	Package string
	Seq     []string
}

var rawTemplate = template.Must(template.New("").Funcs(template.FuncMap{"join": join}).Parse(`// Code generated by go generate; DO NOT EDIT.
package {{ .Package }}

import "strconv"

var pattern = [{{ len .Seq }}]string{ {{ join .Seq ", " }} }
var max = {{ len .Seq }}

func FizzBuzz(n int) string {
	s := pattern[n % max]
	if s == "" {
		return strconv.Itoa(n)
	}
	return s
}
`))

// Slightly modified version of `strings.Join` which wraps every string in quotes
func join(ss []string, sep string) string {
	switch len(ss) {
	case 0:
		return ""
	case 1:
		return ss[0]
	}
	n := len(sep) * (len(ss) - 1)
	for i := 0; i < len(ss); i++ {
		n += len(ss[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString("\"" + ss[0] + "\"")
	for _, s := range ss[1:] {
		b.WriteString(sep)
		b.WriteString("\"" + s + "\"")
	}
	return b.String()
}
