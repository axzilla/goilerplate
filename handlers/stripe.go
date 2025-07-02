package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(os.Getenv("STRIPE_PRICE_ID")),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(fmt.Sprintf("%s/success?session_id={CHECKOUT_SESSION_ID}", os.Getenv("BASE_URL"))),
		CancelURL:  stripe.String(fmt.Sprintf("%s/#pricing", os.Getenv("BASE_URL"))),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		http.Error(w, "Failed to create checkout session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func HandleSuccess(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	
	s, err := session.Get(sessionID, nil)
	if err != nil {
		log.Printf("session.Get: %v", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	SuccessPage(s.CustomerDetails.Email).Render(r.Context(), w)
}