package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var jwtKey = []byte("your_secret_key") // Lembre-se de substituir pela sua chave real

func main() {
	InitDB("root:0000@tcp(localhost:3306)/testdb") // Ajuste a conexão conforme necessário

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

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Println("Connected to database")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	// Implementação do registro de usuários
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implementação da validação do usuário no banco de dados e geração do token JWT
	if user.Username == "usuario" && user.Password == "senha" {
		token, err := GenerateJWT(user.Username)
		if err != nil {
			http.Error(w, "Error generating JWT", http.StatusInternalServerError)
			return
		}

		// Retornar o token JWT para o cliente
		json.NewEncoder(w).Encode(TokenResponse{Token: token})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT PartnerId, PartnerName, CustomerId, CustomerName FROM myTable")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []map[string]interface{}
	for rows.Next() {
		var PartnerId, PartnerName, CustomerId, CustomerName string
		err := rows.Scan(&PartnerId, &PartnerName, &CustomerId, &CustomerName)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning database row", http.StatusInternalServerError)
			return
		}
		row := map[string]interface{}{
			"PartnerId":    PartnerId,
			"PartnerName":  PartnerName,
			"CustomerId":   CustomerId,
			"CustomerName": CustomerName,
		}
		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over database rows", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Token expira em 2 horas
	})

	return token.SignedString(jwtKey)
}
