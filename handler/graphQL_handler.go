package handler

import (
	"net/http"
	"github.com/graphql-go/handler"
	"github.com/iqra-sham/jokes/schema"
)

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})
	h.ServeHTTP(w, r)
}
