package table

import (
	"errors"
	"fmt"
	"github.com/gyarang/k-crypto/hangul"
	"github.com/gyarang/k-crypto/util"
)

var (
	validJaum  = []rune{'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	validMoum  = []rune{'ㅏ', 'ㅑ', 'ㅓ', 'ㅕ', 'ㅗ', 'ㅛ', 'ㅜ', 'ㅠ', 'ㅡ', 'ㅣ'}
	tableRunes = []rune{'ㄱ', 'ㄴ', 'ㄷ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅅ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㅏ', 'ㅑ', 'ㅓ', 'ㅕ', 'ㅗ', 'ㅛ', 'ㅜ', 'ㅠ', 'ㅡ', 'ㅣ'}
)

var (
	NotValidLength = errors.New("length of key should be 3")
	NotValidInput  = errors.New("input key is not valid hangul")
)

type Point struct {
	x, y int
}

type Table struct {
	table    [][]rune
	tableMap map[rune][]Point
}

func NewTable(key string) (Table, error) {
	hArr, err := getHangulKey(key)
	if err != nil {
		return Table{}, err
	}

	runeTable := make([][]rune, 9)
	tableMap := make(map[rune][]Point)

	for i, h := range hArr {
		// 세로줄은 9부터 시작
		index := ((3 - i) * 3) - 1
		runeTable[index] = make([]rune, 9)
		runeTable[index][0] = h.Cho
		tableMap[h.Cho] = append(tableMap[h.Cho], Point{x: 0, y: index})
		index--

		runeTable[index] = make([]rune, 9)
		runeTable[index][0] = h.Jung
		tableMap[h.Jung] = append(tableMap[h.Jung], Point{x: 0, y: index})
		index--

		runeTable[index] = make([]rune, 9)
		runeTable[index][0] = h.Jong
		tableMap[h.Jong] = append(tableMap[h.Jong], Point{x: 0, y: index})
	}

	for y, line := range runeTable {
		start := util.GetIndexOfItem(tableRunes, line[0])
		for x := 1; x < 9; x++ {
			start++
			r := tableRunes[start%24]
			line[x] = r
			tableMap[r] = append(tableMap[r], Point{x: x, y: y})
		}
	}

	table := Table{table: runeTable, tableMap: tableMap}

	return table, nil
}

// Print prints bit readable table (hard coded)
func (t Table) Print() {
	fmt.Print("  ")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d  ", i)
	}
	fmt.Println()
	for y := len(t.table) - 1; y >= 0; y-- {
		fmt.Printf("%d ", y+1)
		for x := 0; x < len(t.table[y]); x++ {
			fmt.Printf("%s ", string(t.table[y][x]))
		}
		fmt.Print("\n")
	}
}

func getHangulKey(key string) ([]hangul.Hangul, error) {
	input := []rune(key)
	if len(input) != 3 {
		return nil, NotValidLength
	}

	hanguls := make([]hangul.Hangul, 3)
	for i, v := range input {
		var err error
		hanguls[i], err = hangul.NewHangul(v)
		if err != nil {
			return nil, NotValidInput
		}
	}

	for _, h := range hanguls {
		if !isValidHangul(h) {
			return nil, NotValidInput
		}
	}

	return hanguls, nil
}

func isValidHangul(input hangul.Hangul) bool {
	return util.IsItemInSlice(validJaum, input.Cho) &&
		util.IsItemInSlice(validMoum, input.Jung) &&
		util.IsItemInSlice(validJaum, input.Jong)
}
