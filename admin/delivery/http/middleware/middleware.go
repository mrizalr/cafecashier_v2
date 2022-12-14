package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
)

func BasicAuth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SetContentTypeJSON(w)
		token := r.Header.Get("token")

		jsonData, err := utils.Decode(token)
		if err != nil {
			log.Println(err.Error())
			err = errors.New("you are not authorized, please do some login first")

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.ResponseFormatter(http.StatusUnauthorized, "UNAUTHORIZED", "errors", err.Error()))
			return
		}

		adminData := models.AdminDataToken{}
		err = json.Unmarshal(jsonData, &adminData)
		if err != nil {
			log.Println(err.Error())
			err = errors.New("you are not authorized, please do some login first")

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.ResponseFormatter(http.StatusUnauthorized, "UNAUTHORIZED", "errors", err.Error()))
			return
		}

		ctx := context.WithValue(context.Background(), "admin-data", adminData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
