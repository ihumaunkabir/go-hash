package hash 

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromText(text string) (hash string, err error) {

	if hash, err = generateHashFromText(text); err != nil {
		log.Println("hash : GenerateHashFromText : ", err)
	}

	return
}

func CompareHashAndText(hashFromDB string, text string) (err error) {

	if err = compareHashAndText(hashFromDB, text); err != nil {
		log.Println("hash: CompareHashAndText : ", err)
	}

	return
}

func generateHashFromText(text string) (hash string, err error) {

	var hashByte []byte

	if hashByte, err = bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost); err != nil {
		log.Println("hash: generateHashFromText : ", err)
	}

	hash = string(hashByte[:])

	return
}

func compareHashAndText(hashFromDB, text string) (err error) {

	if err = bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(text)); err != nil {
		log.Println("hash : compareHashAndText : ", err)
	}

	return
}
