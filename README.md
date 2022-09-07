# k-crypto

일본이 쩔쩔맨 K-암호 Go 구현체(https://www.hani.co.kr/arti/society/society_general/1057247.html)

## Build
```bash
make build
```

## Usage

### Encrypt
```bash
k-crypto encrypt -k 금강산 -i "속히 상경하라"
암호화 값:
4725893888 47684889244838687168
```

### Decrypt
```bash
k-crypto decrypt -k 금강산 -i "6725833882 67513283243238511351" 
복호화 값:
ㅅㅗㄱㅎㅣ ㅅㅏㅇㄱㅕㅇㅎㅏㄹㅏ
```

## Todo

- [x] 완성형 한글 초성, 중성, 종성 분해
- [ ] 암호표
  - [x] 암호표 생성
  - [ ] 암호 키 유효성 확인
- [x] 암호화
  - [x] 쌍자음, 이중 모음, 이중 자음 처리
- [ ] 복호화
  - [x] 분해된 자모음으로 복호화
  - [ ] ~~분해된 자모음 -> 문장 변환 파서 구현~~ 
    - 분해된 자모음에 대한 명확한 명세가 없음
      - example: [ㄱ,ㄱ,ㅏ,ㄱ,ㄱ,ㅏ] -> 해당 배열이 '까까' 인지 '깍가' 인지 구분 불가.
