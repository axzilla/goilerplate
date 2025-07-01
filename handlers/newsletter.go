package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/axzilla/goilerplate/ui/components/toast"
)

func HandleNewsletter(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		toast.Toast(toast.Props{
			Title:       "Error processing form",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Get email from form
	email := r.FormValue("email")
	if email == "" {
		toast.Toast(toast.Props{
			Title:       "Please enter an email address",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Get Mailchimp config from environment
	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	serverID := os.Getenv("MAILCHIMP_SERVER_ID")
	listID := os.Getenv("MAILCHIMP_LIST_ID")

	// Check if Mailchimp is configured
	if apiKey == "" || serverID == "" || listID == "" {
		// For development, just log the email and return success
		log.Printf("Newsletter signup (dev mode): %s", email)

		toast.Toast(toast.Props{
			Title:       "You're on the list! 🎉",
			Description: "Check your email for your 50% early bird discount.",
			Variant:     toast.VariantSuccess,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Create subscriber data
	subscriberData := map[string]any{
		"email_address": email,
		"status":        "subscribed",
		"tags":          []string{"goilerplate", "early-bird"},
	}

	jsonData, err := json.Marshal(subscriberData)
	if err != nil {
		toast.Toast(toast.Props{
			Title:       "Error processing request",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Create request to Mailchimp API
	endpoint := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members",
		serverID, listID)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		toast.Toast(toast.Props{
			Title:       "Error creating request",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Set headers and auth
	req.SetBasicAuth("anystring", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		toast.Toast(toast.Props{
			Title:       "Connection error",
			Description: "Please try again later.",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		toast.Toast(toast.Props{
			Title:       "Error reading response",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Check for "Member Exists" error
	if resp.StatusCode == 400 {
		var errorResp map[string]any
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if title, ok := errorResp["title"].(string); ok && title == "Member Exists" {
				toast.Toast(toast.Props{
					Title:       "You're already on the list!",
					Description: "Thank you for your continued interest.",
					Variant:     toast.VariantSuccess,
					Position:    "bottom-right",
					Duration:    5000,
					Dismissible: true,
					Icon:        true,
				}).Render(r.Context(), w)
				return
			}
		}
	}

	// Handle other errors
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("Mailchimp error: Status %d, Body: %s", resp.StatusCode, string(body))
		toast.Toast(toast.Props{
			Title:       "Unable to subscribe",
			Description: "Please try again later.",
			Variant:     toast.VariantError,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Success!
	toast.Toast(toast.Props{
		Title:       "You're on the list! 🎉",
		Description: "Check your email for your 50% early bird discount.",
		Variant:     toast.VariantSuccess,
		Position:    "bottom-right",
		Duration:    5000,
		Dismissible: true,
		Icon:        true,
	}).Render(r.Context(), w)
}