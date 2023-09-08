package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "bytes"

    "github.com/gin-gonic/gin"
)

func TestRegisterUserEndpoint(t *testing.T) {
    // Create a new HTTP request to register a user.
    userRegistration := map[string]interface{}{
        "name":   "TestUser",
        "email":  "test@example.com",
        "passwd": "test123",
    }
    requestBody, _ := json.Marshal(userRegistration)
    req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to capture the response.
    rr := httptest.NewRecorder()

    // Create a Gin router and set up your routes and handlers.
    r := gin.Default()
    handlers := newHandlers(nil) // You can provide a mock database if needed.
    r.POST("/register", handlers.registerUser)

    // Serve the request to your handler.
    r.ServeHTTP(rr, req)

    // Check the response status code.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v but got %v", http.StatusOK, status)
    }
}
