package shortener

import(
	"crypto/sha256"
	"fmt"
	"os"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Of(s string) []byte {
	algo := sha256.New()
	algo.Write([]byte(s))
	return algo.Sum(nil)
}

func base58Encode(raw []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(raw)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink, userId string) string {
	urlHash := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHash).Uint64()
	finalString := base58Encode([]byte(fmt.Sprintf("%d", generatedNumber)))[:8]
	return finalString
}