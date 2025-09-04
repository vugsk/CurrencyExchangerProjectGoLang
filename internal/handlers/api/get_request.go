package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/models"
)

func GetStatusUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth-token")
	if err != nil {
		fmt.Println("auth-token cookie not found")
		w.WriteHeader(401)
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Code:    "USER_EXISTING",
			Message: "text",
		})
		return
	}
	fmt.Print("status user: ")
	fmt.Println(cookie)
	fmt.Println(r.Cookie("auth-token"))
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Code:    "USER_SESSION_ID_VALID",
		Message: "text",
	})
}

func GetPermissionLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Print("logout ")
	fmt.Println(r.Cookie("auth-token"))
	cookie := &http.Cookie{
		Name:   "auth-token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Code:    "USER_SUCCESS_LOGOUT",
		Message: "text",
	})
}
