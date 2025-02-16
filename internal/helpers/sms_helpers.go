package helpers

import (
	"fmt"
	"net/http"
	"strings"
)

// SendSMS sends an SMS using an SMS API
func SendSMS(to, message string) error {
	config := configs.LoadConfig()

	url := fmt.Sprintf("%s?to=%s&message=%s", config.SMSAPI, to, strings.ReplaceAll(message, " ", "%20"))

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send SMS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS: %s", resp.Status)
	}

	return nil
}