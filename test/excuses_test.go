package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/mantil-io/template-excuses/api/excuses"
)

func TestExcuses(t *testing.T) {
	api := httpexpect.New(t, apiURL)

	// clear api state
	api.POST("/excuses/clear").
		Expect().
		Status(http.StatusNoContent)

	// check that the count is 0
	api.POST("/excuses/count").
		Expect().
		ContentType("application/json").
		Status(http.StatusOK).
		JSON().Object().
		Value("Count").Number().Equal(0)

	// check that random returns error
	api.POST("/excuses/random").
		Expect().
		ContentType("application/json").
		Status(http.StatusInternalServerError).
		Header("x-api-error").Equal("no excuses")

	// call load with invalid URL and expect error
	req := excuses.LoadRequest{
		URL: "https://this.is.not.valid.url",
	}
	api.POST("/excuses/load").
		WithJSON(req).
		Expect().
		ContentType("application/json").
		Status(http.StatusInternalServerError).
		Header("x-api-error").Contains("no such host")

	// call load with valid url and expect valid countd
	req = excuses.LoadRequest{
		URL: "https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt",
	}
	api.POST("/excuses/load").
		WithJSON(req).
		Expect().
		ContentType("application/json").
		Status(http.StatusOK).
		JSON().Object().
		Value("Count").Number().Equal(63)

	// get value
	obj := api.POST("/excuses/random").
		Expect().
		ContentType("application/json").
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("Excuse")

	t.Logf("got excuse: %s", obj.Value("Excuse").String().Raw())

	// get value from default method
	api.POST("/excuses").
		Expect().
		ContentType("application/json").
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("Excuse")

	// method which don't exists
	api.POST("/excuses/non-existent-method").
		Expect().
		Status(http.StatusNotImplemented)
}
