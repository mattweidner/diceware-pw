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

var version string

func dice_roll(max int64) int {
	r, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("Error generating random number..")
	}
	return int(r.Int64() + 1)
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
	fmt.Println("Random passphrase generator", version)
	fmt.Println("--------------------------------------------")
	fmt.Println("Generating passphrases...")
	fmt.Println("--------------------------------------------")
	//word_bank, err := read_wordlist("wordlists/beale.diceware.wordlist.asc")
	//word_bank, err := read_wordlist("wordlists/eff_large_wordlist.txt")
	word_bank, err := read_wordlist("wordlists/google-20000-english.txt")
	if err != nil {
		fmt.Println(err)
	}
	const PHRASE_LENGTH int = 7
	for i := 0; i < 10; i++ {
		for j := PHRASE_LENGTH; j > 0; j-- {
			var word_index int = dice_roll(int64(len(word_bank)))
			print(strings.ToUpper(word_bank[word_index]), " ")
		}
		print("\n")
	}
}
