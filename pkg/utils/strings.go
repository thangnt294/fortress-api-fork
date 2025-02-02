package utils

import (
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
)

// IsNumber checks if a string is a number
func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func FormatNumber(n int64) string {
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func RemoveEmptyString(in []string) []string {
	out := make([]string, 0)
	for _, status := range in {
		if RemoveAllSpace(status) != "" {
			out = append(out, status)
		}
	}

	return out
}

func RemoveAllSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func HasDomain(str string) bool {
	u, err := url.Parse(str)
	if u.Scheme == "mailto" {
		return true
	}
	return err == nil && u.Scheme != "" && u.Host != "" && regexp.MustCompile(`^[^.]+\.[^.]+(\.[^.]+)*$`).MatchString(u.Host)
}

func FormatMoney(value float64, currencyCode string) string {
	var result string

	currency := money.New(1, currencyCode)
	tmpValue := value * math.Pow(10, float64(currency.Currency().Fraction))

	formatted := currency.Multiply(int64(tmpValue)).Display()
	parts := strings.Split(formatted, ".00")
	if len(parts) > 0 {
		result = parts[0]
	}

	return result
}
