package common

import (
	"fmt"
	"math"
	"math/rand" // "crypto/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

// Adds a separator S in middle, according to given format
//
// Formats:
//   - 0: NNNSNNNSNNNN - Example: phone number
//   - 1: NNNNSNNNNSNNNNSNNNN - Example: credit card
//
// Another format will return the string untouched
func AddSeparator(num string, format int, separator string) string {
	switch format {
	case 0:
		return strings.Join([]string{num[:3], num[3:6], num[6:]}, separator)
	case 1:
		return strings.Join([]string{num[:4], num[4:8], num[8:12], num[12:]}, separator)
	}
	return num
}

func GeneraPin(digitos int) string {
	pin := rand.Intn(int(math.Pow10(digitos)))
	return fmt.Sprintf("%0"+strconv.Itoa(digitos)+"d", pin)
}

func GeneraIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Random delay in a range of time Duration. If min > max, values will be considered swapped!
//
// Example:
//
//	RandDelayRange(3*time.Second, 5*time.Second)
//
// is a random delay ranging from 3 to 5 seconds
//
// Previous example is equivalent to:
//
//	RandDelayWindowed(4*time.Second, 2*time.Second)
func RandDelayRange(min, max time.Duration) {
	if min > max {
		min, max = max, min
	}
	time.Sleep(min + time.Duration(rand.Intn(int(max-min))))
}

// Delay within a random window of duration. The 'duration' is the middle/average delay
// of a 'window' of allowed deviations against the specified duration.
// If half of the window (the peak) exceeds the duration, the window will be trimmed
// to twice the duration
//
// Example:
//
//	RandDelayWindowed(10*time.Second, 2*time.Second)
//
// would be equivalent to a delay of 10 seconds, in average.
// The 2 seconds window, will give a range of +/- 1 second
// in which the final delay will be randomly ranging
//
// Previous example, in other words, is equivalent to:
//
//	RandDelayRange(9*time.Second, 11*time.Second)
func RandDelayWindowed(duration, window time.Duration) {
	peak := window / 2
	if peak > duration {
		// Trim peak and fit window
		peak, window = duration, 2*duration
	}
	min := duration - peak
	RandDelayRange(min, min+window)
}

func RandUserName(p *gofakeit.PersonInfo) string {
	u := ""
	switch rand.Intn(4) {
	default:
	case 0:
		u = p.LastName + p.FirstName
	case 1:
		u = p.FirstName + p.LastName
	case 2:
		u = p.FirstName
	case 3:
		u = p.LastName
	}
	return u
}

func RandEmail(names string, year int) string {
	// If names is empty, just return a random email with gofakeit
	if names == "" {
		return gofakeit.Email()
	}

	// Split names by space
	nameParts := strings.Fields(names)

	// If more than one name, randomly decide if take one or all names
	if len(nameParts) > 1 && rand.Intn(2) == 0 {
		nameParts = []string{PickRand(nameParts)}
	}

	// For each remaining name, randomly decide among: first letter, first 2 letters, full name
	for i, name := range nameParts {
		switch rand.Intn(3) {
		case 0:
			nameParts[i] = string(name[0]) // First letter
		case 1:
			if len(name) > 1 {
				nameParts[i] = name[:2] // First 2 letters
			}
		}
	}

	// If more than one name, randomly decide if glue them with dot, underscore or nothing
	separator := ""
	if len(nameParts) > 1 {
		switch rand.Intn(3) {
		case 0:
			separator = "."
		case 1:
			separator = "_"
		}
	}
	username := strings.Join(nameParts, separator)

	// If year is 0, randomly decide if add a number (0-99) at the end, or a year (19 to 60 years ago) or nothing
	if year == 0 {
		switch rand.Intn(3) {
		case 0:
			// Generate a year between 19 and 60 years ago
			year = time.Now().Year() - gofakeit.Number(19, 60)
		case 1:
			// Generate a random two-digit number
			year = gofakeit.Number(0, 99)
		}
	}

	// Randomly decide if add the full digits or just the last 2 digits
	yearStr := ""
	if year != 0 {
		if rand.Intn(2) == 0 {
			yearStr = strconv.Itoa(year % 100) // Last 2 digits
			if year%100 < 10 {
				yearStr = "0" + yearStr // Add leading zero
			}
		} else {
			yearStr = strconv.Itoa(year) // Full year
		}
	}

	// If to be added either the number or the year, randomly decide if join with a dot, underscore or nothing
	if yearStr != "" {
		switch rand.Intn(3) {
		case 0:
			username += "." + yearStr
		case 1:
			username += "_" + yearStr
		default:
			username += yearStr
		}
	}

	// Finally, randomly decide the domain (gmail, hotmail, yahoo, etc)
	domains := []string{"gmail.com", "hotmail.com", "yahoo.com", "outlook.com", "protonmail.com"}
	domain := PickRand(domains)

	return fmt.Sprintf("%s@%s", username, domain)
}

func RandUserAgent() string {

	// Generated with https://iplogger.org/useragents/
	UAs := []string{
		"Mozilla/5.0 (Linux; Android 6.0.1; HTC Onemini 2 dual sim Build/MRA58K) AppleWebKit/603.14 (KHTML, like Gecko) Chrome/54.0.3785.133 Mobile Safari/600.4",
		"Mozilla/5.0 (Linux; U; Linux x86_64) Gecko/20130401 Firefox/49.3",
		"Mozilla/5.0 (Windows; Windows NT 10.3; Win64; x64) AppleWebKit/534.50 (KHTML, like Gecko) Chrome/47.0.3151.219 Safari/602",
		"Mozilla/5.0 (Linux; U; Linux x86_64) AppleWebKit/536.32 (KHTML, like Gecko) Chrome/47.0.3730.229 Safari/602",
		"Mozilla/5.0 (Windows; Windows NT 10.4;; en-US) AppleWebKit/600.2 (KHTML, like Gecko) Chrome/54.0.2383.114 Safari/600.9 Edge/14.56090",
		"Mozilla/5.0 (Windows; U; Windows NT 10.4; WOW64; en-US) Gecko/20100101 Firefox/57.2",
		"Mozilla/5.0 (compatible; MSIE 11.0; Windows NT 6.2;; en-US Trident/7.0)",
		"Mozilla/5.0 (Windows NT 10.1;) AppleWebKit/536.41 (KHTML, like Gecko) Chrome/47.0.3870.188 Safari/602.8 Edge/10.74805",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; WOW64 Trident/6.0)",
		"Mozilla/5.0 (Windows NT 6.0;) AppleWebKit/600.27 (KHTML, like Gecko) Chrome/51.0.2931.167 Safari/603",
	}

	return PickRand(UAs)
}

func RandPass1(min, max int) string {
	return gofakeit.Password(gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Number(min, max))
}

// Picks a random element of a slice
func PickRand[T any](options []T) T {
	return options[rand.Intn(len(options))]
}
