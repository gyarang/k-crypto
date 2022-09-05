package table

import (
	"bytes"
	"github.com/gyarang/k-crypto/hangul"
	"github.com/gyarang/k-crypto/util"
	"strconv"
)

func splitMixedJamo(r rune) []rune {
	switch r {
	case 'ㄲ':
		return []rune{'ㄱ', 'ㄱ'}
	case 'ㄸ':
		return []rune{'ㄷ', 'ㄷ'}
	case 'ㅃ':
		return []rune{'ㅂ', 'ㅂ'}
	case 'ㅆ':
		return []rune{'ㅅ', 'ㅅ'}
	case 'ㄳ':
		return []rune{'ㄱ', 'ㅅ'}
	case 'ㄵ':
		return []rune{'ㄴ', 'ㅈ'}
	case 'ㄶ':
		return []rune{'ㄴ', 'ㅎ'}
	case 'ㄺ':
		return []rune{'ㄹ', 'ㄱ'}
	case 'ㄻ':
		return []rune{'ㄹ', 'ㅁ'}
	case 'ㄼ':
		return []rune{'ㄹ', 'ㅂ'}
	case 'ㄽ':
		return []rune{'ㄹ', 'ㅅ'}
	case 'ㄾ':
		return []rune{'ㄹ', 'ㅌ'}
	case 'ㄿ':
		return []rune{'ㄹ', 'ㅂ'}
	case 'ㅀ':
		return []rune{'ㄹ', 'ㅎ'}
	case 'ㅐ':
		return []rune{'ㅏ', 'ㅣ'}
	case 'ㅒ':
		return []rune{'ㅑ', 'ㅣ'}
	case 'ㅔ':
		return []rune{'ㅓ', 'ㅣ'}
	case 'ㅖ':
		return []rune{'ㅕ', 'ㅣ'}
	case 'ㅘ':
		return []rune{'ㅗ', 'ㅏ'}
	case 'ㅙ':
		return []rune{'ㅗ', 'ㅏ', 'ㅣ'}
	case 'ㅚ':
		return []rune{'ㅗ', 'ㅣ'}
	case 'ㅝ':
		return []rune{'ㅜ', 'ㅓ'}
	case 'ㅞ':
		return []rune{'ㅜ', 'ㅓ', 'ㅣ'}
	case 'ㅟ':
		return []rune{'ㅜ', 'ㅣ'}
	case 'ㅢ':
		return []rune{'ㅡ', 'ㅣ'}
	default:
		return []rune{r}
	}
}

func (t Table) Encrypt(input string) (string, error) {
	inputRune := []rune(input)
	var jamos []rune

	for _, r := range inputRune {
		if r == ' ' {
			jamos = append(jamos, ' ')
			continue
		}
		h, err := hangul.NewHangul(r)
		if err != nil {
			return "", err
		}
		jamos = append(jamos, splitMixedJamo(h.Cho)...)
		jamos = append(jamos, splitMixedJamo(h.Jung)...)
		jamos = append(jamos, splitMixedJamo(h.Jong)...)
	}

	strBuf := bytes.Buffer{}
	for _, r := range jamos {
		if r == 0 {
			continue
		}
		if r == ' ' {
			strBuf.WriteRune(' ')
			continue
		}
		strBuf.WriteString(strconv.Itoa(util.RandomSelect(t.tableMap[r]).y + 1))
		strBuf.WriteString(strconv.Itoa(util.RandomSelect(t.tableMap[r]).x + 1))
	}

	return strBuf.String(), nil
}
