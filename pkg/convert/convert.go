//回應處理:類型轉換

package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
	//StrTo轉成string
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
	//StrTo轉成int
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
	//這類func的必要性？
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
	//StrTo轉成uint32
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
