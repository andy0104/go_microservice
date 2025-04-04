package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"go_microservice/user_service/config"
	"go_microservice/user_service/dto"
	"go_microservice/user_service/handlers"
	"go_microservice/user_service/repository"
	"go_microservice/user_service/services"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"go.uber.org/zap"
)

func TestUserSignUpApi(t *testing.T) {
	// set up mock database
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	//initialize the logger
	logger := zap.Must(zap.NewDevelopment()).Sugar()

	// initialize the repositories
	repo := repository.NewRepository(db)

	// initialize the services
	svcs := services.NewServices(repo, logger)

	// initialize the handlers
	hndlrs := handlers.NewIndexHandler(svcs)

	// mock sqlx db
	// mockDb := sqlx.NewDb(db, "sqlmock")
	_ = mock
	// _ = mockDb

	app := fiber.New()
	// setup app routes
	config.InitServer(app, hndlrs)

	t.Run("test signup api call first name validation error", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "",
			LastName:  "Kar",
			Email:     "a@c.com",
			Password:  "1234567890",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
		assert.Nil(t, response["data"])
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "validation error")
	})

	t.Run("test signup api call last name validation error", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "",
			Email:     "a@c.com",
			Password:  "1234567890",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
		assert.Nil(t, response["data"])
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "validation error")
	})

	t.Run("test signup api call email validation error", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "Kar",
			Email:     "",
			Password:  "1234567890",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
		assert.Nil(t, response["data"])
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "validation error")
	})

	t.Run("test signup api call invalid email validation error", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "Kar",
			Email:     "a.com",
			Password:  "1234567890",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
		assert.Nil(t, response["data"])
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "validation error")
	})

	t.Run("test signup api call password validation error", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "Kar",
			Email:     "a@b.com",
			Password:  "",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
		assert.Nil(t, response["data"])
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "validation error")
	})

	t.Run("test signup api call email exist error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT user_id, first_name, last_name, email_id, user_password FROM "Users" WHERE email_id = \$1`).
			WithArgs("a@c.com").
			WillReturnRows(sqlxmock.NewRows([]string{"user_id", "first_name", "last_name", "email_id", "user_password"}).
				AddRow(1, "Aninda", "Kar", "a@c.com", "1234567890"))

		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "Kar",
			Email:     "a@c.com",
			Password:  "1234567890",
		}
		jsonReq, _ := json.Marshal(signupRequest)
		bodyReq := bytes.NewBuffer(jsonReq)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bodyReq)
		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Could not create signup request: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}
		var response map[string]map[string]any
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusConflict)
		assert.NotNil(t, response["error"])
		assert.Equal(t, response["error"]["message"], "email is already in use")
		assert.Nil(t, response["data"])
	})

	t.Run("test signup api call success", func(t *testing.T) {
		// create signup request payload
		signupRequest := dto.UserSignupRequest{
			FirstName: "Aninda",
			LastName:  "Kar",
			Email:     "a@c.com",
			Password:  "1234567890",
		}

		// mock get email by id call
		mock.ExpectQuery(`SELECT user_id, first_name, last_name, email_id, user_password FROM "Users" WHERE email_id = \$1`).
			WithArgs("a@c.com").
			WillReturnError(sql.ErrNoRows)

		// mock the create user call
		mock.ExpectExec(`INSERT INTO "Users"`).
			WithArgs("Aninda", "Kar", "a@c.com", sqlxmock.AnyArg()).
			WillReturnResult(sqlxmock.NewResult(1, 1))

		jsonReq, _ := json.Marshal(signupRequest)
		req, err := http.NewRequest("POST", "/api/v1/user/signup", bytes.NewBuffer(jsonReq))
		if err != nil {
			t.Fatalf("Could not create new request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("Signup request error: %v", err)
		}

		// parse the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response body: %v", err)
		}

		var response map[string]map[string]any
		if err = json.Unmarshal(bodyBytes, &response); err != nil {
			t.Fatalf("Could not parse the response body: %v", err)
		}

		assert.Equal(t, res.StatusCode, fiber.StatusCreated)
		assert.NotNil(t, response["data"])
		assert.Equal(t, response["data"]["message"], "user signup")
		assert.Nil(t, response["error"])
	})
}
