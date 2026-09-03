package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Use(AuthMiddleware)

	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}

func parseResponseBodyJson(t *testing.T, resp *http.Response) map[string]interface{} {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal("Failed to read body: ", err)
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		t.Fatal("Failed to parse JSON: ", err)
	}

	return responseBody
}

// Test 1: Should return status code 401 (Unauthorized) when no token sent
func TestAuthMiddleware_UnauthorizedNoToken(t *testing.T) {
	// Setup HTTP
	app := setupApp()

	// Send fake requeset without auth token
	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	// 1 - Should return status code 401 (Unauthorized)
	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("Expected status %v, got %v", fiber.StatusUnauthorized, resp.StatusCode)
	}

	// 2 - Should return correct json response
	responseBody := parseResponseBodyJson(t, resp)

	expectMessage := "Unauthorized"
	if responseBody["message"] != expectMessage {
		t.Errorf("Invalid unauthorized message, Expected: %s, got: %s", expectMessage, responseBody["message"])
	}
}

func TestAuthMiddleware_InvalidTokenFormat(t *testing.T) {
	app := setupApp()

	// Send fake request
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "a")
	resp, _ := app.Test(req)

	responseBody := parseResponseBodyJson(t, resp)

	expectedMessage := "Invalid token format"
	if responseBody["message"] != expectedMessage {
		t.Errorf("Invalid invalid token message, Expected %s, got %s", expectedMessage, responseBody["message"])
	}

}
