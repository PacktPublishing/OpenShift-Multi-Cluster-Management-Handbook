// main_test.go

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	initCloudsObject()
	code := m.Run()
	os.Exit(code)
}

func TestReturnAllClouds(t *testing.T) {

	req, err := http.NewRequest("GET", "/cloud", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllClouds)
	handler.ServeHTTP(rr, req)

	expected := `[{"Id":"1","Title":"AWS","desc":"Amazon Web Services","content":"Amazon Web Services, Inc. is a subsidiary of Amazon providing on-demand cloud computing platforms and APIs to individuals, companies, and governments, on a metered pay-as-you-go basis."},{"Id":"2","Title":"Azure","desc":"Microsoft Azure","content":"Microsoft Azure, often referred to as Azure, is a cloud computing service operated by Microsoft for application management via Microsoft-managed data centers."},{"Id":"3","Title":"GCP","desc":"Google Cloud Platform","content":"Google Cloud Platform, offered by Google, is a suite of cloud computing services that runs on the same infrastructure that Google uses internally for its end-user products, such as Google Search, Gmail, Google Drive, and YouTube."},{"Id":"4","Title":"IBM Cloud","desc":"IBM Cloud","content":"IBM cloud computing is a set of cloud computing services for business offered by the information technology company IBM."}]`

	checkResponseCode(t, http.StatusOK, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func TestReturnSingleCloud(t *testing.T) {

	req, err := http.NewRequest("GET", "/cloud/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	returnSingleCloud(rr, req)

	expected := `{"Id":"1","Title":"AWS","desc":"Amazon Web Services","content":"Amazon Web Services, Inc. is a subsidiary of Amazon providing on-demand cloud computing platforms and APIs to individuals, companies, and governments, on a metered pay-as-you-go basis."}`

	checkResponseCode(t, http.StatusOK, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func TestReturnSingleCloudNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/cloud/9999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "99999",
	}
	req = mux.SetURLVars(req, vars)

	returnSingleCloud(rr, req)

	expected := `{"Id":"","Title":"Not Found","desc":"Not Found","content":"Not Found"}`

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func TestCreateNewCloud(t *testing.T) {

	expected := `{"Id":"5","Title":"Alibaba","desc":"Alibaba Cloud","content":"Alibaba Cloud"}`
	var jsonStr = []byte(expected)

	req, err := http.NewRequest("POST", "/cloud", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewCloud)
	handler.ServeHTTP(rr, req)

	checkResponseCode(t, http.StatusOK, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func TestUpdateCloud(t *testing.T) {
	expected := `{"Id":"5","Title":"Alibaba v2","desc":"Alibaba Cloud v2","content":"Alibaba Cloud v2"}`
	var jsonStr = []byte(expected)

	req, err := http.NewRequest("PUT", "/cloud", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewCloud)
	handler.ServeHTTP(rr, req)

	checkResponseCode(t, http.StatusOK, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func TestDeleteCloud(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/cloud/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "5",
	}
	req = mux.SetURLVars(req, vars)

	deleteCloud(rr, req)

	expected := ``

	checkResponseCode(t, http.StatusOK, rr.Code)
	checkResponseBody(t, expected, rr.Body.String())
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected string, actual string) {
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected %s array. Got %s", expected, actual)
	}
}
