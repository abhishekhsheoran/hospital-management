package doctor

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hospital-management/utils"
)

func TestCreatDoctorPass(t *testing.T) {
	// prepare input params
	// Prepare req :
	// 1. Preapre req body
	var body = `
	{
		"name":"Vipin",
		"age" : 31
	}	
	`
	utils.InitializeDatabase()
	var jsonStr = []byte(body)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest(http.MethodPost, "http://localhot:8080/api/v1/doctor", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("expected no err but got %v", err)
	}
	// create response obj
	response := httptest.NewRecorder()
	// call function
	CreateDoctor(response, req)
	if response.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected 200 but got %v", response.Result().StatusCode)
	}
}

func TestCreatDoctorFail(t *testing.T) {
	// prepare input params
	// Prepare req :
	// 1. Preapre req body
	var body = `
	`
	utils.InitializeDatabase()
	var jsonStr = []byte(body)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest(http.MethodPost, "http://localhot:8080/api/v1/doctor", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("expected no err but got %v", err)
	}
	// create response obj
	response := httptest.NewRecorder()
	// call function
	CreateDoctor(response, req)
	if response.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected 200 but got %v", response.Result().StatusCode)
	}
}
