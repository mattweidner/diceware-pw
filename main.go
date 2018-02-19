// dwpw.rb
// Matt Weidner <matt.weidner@gmail.com>
// Diceware passphrase generator
package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"math/big"
	"io/ioutil"
	"crypto/rand"
)

func dice_roll() int {
	r, err := rand.Int(rand.Reader, big.NewInt(7775))
	if err != nil {
		fmt.Println("Error generating random number..")
	}
	return int(r.Int64() + 1)
}

func test_all() {
	fmt.Println("Testing dice_roll...", dice_roll())
	fmt.Println("Testing dice_roll...", dice_roll())
	_, err := read_wordlist("wordlists/beale.diceware.wordlist.asc")
	if err != nil {
		fmt.Println("read_wordlist()... Fail")
	}
	fmt.Println("read_wordlist()... Success")
}

func read_wordlist(wordlist string) ([]string, error) {
	f, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening file:", wordlist)
		return nil, err
	} else {
		defer f.Close()
		reader := bufio.NewReader(f)
		contents, _ := ioutil.ReadAll(reader)
		return strings.Split(string(contents), "\n"), nil
	}
}

func main() {
	//test_all()
	fmt.Println("-----------------------------------------")
	fmt.Println("Generating passphrases...")
	fmt.Println("-----------------------------------------")
	word_bank, err := read_wordlist("wordlists/beale.diceware.wordlist.asc")
	if err != nil {
		fmt.Println(err)
	}
	const PHRASE_LENGTH int = 6
	for i := 0; i < 10; i++ {
		for j := PHRASE_LENGTH; j > 0; j-- {
			var word_index int = dice_roll()
			print(strings.ToUpper(word_bank[word_index]), " ")
		}
		print("\n")
	}
}
