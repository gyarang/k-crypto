package table

import "errors"

var (
	DecryptNotValidInput = errors.New("decrypt: Not valid input")
)

func (t Table) Decrypt(input string) ([]rune, error) {
	inputRunes := []rune(input)
	var decryptRunes []rune

	i := 0
	for i < len(inputRunes) {
		if inputRunes[i] == ' ' {
			decryptRunes = append(decryptRunes, ' ')
			i++
			continue
		}

		if i == len(inputRunes)-1 {
			return []rune{}, DecryptNotValidInput
		}

		if inputRunes[i] < '1' || inputRunes[i] > '9' ||
			inputRunes[i+1] < '1' || inputRunes[i+1] > '9' {
			return []rune{}, DecryptNotValidInput
		}

		y, x := inputRunes[i]-'1', inputRunes[i+1]-'1'
		r := t.table[y][x]
		decryptRunes = append(decryptRunes, r)
		i += 2
	}

	return decryptRunes, nil
}
