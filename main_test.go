package main

import "testing"

var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestEncrypt(t *testing.T) {
	plain := "HELLO"
	secret := "XMCKL"
	expected := "4QNVZ"

	cipher := encrypt(plain, secret, charset)

	if cipher != expected {
		t.Errorf("got %s, expected %s", cipher, expected)
	}
}

func TestDecrypt(t *testing.T) {
	expected := "HELLO"
	secret := "XMCKL"
	cipher := "4QNVZ"

	plain := decrypt(cipher, secret, charset)

	if plain != expected {
		t.Errorf("got %s, expected %s", plain, expected)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encrypt("X9BPNK2GV8GXJKN38Y6DW297TLTNFSJFKN2WHASQLMECT2QVNAZ66W4SM4R6RZILQCA7OSZAHU6W2YNXVREMTH7J56TGCJ1L1IB58DY2KJ412W1BLU4U03IPAZKXFHDB", "45H98I8MHV4WL8DT7D71U9KI8E2L7P1ZQ52H7JIUR3OOVXHRO1G1P8OUA7C3KBE0GC1G75FXUW5ZGDJIH6KUN5XDHZIS0TN4RSG7CYO6Z34IDQ13E5SNPVP9Y0GMI4YK", charset)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decrypt("4MBWYWCIWT3I6VWIX2U6TR2CI44LIE5J6SHIFYSR01CVJYCA4D28IH2UT3GQ6LUAQPD2UALWL8ANKZ5URQUH4WG3OXXIYGHBEN0M7FS23OI2KS9H9UQOFR4MLBIBKRC6", "RDD1D0IDOVA5SF6OCRV8XQCK0KYE9N827JV5Q7Q5NYAJI2ON7LOE6OZ30FJN5EFAERDPTH3U6MRAG61H5YWNVOH7FU1NV3C0TNP728SZFEBM0GUTII5VM89BO84LLEAS", charset)
	}
}
