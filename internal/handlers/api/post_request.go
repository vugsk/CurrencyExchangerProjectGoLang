package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/models"
)

func sendError(w http.ResponseWriter, err models.ErrorResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err.Message)
}

func CreateUser(w http.ResponseWriter, r *http.Request) interface{} {
	var req models.RequestRegistration
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		var responsesError models.ErrorResponse = models.ErrorResponse{
			Message: "Запрос не содержит необходимых данных",
			Code:    "NO_NECESSARY_DATA",
		}
		sendError(w, responsesError, http.StatusBadRequest)
		return nil
	}

	if req.Email == "" || req.Login == "" || req.Password == "" {
		var responsesError models.ErrorResponse = models.ErrorResponse{
			Message: "Данные есть, но не прошли валидацию",
			Code:    "FAILED_VALIDATION",
		}
		sendError(w, responsesError, http.StatusUnprocessableEntity)
		return nil
	}

	fmt.Println(req.Email, req.Login, req.Password)
	if req.Login == "kollok" {
		w.WriteHeader(http.StatusConflict) // 409 - Conflict
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Пользователь с таким login уже существует",
			Code:    "LOGIN_EXISTS",
		})
	}
	return req
}

func ChekUser(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		fmt.Println(r.Method)
		return false
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		return false
	}

	if req.Password == "" {
		return false
	}

	fmt.Println(req.LoginOrEmail, req.Password)

	if req.LoginOrEmail == "kollok" && req.Password == req.Password {
		fmt.Println("GetCookieProfile")
		cookie := http.Cookie{
			Name:     "auth-token",
			Value:    "h34fdf62df7f3h34fdf62df7f3h34fdf",
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			User    struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"user"`
			Timestamp time.Time `json:"timestamp"`
		}{
			Message: "ok",
			User: struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			}{Name: "kollok", Id: "df62df7f3h34f"},
			Timestamp: time.Now(),
		})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusLocked)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Wrong password",
			Code:    "PASSWORD_NOT_FOUND",
		})
		return false
	}
	return true
}
