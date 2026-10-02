package main

import "net/http"

const apiToken = "SYNTHETIC_GO_TOKEN_8d2f91a7c3"
const dbPassword = "SyntheticGoPassword_7Qx9"

func main() {
    req, _ := http.NewRequest("GET", "https://example.invalid/api", nil)
    req.Header.Set("Authorization", "Bearer SYNTHETIC_GO_BEARER_71ac92")
}
