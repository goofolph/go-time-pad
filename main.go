package main

import (
	"fmt"
	"strings"
)

func decryptJoin(cipher string, secret string) string {
	plain := make([]string, len(cipher), len(cipher))

	// decrypt string
	for i, c := range cipher {
		c_value := int(c) - int('a')
		s_value := int([]rune(secret)[i%len(secret)]) - int('a')
		p_value := c_value - s_value
		if p_value < 0 {
			p_value += 26
		}
		p := rune(p_value + int('a'))

		plain[i] = string(p)
	}

	// return plain decrypted string
	return strings.Join(plain, "")
}

func decrypt(cipher string, secret string) string {
	plain := "" // decrypted string

	// decrypt string
	for i, c := range cipher {
		c_value := int(c) - int('a')
		s_value := int([]rune(secret)[i%len(secret)]) - int('a')
		p_value := c_value - s_value
		if p_value < 0 {
			p_value += 26
		}
		p := rune(p_value + int('a'))

		plain += string(p)
	}

	// return plain decrypted string
	return plain
}

func encrypt(plain string, secret string) string {
	cipher := ""

	// encrypt string
	for i, p := range plain {
		p_value := int(p) - int('a')
		s_value := int([]rune(secret)[i%len(secret)]) - int('a')
		c_value := p_value + s_value
		if c_value > 26 {
			c_value -= 26
		}
		c := rune(c_value + int('a'))

		cipher += string(c)
	}

	// return plain decrypted string
	return cipher
}

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter ciphertext:")
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }
	// cipher := strings.TrimSuffix(input, "\n")

	// fmt.Print("Enter secret:")
	// input, err = reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }
	// secret := strings.TrimSuffix(input, "\n")

	plain := "canwestrustdiscord"
	secret := "thisisaverystrongkey"
	cipher := encrypt(plain, secret)

	fmt.Println(plain, " => ", cipher)
}
