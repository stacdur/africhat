package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	db          *mongo.Database
)

func main() {
	// Charger .env
	godotenv.Load()

	// MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB error:", err)
	}
	mongoClient = client
	db = client.Database("africhat")
	log.Println("✅ MongoDB connected")

	// Gin Router
	router := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Routes
	setupRoutes(router)

	// Démarrer le serveur
	port := getEnv("PORT", "8080")
	log.Printf("🚀 AfriChat Server started on port %s", port)
	log.Printf("📡 API: http://localhost:%s/api", port)
	router.Run(":" + port)
}

func setupRoutes(r *gin.Engine) {
	// Routes publiques
	r.POST("/api/auth/register", RegisterHandler)
	r.POST("/api/auth/login", LoginHandler)
	r.POST("/api/auth/request-otp", RequestOTPHandler)
	r.POST("/api/auth/verify-otp", VerifyOTPHandler)

	// Routes protégées
	auth := r.Group("/api")
	auth.Use(AuthMiddleware())
	{
		// Users
		auth.GET("/users/profile", GetProfileHandler)
		auth.PUT("/users/profile", UpdateProfileHandler)

		// Chats
		auth.GET("/chats", GetChatsHandler)
		auth.POST("/chats", CreateChatHandler)
		auth.GET("/chats/:id/messages", GetMessagesHandler)
		auth.POST("/chats/:id/messages", SendMessageHandler)

		// Mobile Money
		auth.POST("/mobile-money/send", SendMoneyHandler)
		auth.GET("/mobile-money/transactions", GetTransactionsHandler)

		// WebRTC Calls
		auth.POST("/webrtc/call/initiate", InitiateCallHandler)
		auth.POST("/webrtc/call/:callId/answer", AnswerCallHandler)
		auth.POST("/webrtc/call/:callId/end", EndCallHandler)
	}

	// WebSocket
	r.GET("/ws", ServeWebSocket)
	r.GET("/api/webrtc/signaling", WebRTCSignalingHandler)

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "africhat-backend",
			"version": "1.0.0",
		})
	})
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// Placeholder handlers (à implémenter)
func RegisterHandler(c *gin.Context)          { c.JSON(200, gin.H{"message": "register"}) }
func LoginHandler(c *gin.Context)             { c.JSON(200, gin.H{"message": "login"}) }
func RequestOTPHandler(c *gin.Context)        { c.JSON(200, gin.H{"message": "otp", "otp": "123456"}) }
func VerifyOTPHandler(c *gin.Context)         { c.JSON(200, gin.H{"message": "verified"}) }
func GetProfileHandler(c *gin.Context)        { c.JSON(200, gin.H{"message": "profile"}) }
func UpdateProfileHandler(c *gin.Context)     { c.JSON(200, gin.H{"message": "updated"}) }
func GetChatsHandler(c *gin.Context)          { c.JSON(200, []string{}) }
func CreateChatHandler(c *gin.Context)        { c.JSON(200, gin.H{"message": "created"}) }
func GetMessagesHandler(c *gin.Context)       { c.JSON(200, []string{}) }
func SendMessageHandler(c *gin.Context)       { c.JSON(200, gin.H{"message": "sent"}) }
func SendMoneyHandler(c *gin.Context)         { c.JSON(200, gin.H{"message": "money sent"}) }
func GetTransactionsHandler(c *gin.Context)   { c.JSON(200, []string{}) }
func InitiateCallHandler(c *gin.Context)      { c.JSON(200, gin.H{"message": "call initiated"}) }
func AnswerCallHandler(c *gin.Context)        { c.JSON(200, gin.H{"message": "answered"}) }
func EndCallHandler(c *gin.Context)           { c.JSON(200, gin.H{"message": "ended"}) }
func AuthMiddleware() gin.HandlerFunc         { return func(c *gin.Context) { c.Next() } }
func ServeWebSocket(c *gin.Context)           { c.String(200, "WebSocket endpoint") }
func WebRTCSignalingHandler(c *gin.Context)   { c.String(200, "WebRTC signaling") }
