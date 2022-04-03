# FizzBuzz in O(1) Using Codegen #

Why you may ask? Well... I was bored, that's why 🤷‍♂️

## How does it work?

In the traditional FizzBuzz problem, using the factors 3 and 5, we can notice an
infinite sequence when printing the whole sequence of numbers. That sequence has
a length of 15, which happens to be the *product of all the factors* that we
want to include in our algorithm.

```text
1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz
```

This works for 3 and 5, but it also works for any pair of numbers, and even also
works for any *set* of numbers.

For example, given the set `[a, b, ..., n]` we can determine that the repeating
sequence of our FizzBuzz game will have a length of `a x b x ... x n`. We can
also write an algorithm that is going to sieve through all of our factors in
this set and generate the sequence for us, with the right word substituted in
the right place.

From there, all we have to do is take our input number, take its modulo against
the length of the sequence, and we have the index in the sequence where the
answer will be. If that index returns an empty string, then we simply return the
input number back.

## How to use this thing?

The code generator is located in `codegen/main.go` and takes to named parameters:

- `-o` - File path where the generated code will be written to.
- `-p` - Package name for the generated file. Defaults to `main`.

The factors to use are positional arguments, they should be alternated with the
word to be used in the substitution.

complete example:

```bash
$ go run ./codegen -o lib.go -p main 3 Fizz 5 Buzz
```

### Go Generate

The entrypoint `main.go` already contains a `//go:generate` comment so simply
running `go generate` will also work.