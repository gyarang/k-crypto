// 한글 유니코드표
// ⌜     587        ⌝
// 가, 각, 갂, ...깊, 깋 ⌝
// 까, 깍, 깎, ...낖, 낗 ⌟ 588

package hangul

import (
	"errors"
)

var (
	chos  = []rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	jungs = []rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}
	jongs = []rune{0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
)

var (
	NotValidHangulErr = errors.New("input value is not hangul")
)

const (
	hangulStart = '가'
	hangulEnd   = '힣'
)

type Hangul struct {
	Cho  rune
	Jung rune
	Jong rune
}

// NewHangul returns Hangul value of input value
// if input is not valid returns NotValidHangulErr
func NewHangul(input rune) (Hangul, error) {
	if !isValidHangul(input) {
		return Hangul{}, NotValidHangulErr
	}

	h := Hangul{}
	temp := input - hangulStart
	choTemp := temp / 588
	h.Cho = chos[choTemp]
	h.Jung = jungs[(temp-(choTemp*588))/28]
	h.Jong = jongs[temp%28]

	return h, nil
}

// isValidHangul returns true when input is valid
func isValidHangul(input rune) bool {
	return input >= hangulStart && input <= hangulEnd
}
