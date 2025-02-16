package utils

import (
	"math/rand"
	"time"
	"strings"
)

// GenerateRandomString generates a random string of a given length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateTempPassword generates a temporary password
func GenerateTempPassword() string {
	return GenerateRandomString(10) // 10-character temporary password
}

// FormatPhoneNumber formats a phone number to a standard format
func FormatPhoneNumber(phone string) string {
	// Remove all non-numeric characters
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Add country code if missing
	if !strings.HasPrefix(phone, "+") {
		phone = "+1" + phone // Assuming default country code is +1 (USA/Canada)
	}

	return phone
}

// FormatDate formats a time.Time object to a readable string
func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

// Contains checks if a slice contains a specific value
func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// IsEmpty checks if a string is empty or consists only of whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}