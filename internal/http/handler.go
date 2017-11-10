package http

import (
	"context"
	"net/http"

	"github.com/denniszl/go-debug/internal/data"
	"code.cfops.it/bill/entitlements-api/internal/entitlements"
	"code.cfops.it/bill/go-lib/logger"
	"code.cfops.it/bill/go-lib/metrics"
	"github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler initializes a go-kit http service
func MakeHTTPHandler(
	endpoints entitlements.Endpoints,
	auth Auth,
) http.Handler {
	r := mux.NewRouter()
	// Server options for all endpoints

	// GET /thing
	thing := endpoints.Healthcheck
	r.Methods("GET").Path("/thing").Handler(
		httptransport.NewServer(
			thing,
			decodeRequest,
			encodeResponse,
			options...,
		),
	).Name("thing")

	// Not Found
	notFound := func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, ErrNotFound
	}
	notFound = loggingMiddleware(logger, "NotFound")(notFound)
	r.NotFoundHandler = httptransport.NewServer(
		notFound,
		decodeRequest,
		encodeResponse,
		options...,
	)

	return r
}
