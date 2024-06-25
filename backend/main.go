package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func main() {
	users = append(users, User{
		ID:       1,
		Email:    "test@example.com",
		Password: hashPassword("password123"),
	})

	r := mux.NewRouter()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	r.Use(corsMiddleware)

	r.HandleFunc("/api/register", registerUser).Methods("POST")
	r.HandleFunc("/api/login", loginUser).Methods("POST")
	r.HandleFunc("/api/data", authMiddleware(dataHandler)).Methods("GET")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Register request received")
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("Error decoding new user:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == newUser.Email {
			log.Println("User already exists:", newUser.Email)
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}

	hashedPassword := hashPassword(newUser.Password)
	newUser.Password = hashedPassword
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	log.Println("New user registered:", newUser.Email)
	json.NewEncoder(w).Encode(newUser)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Login request received")
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Println("Error decoding credentials:", err)
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	log.Println("Credentials decoded:", creds.Email)

	var user User
	for _, u := range users {
		if u.Email == creds.Email {
			user = u
			break
		}
	}

	if user.Email == "" {
		log.Println("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		log.Println("Invalid password")
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error creating token:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Login successful, returning token")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Data from server"}
	json.NewEncoder(w).Encode(data)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
