package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func GetAdvertByIdCount(db *sql.DB, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		advertID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logger.Info("Invalid advert ID")
			http.Error(w, "Invalid advert ID", http.StatusBadRequest)
			return
		}

		// Выполняем запрос к базе данных для получения объявления
		row := db.QueryRow("SELECT id, user_id, type_placement, title, description, phone, is_agent, priority, created_at, is_deleted FROM advert WHERE id = $1", advertID)

		var advert models.Advert
		err = row.Scan(&advert.ID, &advert.UserID, &advert.AdvertTypeSale, &advert.Title, &advert.Description, &advert.Phone, &advert.IsAgent, &advert.Priority, &advert.DateCreation, &advert.IsDeleted)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Info("Advert not found")
				http.Error(w, "Advert not found", http.StatusBadRequest)
				return
			}
			logger.Info("Error getting advert")
			http.Error(w, fmt.Sprintf("Error getting advert: %v", err), http.StatusInternalServerError)
			return
		}

		// Получаем количество лайков для данного объявления
		var likesCount int
		err = db.QueryRow("SELECT COUNT(*) FROM favourite_advert WHERE advert_id = $1 AND is_deleted = false", advertID).Scan(&likesCount)
		if err != nil && err != sql.ErrNoRows {
			logger.Info("Error getting likes count")
			http.Error(w, fmt.Sprintf("Error getting likes count: %v", err), http.StatusInternalServerError)
			return
		}

		// Добавляем количество лайков к объявлению
		advert.Likes = likesCount

		// Отправляем ответ в формате JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(advert)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting likes count: %v", err), http.StatusInternalServerError)
		}
	}
}

func GetAdvertById(db *sql.DB, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		advertID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logger.Info("Invalid advert ID")
			http.Error(w, "Invalid advert ID", http.StatusBadRequest)
			return
		}

		// Выполняем запрос к базе данных для получения объявления
		row := db.QueryRow("SELECT id, user_id, type_placement, title, description, phone, is_agent, priority, likes, created_at,  is_deleted FROM advert WHERE id = $1", advertID)

		var advert models.Advert
		err = row.Scan(&advert.ID, &advert.UserID, &advert.AdvertTypeSale, &advert.Title, &advert.Description, &advert.Phone, &advert.IsAgent, &advert.Priority, &advert.Likes, &advert.DateCreation, &advert.IsDeleted)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Info("Advert not found")
				http.Error(w, "Advert not found", http.StatusNotFound)
				return
			}
			logger.Info("Error getting likes count")
			http.Error(w, fmt.Sprintf("Error getting advert: %v", err), http.StatusInternalServerError)
			return
		}

		// Отправляем ответ в формате JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(advert)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting likes count: %v", err), http.StatusInternalServerError)
		}
	}
}
