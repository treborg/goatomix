package history

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/treborg/goatomix/levelsets"
)

func sg(sGrid []string) [][]byte {
	gGrid, _ := levelsets.GridToBytes(sGrid)
	return gGrid
}
func cases(s string) []byte {
	bb := []byte(s)
	for i, b := range bb {
		bb[i] = b - 'a'
	}
	return bb
}

func TestCheckMove(t *testing.T) {
	assert := assert.New(t)
	assert.False(isEmpty('#'), "want false, got true")
	assert.False(isEmpty('A'), "want false, got true")
	assert.False(isEmpty(' '), "want false, got true")
	assert.True(isEmpty('.'), "want true, got false")

	grid := sg([]string{
		/*
		 abcdefghijk   */
		"###########", //a
		"#B..#.....#", //b
		"#.........#", //c
		"#..D......#", //d
		"##........#", //e
		"#.........#", //f
		"#.........#", //g
		"###########", //h
	})

	cases := []struct {
		Move string
		RE   string
	}{
		{"bbbb", `no move`},
		{"aaba", `no atom`},
		{"bbaa", `diagonal`},

		{"bbhb", `out of bounds`}, // A down
		{"bbkh", `out of bounds`}, // A right

		{"bbgb", `down, move blocked`},
		{"bbbg", `right, move blocked`},

		{"dddb", ""}, // left wall
		{"dddj", ""}, // right wall
		{"ddbd", ""}, // top wall
		{"ddgd", ""}, // bottom wall

		{"dddc", "left, move NOT blocked"},
		{"dddi", "right, move NOT blocked"},
		{"ddcd", "up, move NOT blocked"},
		{"ddfd", "down, move NOT blocked"},
	}
	for _, tc := range cases {

		move := NewMove(tc.Move)
		re := tc.RE

		msg := fmt.Sprintf("move %#v, error like: %s", move, re)

		t.Run(msg, func(t *testing.T) {
			err := CheckMove(grid, move)
			if re == "" {
				assert.Nilf(err, "move '%s', failed unexpectedly", move)
				return
			}
			if assert.Error(err, "wanted error, got nil") {
				matched, _ := regexp.MatchString(re, err.Error())
				if !matched {
					t.Errorf("wanted '%s', got  '%s'", re, err)
				}
			}
		})
	}
	//*/
}
