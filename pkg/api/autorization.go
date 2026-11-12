package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

func generateToken(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, ErrResp{Error: "err with req mrthod"})
		return
	}
	var req struct {
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "err with decode req"})
		return
	}
	expPassword := os.Getenv("TODO_PASSWORD")
	if req.Password != expPassword {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "incorrect pass"})
		return
	}
	token := generateToken(expPassword)
	writeJson(w, http.StatusOK, map[string]string{"token": token})
}
