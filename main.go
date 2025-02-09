package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type NumberResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
	Error      bool     `json:"error,omitempty"`
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

func isArmstrong(n int) bool {
	sum, temp, digits := 0, n, len(strconv.Itoa(n))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func getFunFact(number int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}

func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Get "number" query parameter
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, `{"error": true, "message": "Missing 'number' parameter"}`, http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, `{"message": "Alphabet,""error": true, }`, http.StatusBadRequest)
		return
	}

	properties := []string{}
	if isArmstrong(number) {
		if number%2 == 0 {
			properties = append(properties, "armstrong", "even")
		} else {
			properties = append(properties, "armstrong", "odd")
		}
	} else {
		if number%2 == 0 {
			properties = append(properties, "even")
		} else {
			properties = append(properties, "odd")
		}
	}

	funFact, err := getFunFact(number)
	if err != nil {
		funFact = "Could not fetch fun fact."
	}

	response := NumberResponse{
		Number:     number,
		IsPrime:    isPrime(number),
		IsPerfect:  isPerfect(number),
		Properties: properties,
		DigitSum:   digitSum(number),
		FunFact:    funFact,
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.MarshalIndent(response, "", "  ")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/api/classify-number", classifyNumberHandler)
	fmt.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
