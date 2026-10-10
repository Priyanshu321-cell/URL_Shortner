package main

import (
	"fmt"
	"math/big"
	"slices"
)

const base_62_map string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Id_to_base_62(val int64) (string, error) {
	var res []byte
	if val <= 0 {
		return string(res), nil
	}

	rem := big.NewInt(val)
	bigVal := big.NewInt(val)
	div := big.NewInt(62)
	zero := big.NewInt(0)
	for bigVal.Cmp(zero) != 0 {
		rem.Mod(bigVal, div)
		bigVal.Div(bigVal, div)
		res = append(res, base_62_map[rem.Int64()])
	}

	slices.Reverse(res)
	return string(res), nil
}

func Base_62_to_Id(val string) (int64, error) {
	if val == "" {
		return -1, nil
	}

	num := int64(0)
	for _, ch := range val {
		idx := -1
		for i := range len(base_62_map) {
			if base_62_map[i] == byte(ch) {
				idx = i
				break
			}
		}
		if idx == -1 {
			return 0, fmt.Errorf("invalid character: %c", ch)
		}
		num = num*62 + int64(idx)
	}
	return num, nil
}
