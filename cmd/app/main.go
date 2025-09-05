package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/handlers/api"
	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/models"
	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/services"
)

type User1 struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	TelegramId string `json:"telegramId"`
}

type Currency struct {
	Bank        string `json:"bank"`
	Symbol      string `json:"symbol"`
	SumCurrency any    `json:"sumCurrency"`
}

type Crypt struct {
	FullName  string `json:"fullName"`
	Name      string `json:"name"`
	SumCrypto any    `json:"sumCrypto"`
}

type ExchangeTransfer struct {
	CryptoWallet string   `json:"cryptoWallet"`
	BankCard     string   `json:"bankCard"`
	Crypt        Crypt    `json:"crypt"`
	Currency     Currency `json:"currency"`
}

type CreateRequestRequest struct {
	User     User1            `json:"user"`
	Transfer ExchangeTransfer `json:"transfer"`
}

type CreateRequestRegistrationRequest struct {
	TelegramId string           `json:"telegramId"`
	Transfer   ExchangeTransfer `json:"transfer"`
}

func main() {
	database := services.DataBaseService{}

	loginError := database.Connect("root", "admin", "databaseusers")
	if loginError != nil {
		return
	}
	defer func(database services.DataBaseService) {
		closeError := database.Close()
		if closeError != nil {
			return
		}
	}(database)

	//http.FileServer(http.Dir("C:\\Users\\nikit\\WebstormProjects\\CurrencyExchangerProjectFrontend\\dist\\CurrencyExchangerProjectFrontend\\browser"))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	// Если запрос к API, то 404
	//	if strings.HasPrefix(r.URL.Path, "/api/") {
	//		http.NotFound(w, r)
	//		return
	//	}
	//
	//	// Если файл не существует - отдаем index.html (для SPA routing)
	//	indexFile := "C:\\Users\\nikit\\WebstormProjects\\CurrencyExchangerProjectFrontend\\dist\\CurrencyExchangerProjectFrontend\\browser\\index.html"
	//	if _, err := os.Stat(indexFile); err == nil {
	//		http.ServeFile(w, r, indexFile)
	//		return
	//	}
	//
	//	http.NotFound(w, r)
	//})

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Query().Get("operation") == "login" {
			api.ChekUser(w, r)
		} else if r.Method == "POST" && r.URL.Query().Get("operation") == "create" {
			err := database.Insert("users", api.CreateUser(w, r))
			if err != nil {
				return
			}
		} else if r.Method == "GET" && r.URL.Query().Get("operation") == "status" {
			api.GetStatusUser(w, r)
		} else if r.Method == "GET" && r.URL.Query().Get("operation") == "logout" {
			api.GetPermissionLogout(w, r)
		} else if r.Method == "POST" && r.URL.Query().Get("operation") == "recovery" {
			type Sdd struct {
				EmailOrLogin string `json:"input_login_or_email"`
			}
			var req Sdd
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println(err)
			}
			fmt.Println("ddd ", req.EmailOrLogin)
		} else if r.Method == "POST" && r.URL.Query().Get("operation") == "sendCodeEmail" {
			type Sdd struct {
				Code string `json:"code"`
			}
			var req Sdd
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println(err)
			}
			fmt.Println("ddd ", req.Code)

			w.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Code:    "USER_SUCCESS",
				Message: "text",
			})
		} else if r.Method == "POST" && r.URL.Query().Get("operation") == "sendNewPassword" {
			type Sdd struct {
				NewPassword string `json:"new_password"`
			}
			var req Sdd
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println(err)
			}
			fmt.Println("ddd ", req.NewPassword)
		}
	})

	http.HandleFunc("/api/create_request", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())
		if r.URL.Query().Get("registration") == "true" {
			var req CreateRequestRegistrationRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println(err)
			}
			fmt.Println(req)
		} else if r.URL.Query().Get("registration") == "false" {
			var req CreateRequestRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println(err)
			}
			fmt.Println(req)
		}
	})

	server := &http.Server{
		Addr:                         ":8080",
		DisableGeneralOptionsHandler: true,
	}

	log.Fatal(server.ListenAndServe())
}
