package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID          string    `json:"_id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
}

type itemEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	CreatedAt   time.Time          `bson:"createdAt"`
}

var (
	client     *mongo.Client
	itemColl   *mongo.Collection
	ctxTimeout = 10 * time.Second
	mockItems  = []Item{
		{
			ID:          "1",
			Name:        "手工毛線球",
			Description: "柔軟舒適的毛線球，適合貓咪玩耍",
			Price:       299,
			CreatedAt:   time.Now(),
		},
		{
			ID:          "2",
			Name:        "貓薄荷玩具",
			Description: "含天然貓薄荷，令貓咪歡樂",
			Price:       199,
			CreatedAt:   time.Now(),
		},
	}
)

func main() {
	_ = godotenv.Load()

	port := getEnv("PORT", "3001")
	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI != "" {
		if err := connectMongo(mongoURI); err != nil {
			log.Printf("❌ MongoDB 連接失敗: %v", err)
		} else {
			log.Println("✅ MongoDB 連接成功")
		}
	} else {
		log.Println("⚠️ MONGODB_URI 未設定，將使用模擬資料")
	}

	router := chi.NewRouter()
	router.Use(corsMiddleware)

	router.Get("/api/items", getItems)
	router.Post("/api/items", createItem)
	router.Get("/api/items/{id}", getItemByID)
	router.Put("/api/items/{id}", updateItem)
	router.Delete("/api/items/{id}", deleteItem)
	router.Get("/api/health", healthCheck)

	log.Printf("🚀 API 伺服器運行在 http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func connectMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	itemColl = client.Database("catmicraft").Collection("items")
	return nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			allowed := []string{"http://localhost:5173", "https://yourusername.github.io"}
			for _, o := range allowed {
				if o == origin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getItems(w http.ResponseWriter, r *http.Request) {
	if itemColl == nil {
		respondJSON(w, http.StatusOK, mockItems)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	cursor, err := itemColl.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}))
	if err != nil {
		handleError(w, err)
		return
	}
	defer cursor.Close(ctx)

	var items []Item
	for cursor.Next(ctx) {
		var entity itemEntity
		if err := cursor.Decode(&entity); err != nil {
			handleError(w, err)
			return
		}
		items = append(items, entity.toItem())
	}
	respondJSON(w, http.StatusOK, items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var payload Item
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handleError(w, fmt.Errorf("無效的請求內容: %w", err))
		return
	}

	if strings.TrimSpace(payload.Name) == "" || strings.TrimSpace(payload.Description) == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "名稱和描述為必填項"})
		return
	}

	payload.CreatedAt = time.Now()
	if itemColl == nil {
		payload.ID = fmt.Sprintf("%d", time.Now().UnixNano())
		mockItems = append([]Item{payload}, mockItems...)
		respondJSON(w, http.StatusCreated, payload)
		return
	}

	entity := bson.D{
		{Key: "name", Value: payload.Name},
		{Key: "description", Value: payload.Description},
		{Key: "price", Value: payload.Price},
		{Key: "createdAt", Value: payload.CreatedAt},
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	res, err := itemColl.InsertOne(ctx, entity)
	if err != nil {
		handleError(w, err)
		return
	}

	payload.ID = res.InsertedID.(primitive.ObjectID).Hex()
	respondJSON(w, http.StatusCreated, payload)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if itemColl == nil {
		for _, item := range mockItems {
			if item.ID == id {
				respondJSON(w, http.StatusOK, item)
				return
			}
		}
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "無效的項目 ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	var entity itemEntity
	if err := itemColl.FindOne(ctx, bson.M{"_id": objID}).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
			return
		}
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, entity.toItem())
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload Item
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handleError(w, err)
		return
	}

	if strings.TrimSpace(payload.Name) == "" || strings.TrimSpace(payload.Description) == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "名稱和描述為必填項"})
		return
	}

	if itemColl == nil {
		for index, item := range mockItems {
			if item.ID == id {
				mockItems[index].Name = payload.Name
				mockItems[index].Description = payload.Description
				mockItems[index].Price = payload.Price
				respondJSON(w, http.StatusOK, mockItems[index])
				return
			}
		}
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "無效的項目 ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	update := bson.M{
		"name":        payload.Name,
		"description": payload.Description,
		"price":       payload.Price,
	}

	result := itemColl.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var entity itemEntity
	if err := result.Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
			return
		}
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, entity.toItem())
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if itemColl == nil {
		for index, item := range mockItems {
			if item.ID == id {
				deleted := mockItems[index]
				mockItems = append(mockItems[:index], mockItems[index+1:]...)
				respondJSON(w, http.StatusOK, map[string]any{"message": "項目已刪除", "item": deleted})
				return
			}
		}
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "無效的項目 ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	result, err := itemColl.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		handleError(w, err)
		return
	}
	if result.DeletedCount == 0 {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "項目未找到"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "項目已刪除"})
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	status := "未連接"
	if itemColl != nil {
		status = "連接成功"
	}
	respondJSON(w, http.StatusOK, map[string]any{
		"status":  "OK",
		"timestamp": time.Now().Format(time.RFC3339),
		"mongodb":  status,
	})
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("error: %v", err)
	respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func (e itemEntity) toItem() Item {
	return Item{
		ID:          e.ID.Hex(),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CreatedAt:   e.CreatedAt,
	}
}
