package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DateKey int = iota
	FromKey
	ToKey
	IsWeekAverageKey
	JSONContent

	FromQS = "from"
	ToQS   = "to"
	DateQS = "date"
)

var ContextQueryStringList = map[int]string{
	DateKey: DateQS,
	FromKey: FromQS,
	ToKey:   ToQS,
}

// JSONToContext value
func JSONToContext(ctx context.Context, req *http.Request) context.Context {

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	return context.WithValue(ctx, JSONContent, body)
}

// JSONFromContext value
func JSONFromContext(ctx context.Context, i interface{}) bool {
	b, ok := ctx.Value(JSONContent).([]byte)
	if !ok {
		return false
	}

	if err := json.Unmarshal(b, i); err != nil {
		log.Println("func JSONFromContext", err)
		return false
	}
	return true
}
func DateKeyFromContext(ctx context.Context) (string, bool) {
	ua, ok := ctx.Value(DateKey).(string)
	return ua, ok
}

func FromKeyFromContext(ctx context.Context) (string, bool) {
	ua, ok := ctx.Value(FromKey).(string)
	return ua, ok
}

func ToKeyFromContext(ctx context.Context) (string, bool) {
	ua, ok := ctx.Value(ToKey).(string)
	return ua, ok
}

func AllQSToContext(ctx context.Context, req *http.Request) context.Context {
	for key, qs := range ContextQueryStringList {
		ctx = context.WithValue(ctx, key, req.FormValue(qs))
	}
	return ctx
}

func GetContextFromRequest(req *http.Request, opt ...func(context.Context, *http.Request) context.Context) context.Context {
	ctx := req.Context()

	for _, f := range opt {
		ctx = f(ctx, req)
	}

	return ctx
}

func GetAllContextFromRequest(req *http.Request) context.Context {
	contentType := req.Header.Get("Content-Type")
	if contentType == "application/vnd.api+json" || contentType == "application/json" {
		return GetContextFromRequest(req, JSONToContext)
	}
	return GetContextFromRequest(req, AllQSToContext, JSONToContext)
}
