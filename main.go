package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	v4 "github.com/aws/aws-sdk-go-v2/signer/v4"
)

type sigV4RoundTripper struct {
	transport http.RoundTripper
	signer    *v4.Signer
}

func (rt *sigV4RoundTripper) calculatePayloadHash(req *http.Request) (string, error) {
	// Clone the request body
	var buf bytes.Buffer
	bodyClone := io.TeeReader(req.Body, &buf)
	req.Body = ioutil.NopCloser(bodyClone)

	// Read the cloned request body
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read request body: %v", err)
	}

	// Calculate the payload hash
	hash := sha256.New()
	hash.Write(bodyBytes)
	payloadHash := hex.EncodeToString(hash.Sum(nil))

	return payloadHash, nil
}

func (rt *sigV4RoundTripper) signRequest(ctx context.Context, req *http.Request) error {
	payloadHash, err := rt.calculatePayloadHash(req)
	if err != nil {
		return err
	}

	// Sign the request using AWS SDK for Go v2 SigV4 signer
	err = rt.signer.SignHTTP(ctx, aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return rt.signer.Credentials, nil
	}), req, payloadHash, "execute-api", "", v4.SignHTTPRequestPayload)
	if err != nil {
		return fmt.Errorf("failed to sign request: %v", err)
	}

	return nil
}

func (rt *sigV4RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Clone the request
	clone := req.Clone(ctx)

	// Sign the cloned request
	err := rt.signRequest(ctx, clone)
	if err != nil {
		return nil, err
	}

	// Send the signed request using the underlying transport
	return rt.transport.RoundTrip(clone)
}

// ... main function
package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/credentials"
	v4 "github.com/aws/aws-sdk-go-v2/signer/v4"
	"github.com/stretchr/testify/assert"
)

func TestSigV4RoundTripper(t *testing.T) {
	// Mock API Gateway REST API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Dummy AWS credentials for testing
	accessKeyID := "AKIATESTKEYID"
	secretAccessKey := "testSecretAccessKey"
	sessionToken := ""

	// Create a custom RoundTripper
	rt := &sigV4RoundTripper{
		transport: http.DefaultTransport,
		signer: v4.NewSigner(credentials.NewStaticCredentialsProvider(
			accessKeyID,
			secretAccessKey,
			sessionToken,
		)),
	}

	client := &http.Client{
		Transport: rt,
	}

	payload := `{"Key1":"Value1","Key2":"Value2"}`
	req, _ := http.NewRequest("POST", ts.URL, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify the presence of the expected authorization header
	assert.NotEmpty(t, resp.Request.Header.Get("Authorization"))
}

func TestCalculatePayloadHash(t *testing.T) {
	rt := &sigV4RoundTripper{}

	// Create a dummy request with a body
	payload := `{"Key1":"Value1","Key2":"Value2"}`
	req, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(payload))

	// Calculate the payload hash using the calculatePayloadHash function
	payloadHash, err := rt.calculatePayloadHash(req)
	assert.NoError(t, err)

	// Calculate the expected payload hash manually
	hash := sha256.New()
	hash.Write([]byte(payload))
	expectedPayloadHash := hex.EncodeToString(hash.Sum(nil))

	// Verify that the payload hash matches the expected hash
	assert.Equal(t, expectedPayloadHash, payloadHash)
}
