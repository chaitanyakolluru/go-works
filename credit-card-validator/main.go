package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Valid bool `json:"valid"`
}

// determines if a number is a valid Luhn number
func isValidLuhnNumber(digits []int) response {
	sum := 0
	parity := len(digits) % 2
	for i, digit := range digits {
		if i%2 != parity {
			sum += digit
		} else if digit > 4 {
			sum += digit*2 - 9
		} else {
			sum += digit * 2
		}
	}

	return response{
		Valid: digits[len(digits)-1] == (10 - (sum % 10)),
	}
}

func convertNumberToDigits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append([]int{number % 10}, digits...)
		number = number / 10
	}
	return digits
}

func main() {

	payload := map[string]int{}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		if err := c.Bind(&payload); err != nil {
			return err
		}

		if payload["digits"] == 0 {
			return c.JSON(http.StatusBadRequest, "No digits provided")
		}

		return c.JSON(
			http.StatusOK,
			isValidLuhnNumber(
				convertNumberToDigits(payload["digits"]),
			),
		)
	})

	e.Logger.Fatal(e.Start(":1323"))

}
