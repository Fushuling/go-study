package main

import (
	"bufio"
	"fmt"
	"os"
	"strings" // 包含字符串操作相关方法
)

func caesarDe(strCipher string, step_move byte) string {
	str_cipher := strings.ToLower(strCipher)
	str_slice_src := []byte(str_cipher)
	str_slice_dst := make([]byte, len(str_slice_src), len(str_slice_src))
	for i := 0; i < len(str_slice_src); i++ {
		if str_slice_src[i] >= 97+step_move {
			str_slice_dst[i] = str_slice_src[i] - step_move
		} else {
			str_slice_dst[i] = str_slice_src[i] + 26 - step_move
		}
	}
	fmt.Println("明文：", string(str_slice_dst))
	return string(str_slice_dst)
}

func main() {
	var step byte
	fmt.Scanln(&step)
	reader := bufio.NewReader(os.Stdin)
	string, _ := reader.ReadString('\n')
	caesarDe(string, step)

}
