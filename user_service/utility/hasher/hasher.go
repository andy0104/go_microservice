package hasher

import "golang.org/x/crypto/bcrypt"

func HashText(txt string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CompareHash(hash string, str string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)); err != nil {
		return false
	}

	return true
}
