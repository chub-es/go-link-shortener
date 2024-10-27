package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "http://app:8000"
	healthPath = host + "/healthz"
	attempts   = 20

	// HTTP REST
	apiPath = host + "/api/v1"
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP GET: /:short_url
func TestHTTPRedirect(t *testing.T) {
	Test(t,
		Description("Redirect Success"),
		Get(host+"/test"),
		Expect().Status().Equal(http.StatusNotFound),
	)
}

// HTTP POST: /url
func TestHTTPDoCreateLink(t *testing.T) {
	body := `{
		"original_url": "https://www.youtube.com/"
	}`
	Test(t,
		Description("DoCreateLink Success"),
		Post(apiPath+"/url"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`"short_url":"`),
	)

	body = `{}`
	Test(t,
		Description("DoCreateLink Fail"),
		Post(apiPath+"/url"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}
