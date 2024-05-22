package main

import (
	"bufio"
	"fmt"
	"os"
)

// {{{ Part 1

const NoDigit = -1

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func intFromDigit(c rune) int {
	return int(c - '0')
}

func extractNumber(line string) int {
	const empty = ' '
	first := empty
	last := empty

	for _, c := range line {
		if isDigit(c) {
			if first == empty {
				first = c
			}
			last = c
		}
	}

	f := intFromDigit(first)
	l := intFromDigit(last)

	return f*10 + l
}

// }}}

type StringMatchNode struct {
	Current         rune
	Next            map[rune]*StringMatchNode
	ResultIfMatched int
}

type StringMatchTree struct {
	base            *StringMatchNode
	currentMatching *StringMatchNode
}

const NullRune = rune(0)

func NewStringMatchNode(current rune) *StringMatchNode {
	return &StringMatchNode{
		Current:         current,
		Next:            make(map[rune]*StringMatchNode),
		ResultIfMatched: NoDigit,
	}
}

type NewMatchTreeArgs []struct {
	Word   string
	Result int
}

func NewStringMatchTree(args NewMatchTreeArgs) StringMatchTree {
	tree := StringMatchTree{}
	tree.base = NewStringMatchNode(NullRune)
	tree.currentMatching = tree.base

	for _, a := range args {
		tree.Add(a.Word, a.Result)
	}

	return tree
}

func (t *StringMatchTree) Add(s string, result int) {
	current := t.base
	for _, c := range s {
		next, ok := current.Next[c]
		if !ok {
			current.Next[c] = NewStringMatchNode(c)
			next = current.Next[c]
		}
		current = next
	}
	current.ResultIfMatched = result
}

func (t *StringMatchTree) ResetMatching() {
	t.currentMatching = t.base
}

func (t *StringMatchTree) CompleteMatch() bool {
	if len(t.currentMatching.Next) == 0 {
		return true
	}
	return false
}

func (t *StringMatchTree) CurrentResult() int {
	return t.currentMatching.ResultIfMatched
}

func (t *StringMatchTree) MatchNext(c rune) bool {
	match, ok := t.currentMatching.Next[c]
	if !ok {
		return false
	}
	t.currentMatching = match
	return true
}

func (t *StringMatchTree) MatchFrom(s string) (int, bool) {
	t.ResetMatching()
	for _, c := range s {
		if !t.MatchNext(c) {
			return 0, false
		}
		if t.CompleteMatch() {
			return t.CurrentResult(), true
		}
	}
	return 0, false
}

func realExtractNumber(line string, matchTree *StringMatchTree) int {
	first, last := NoDigit, NoDigit

	for i := 0; i < len(line); i++ {
		r := rune(line[i])
		digit := NoDigit

		if isDigit(r) {
			digit = intFromDigit(r)
		} else if match, ok := matchTree.MatchFrom(line[i:]); ok {
			digit = match
		}

		if digit != NoDigit {
			if first == NoDigit {
				first = digit
			}
			last = digit
		}
	}

	return first*10 + last
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	matchTree := NewStringMatchTree(NewMatchTreeArgs{
		{
			Word:   "one",
			Result: 1,
		},
		{
			Word:   "two",
			Result: 2,
		},
		{
			Word:   "three",
			Result: 3,
		},
		{
			Word:   "four",
			Result: 4,
		},
		{
			Word:   "five",
			Result: 5,
		},
		{
			Word:   "six",
			Result: 6,
		},
		{
			Word:   "seven",
			Result: 7,
		},
		{
			Word:   "eight",
			Result: 8,
		},
		{
			Word:   "nine",
			Result: 9,
		},
	})

	for scanner.Scan() {
		line := scanner.Text()
		number := realExtractNumber(line, &matchTree)
		sum += number
	}

	fmt.Println(sum)
}
