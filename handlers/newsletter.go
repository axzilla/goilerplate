package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/axzilla/goilerplate/ui/components/toast"
)

func HandleNewsletter(w http.ResponseWriter, r *http.Request) {
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

	apiKey := os.Getenv("BREVO_API_KEY")
	
	if apiKey == "" {
		log.Printf("Newsletter signup (dev mode): %s", email)
		
		toast.Toast(toast.Props{
			Title:       "You're on the list! 🎉",
			Description: "Thanks for your interest.",
			Variant:     toast.VariantSuccess,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	contactData := map[string]any{
		"email":   email,
		"listIds": []int{2},
	}

	jsonData, err := json.Marshal(contactData)
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

	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/contacts", bytes.NewBuffer(jsonData))
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

	req.Header.Set("api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

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

	body, _ := io.ReadAll(resp.Body)
	
	if resp.StatusCode == 201 {
		toast.Toast(toast.Props{
			Title:       "You're on the list! 🎉",
			Description: "Thanks for your interest.",
			Variant:     toast.VariantSuccess,
			Position:    "bottom-right",
			Duration:    5000,
			Dismissible: true,
			Icon:        true,
		}).Render(r.Context(), w)
		return
	}

	// Brevo returns 400 with "Contact already exist" message
	if resp.StatusCode == 400 {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if code, ok := errorResp["code"].(string); ok && code == "duplicate_parameter" {
				toast.Toast(toast.Props{
					Title:       "You're already on the list!",
					Description: "Thanks for your continued interest.",
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

	log.Printf("Brevo response: Status %d, Body: %s", resp.StatusCode, string(body))
	toast.Toast(toast.Props{
		Title:       "Something went wrong",
		Description: "Please try again later.",
		Variant:     toast.VariantError,
		Position:    "bottom-right",
		Duration:    5000,
		Dismissible: true,
		Icon:        true,
	}).Render(r.Context(), w)
}