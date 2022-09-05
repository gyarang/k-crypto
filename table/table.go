package table

import (
	"errors"
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

type Table [][]rune

func NewTable(key string) (Table, error) {
	hArr, err := getHanguls(key)
	if err != nil {
		return nil, err
	}

	table := make([][]rune, 9)
	for i, h := range hArr {
		// 세로줄은 9부터 시작
		index := ((3 - i) * 3) - 1
		table[index] = make([]rune, 9)
		table[index][0] = h.Cho
		index--

		table[index] = make([]rune, 9)
		table[index][0] = h.Jung
		index--

		table[index] = make([]rune, 9)
		table[index][0] = h.Jong
	}

	for _, line := range table {
		start := util.GetIndexOfItem(tableRunes, line[0])
		for i := 1; i < 9; i++ {
			start++
			line[i] = tableRunes[start%24]
		}
	}

	return table, nil
}

func getHanguls(key string) ([]hangul.Hangul, error) {
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
