package main

import "testing"

func TestEncrypt(t *testing.T) {
	plain := "hello"
	secret := "xmckl"
	expected := "eqnvz"

	cipher := encrypt(plain, secret)

	if cipher != expected {
		t.Errorf("got %s, expected %s", cipher, expected)
	}
}

func TestDecryptJoin(t *testing.T) {
	expected := "hello"
	secret := "xmckl"
	cipher := "eqnvz"

	plain := decrypt(cipher, secret)

	if plain != expected {
		t.Errorf("got %s, expected %s", plain, expected)
	}
}

func TestDecrypt(t *testing.T) {
	expected := "hello"
	secret := "xmckl"
	cipher := "eqnvz"

	plain := decrypt(cipher, secret)

	if plain != expected {
		t.Errorf("got %s, expected %s", plain, expected)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encrypt("kutmxpbauapewfqurwtxhkcmgzgruubfxqctumrggiiomrgkkqyoocvbimmqivpbhthypglqmyhplmbaqhtbbelkpagsqrmteuuegufwtjgcfronbwimpnzcltohgffv", "unzdnsxanbwuchmwtidmxkrdisaxivsqvwfbddaivxwmwswlcwhviqaovhrxuttvsavpdhtbnekvugcrnsarkczlastbpylsvhzjmeszxwkyrlhqmghiawbfsotadrgk")
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decrypt("izqfxiqjnvrwrmptuapxmczzsibcuacowousnqykqetmtyggtzpecneprhtwctrevafeejvvedjevaulfggftkttfucjhjyiukpzsxwlrhfmmsjlpcheyoewbmisuvyv", "rrghmquissytdvqeubvnghcyhmrkrxftnrrveeckbsfyekdtjviptxtjoguevufndoxkkucsbitevnwyvkhmhudutubcswzimchsakcbhgezxchciwymfklerqhuyuxf")
	}
}

func BenchmarkDecryptJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decryptJoin("bnwxqkssclpqvnxryuwsqbzjhqozyfsgfdmriooveusrjwnhozhomooocfywwmfnjyxhzbrqdhqsbewdonbzhtezsmwktovqjlhsdgjidcfdmgtwadhniamplgwmqxoo", "nazfdblnlfsnvvgsmybcjdutmdtdqqcsezxsjcmjpjfclauhvwwebnqnjoxxclldnlylaplpvckiwcjvfpoperthxlwcmyeempkajzkuywnbxlvzarlhsjeofifrexpv")
	}
}
