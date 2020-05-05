package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/eriktate/wrkhub/env"
	"github.com/eriktate/wrkhub/uid"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// const defaultKey = "0Go7N9j2SEvU6R14pDtakSQfANwRF5Va3YGHb9COqFc6VA4cMRmrt31fSAASxuz9lWH/3kj3cU6OMtv0se1Fap4ciTSefkIy7MNFcmEyc6+GEAAz1Cd9KfJG3+MeZgGcG77lz0pAfhd6sDiRn+y7oWJyMIweYizwtedgC9adZvgwh8Mss41yss5tfjrI1SBtIhgvj361fhLojrO8nZHXVaqlLW13xSxcuGCe9Q=="
const defaultKey = "supersecret"

const tokenPrefix = "Bearer "

type sessionCtxKey string

const sessionKey = sessionCtxKey("sessionCtxKey")

type Session struct {
	UserID uid.UID
}

// GetSession retrieves a session token from the context.
func GetSession(ctx context.Context) (Session, error) {
	if session, ok := ctx.Value(sessionKey).(Session); ok {
		return session, nil
	}

	return Session{}, errors.New("session not found")
}

func CtxWithSession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func Authenticate(log *logrus.Logger) Middleware {
	signingSecret := env.GetString("WRKHUB_SIGNING_SECRET", defaultKey)
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// short circuit if getting a token
			if strings.Contains(r.URL.Path, "token") {
				handler.ServeHTTP(w, r)
				return
			}

			// tokHeader := r.Header.Get("Authorization")
			// if len(tokHeader) <= len(tokenPrefix) {
			// 	log.WithField("token", tokHeader).Error("auth token too short")
			// 	badRequest(w, "could not parse auth token")
			// 	return
			// }

			log.Info(r.Cookies())
			sessionCookie, err := r.Cookie("wrkhub-session-token")
			if err != nil {
				log.WithError(err).Error("could not read session token")
				badRequest(w, "could not read session token")
				return
			}

			tokHeader := sessionCookie.Value

			token, err := jwt.Parse(tokHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(signingSecret), nil
			})

			if err != nil {
				log.WithError(err).Error("failed to parse auth token")
				badRequest(w, "could not parse auth token")
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				log.Printf("Auth token claims: %+v", claims)
				userIDString, ok := claims["id"].(string)
				if !ok {
					log.Error("user ID claims could not be cast to string")
				}

				userID, err := uid.ParseString(userIDString)
				if err != nil {
					log.WithField("userID", claims["id"]).Error("invalid user ID")
					forbidden(w, "auth token contained invalid user ID")
					return
				}

				handler.ServeHTTP(w, r.WithContext(CtxWithSession(r.Context(), Session{UserID: userID})))
				return
			} else {
				log.WithField("token", token).Error("invalid token")
				forbidden(w, "auth token was invalid")
				return
			}
		})
	}
}

func GetToken(log *logrus.Logger) http.HandlerFunc {
	signingSecret := env.GetString("WRKHUB_SIGNING_SECRET", defaultKey)
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uid.ParseString(chi.URLParam(r, "userID"))
		if err != nil {
			log.WithError(err).WithField("userID", userID).Error("invalid user ID")
			badRequest(w, "invalid user ID")
			return
		}

		claims := jwt.MapClaims(map[string]interface{}{
			"id": userID.String(),
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signed, err := token.SignedString([]byte(signingSecret))
		if err != nil {
			log.WithError(err).Error("failed to sign token")
			serverError(w, "failed to sign token")
			return
		}

		cookie := http.Cookie{
			Name:     "wrkhub-session-token",
			Value:    signed,
			Secure:   false,
			HttpOnly: true,
			Path:     "/",
			// SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, &cookie)
		ok(w, []byte(signed))
	}
}
