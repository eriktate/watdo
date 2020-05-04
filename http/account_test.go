package http_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eriktate/wrkhub"
	whttp "github.com/eriktate/wrkhub/http"
	"github.com/eriktate/wrkhub/mock"
	"github.com/eriktate/wrkhub/uid"
	"github.com/sirupsen/logrus"
)

func Test_PostAccount(t *testing.T) {
	id := uid.New()
	cases := []struct {
		name         string
		body         string
		service      wrkhub.AccountService
		expectedCode int
		expectedBody string
	}{
		{
			name: "valid account creation",
			body: `{"name": "Test Account"}`,
			service: &mock.AccountService{
				SaveAccountFn: func(ctc context.Context, account wrkhub.Account) (uid.UID, error) {
					return id, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedBody: id.JSONString(),
		},
		{
			name: "valid account update",
			body: fmt.Sprintf(`{"id": "%s", "name": "Test Account"}`, id.String()),
			service: &mock.AccountService{
				SaveAccountFn: func(ctc context.Context, account wrkhub.Account) (uid.UID, error) {
					return id, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedBody: id.JSONString(),
		},
		{
			name: "invalid json",
			body: `{"broken": }`,
			service: &mock.AccountService{
				SaveAccountFn: func(ctc context.Context, account wrkhub.Account) (uid.UID, error) {
					return id, nil
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: "could not unmarshal account",
		},
		{
			name: "service error",
			body: `{"name": "Test Account"}`,
			service: &mock.AccountService{
				Error: errors.New("forced error"),
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "could not save account",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBufferString(c.body))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			whttp.PostAccount(c.service, logrus.New())(rec, req)

			if rec.Code != c.expectedCode {
				t.Fatalf("expected status code of %d but got %d", c.expectedCode, rec.Code)
			}

			if rec.Body.String() != c.expectedBody {
				t.Fatalf("body did not match expectation")
			}
		})
	}
}
