package services

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/models"
)

func GeneratingAnIdProfile(i models.RequestRegistration) string {
	var dataByte []byte = []byte(i.Email + i.Login + i.Password + time.Now().String())
	var hash [32]byte = sha256.Sum256(dataByte)
	return fmt.Sprintf("%x", hash)
}
