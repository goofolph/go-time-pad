package main

import (
	"fmt"
	"strings"
)

func decrypt(cipher string, secret string, charset string) string {
	message := ""
	for i, c := range cipher {
		c_value := strings.Index(charset, string(c)) // numerical value of c
		s := secret[i%len(secret)]                   // secret character
		s_value := strings.Index(charset, string(s)) // numerical value of s
		m_value := c_value - s_value                 // message value
		if m_value < 0 {                             // add charset len if negative
			m_value += len(charset)
		}
		m := charset[m_value] // message character
		message += string(m)
	}
	return message
}

func encrypt(message string, secret string, charset string) string {
	cipher := ""

	for i, m := range message {
		m_value := strings.Index(charset, string(m))  // numerical value of m
		s := secret[i%len(secret)]                    // secret character
		s_value := strings.Index(charset, string(s))  // numerical value of s
		c_value := (m_value + s_value) % len(charset) // cipher value
		c := charset[c_value]                         // cipher character
		cipher += string(c)
	}
	return cipher
}

func main() {
	message := "HELLO"
	secret := "XMCKL"
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cipher := encrypt(message, secret, charset)
	fmt.Println(message, "+", secret, "=>", cipher)
}
