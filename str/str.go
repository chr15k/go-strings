package str

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Return the remainder of a string after the first occurrence of a given value.
func After(subject, search string) string {
	if len(search) == 0 {
		return subject
	}

	slice := strings.SplitAfterN(subject, search, 2)

	return slice[len(slice)-1]
}

// Return the remainder of a string after the last occurrence of a given value.
func AfterLast(subject, search string) string {
	if len(search) == 0 {
		return subject
	}

	position := strings.LastIndex(subject, search)

	if position == -1 {
		return subject
	}

	return subject[position+len(search):]
}

// Get the portion of a string before the first occurrence of a given value.
func Before(subject, search string) string {
	if len(search) == 0 {
		return subject
	}

	position := strings.Index(subject, search)

	if position == -1 {
		return subject
	}

	return subject[:position]
}

// Get the portion of a string before the last occurrence of a given value.
func BeforeLast(subject, search string) string {
	if len(search) == 0 {
		return subject
	}

	position := strings.LastIndex(subject, search)

	if position == -1 {
		return subject
	}

	return subject[:position]
}

// Get the portion of a string between two given values.
func Between(subject, from, to string) string {

	if len(from) == 0 || len(to) == 0 {
		return subject
	}

	return BeforeLast(After(subject, from), to)
}

// Get the smallest possible portion of a string between two given values.
func BetweenFirst(subject, from, to string) string {
	if len(from) == 0 || len(to) == 0 {
		return subject
	}

	return Before(After(subject, from), to)
}

// Convert a value to camel case.
func Camel(value string) string {
	return Lcfirst(Studly(value))
}

// Determine if a given string contains a given substring.
// @todo test
func Contains(value, substr string) bool {
	return strings.Contains(value, substr)
}

// Determine if a given string does not contain a given substring.
// @todo test
func DoesntContain(value, substr string) bool {
	return !Contains(value, substr)
}

// Determine if the given string, or any of the given slice of strings, ends with a given substring.
func EndsWith(haystack string, needles interface{}) bool {

	switch needles.(type) {
	case []string:
	case string:
		needles = []string{needles.(string)}
	default:
		return false
	}

	for _, needle := range needles.([]string) {
		if len(needle) > 0 && strings.HasSuffix(haystack, needle) {
			return true
		}
	}
	return false
}

// Adds a single instance of the given value to a string if it does not already end with that value:
func Finish(value, cap string) string {
	quoted := regexp.QuoteMeta(cap)

	re := regexp.MustCompile(fmt.Sprintf(`(?:%s)+$`, quoted))

	return fmt.Sprintf("%s%s", re.ReplaceAllString(value, ""), cap)
}

// Determine if a given string matches a given pattern.
func Is(patterns interface{}, value string) bool {

	switch patterns.(type) {
	case []string:
	case string:
		patterns = []string{patterns.(string)}
	default:
		return false
	}

	if len(patterns.([]string)) == 0 && len(value) == 0 {
		return true
	}

	for _, pattern := range patterns.([]string) {
		if pattern == value {
			return true
		}

		matched, _ := regexp.MatchString(wildCardToRegexp(pattern), value)
		if matched {
			return true
		}
	}

	return false
}

// Determine if a given value is valid JSON.
func IsJSON(value string) bool {
	return json.Valid([]byte(value))
}

// Determine if a given value is a valid URL.
func IsUrl(value string) bool {
	url, err := url.ParseRequestURI(value)
	return err == nil && url.Scheme != "" && url.Host != ""
}

// Determine if a given value is a valid UUID.
func IsUUID(value string) bool {
	re := regexp.MustCompile(`^[\da-fA-F]{8}-[\da-fA-F]{4}-[\da-fA-F]{4}-[\da-fA-F]{4}-[\da-fA-F]{12}$`)
	return len(re.FindString(value)) > 0
}

// Determine if a given value is a valid ULID.
func IsULID(value string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{26}$`)
	if re.MatchString(value) {
		n, _ := strconv.Atoi(string(value[0]))
		return n <= 7
	}
	return false
}

// Convert a string to kebab case.
func Kebab(value string) string {
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	value = matchAllCap.ReplaceAllString(value, "${1}-${2}")

	value = strings.ReplaceAll(value, "_", "-")
	value = strings.ReplaceAll(value, " ", "-")

	return Lower(value)
}

// Get the length of a given string.
func Length(value string) int {
	return utf8.RuneCountInString(value)
}

// Make a string's first character lowercase.
func Lcfirst(value string) string {
	return Lower(Substr(value, 0, 1)) + Substr(value, 1, Length(value))
}

// Limit the number of characters in a string.
func Limit(value string, limit int) string {
	if Length(value) <= limit || limit <= 0 {
		return value
	}
	return fmt.Sprintf("%s...", string([]rune(value)[:limit]))
}

// Convert the given string to lower-case.
func Lower(value string) string {
	return strings.ToLower(value)
}

// Masks a portion of a string with a repeated character.
func Mask(value, character string, index, length int) string {

	if len(character) == 0 {
		return value
	}

	segment := Substr(value, index, length)
	segmentLen := Length(segment)

	if segmentLen == 0 {
		return value
	}

	char := Substr(character, 0, 1)
	valueLen := Length(value)
	startIndex := index

	if valueLen+index <= 0 {
		return strings.Repeat(char, valueLen)
	}

	if index < 0 {
		if index < -valueLen {
			startIndex = 0
		} else {
			startIndex = valueLen + index
		}
	}

	start := Substr(value, 0, startIndex)
	end := Substr(value, startIndex+segmentLen, valueLen)

	return fmt.Sprintf("%s%s%s", start, strings.Repeat(char, segmentLen), end)
}

// Get the string matching the given pattern.
func Match(pattern, value string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(value)
}

// Get the string matching the given pattern.
func MatchAll(pattern, value string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(value, -1)
}

// Remove all non-numeric characters from a string.
func Numbers(value string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(value, "")
}

// Pad both sides of a string with another.
func PadBoth(value string, length int, pad string) string {

	short := math.Max(0, float64(length-Length(value)))
	shortLeft := int(math.Floor(short) / 2)
	shortRight := int(math.Round(math.Ceil(short) / 2))

	return fmt.Sprintf(
		"%s%s%s",
		Substr(strings.Repeat(pad, shortLeft), 0, shortLeft),
		value,
		Substr(strings.Repeat(pad, shortRight), 0, shortRight),
	)
}

// Pad the left side of a string with another.
func PadLeft(value string, length int, pad string) string {

	short := int(math.Max(0, float64(length-Length(value))))

	return Substr(strings.Repeat(pad, length), 0, short) + value
}

// Pad the right side of a string with another.
func PadRight(value string, length int, pad string) string {

	short := int(math.Max(0, float64(length-Length(value))))

	return value + Substr(strings.Repeat(pad, length), 0, short)
}

// Generate a random, secure password.
func Password(length int, includeNumbers bool, includeSpecial bool) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if includeNumbers {
		charset += "0123456789"
	}

	if includeSpecial {
		charset += "!@#$%^&*()_+"
	}

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range password {
		index, _ := rand.Int(rand.Reader, charsetLength)
		password[i] = charset[index.Int64()]
	}

	return string(password)
}

// Generate a random alpha-numeric string.
func Random(length int) string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	str := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range str {
		index, _ := rand.Int(rand.Reader, charsetLength)
		str[i] = charset[index.Int64()]
	}

	return string(str)
}

// Generate a URL friendly "slug" from a given string.
func Slug(value string, overrides map[string]string) string {

	replacements := make(map[string]string)
	replacements["@"] = "-at-"
	replacements["&"] = "-and-"

	for k, v := range overrides {
		replacements[k] = v
	}

	for k, v := range replacements {
		value = strings.ReplaceAll(value, k, v)
	}

	re, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = re.ReplaceAllString(value, " ")

	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, " ", "-")

	return Lower(value)
}

// Convert a string to snake case.
func Snake(value string) string {
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	value = matchAllCap.ReplaceAllString(value, "${1}-${2}")

	value = strings.ReplaceAll(value, "-", "_")
	value = strings.ReplaceAll(value, " ", "_")

	return Lower(value)
}

// Determine if a given string starts with a given substring.
func StartsWith(haystack string, needles interface{}) bool {

	switch needles.(type) {
	case []string:
	case string:
		needles = []string{needles.(string)}
	default:
		return false
	}

	for _, needle := range needles.([]string) {
		if len(needle) > 0 && strings.HasPrefix(haystack, needle) {
			return true
		}
	}
	return false
}

// Convert a value to studly case.
func Studly(value string) string {

	value = strings.Replace(value, "-", " ", -1)
	value = strings.Replace(value, "_", " ", -1)

	split := strings.Split(value, " ")

	output := make([]string, 0, len(split))
	for _, val := range split {
		output = append(output, Ucfirst(val))
	}

	return strings.Join(output, "")
}

// Remove all "extra" blank space from the given string.
func Squish(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// Substr returns the portion of string specified by the start and length parameters
// The behaviour of this function is mostly the same as the PHP mb_substr function.
//
// see http://php.net/manual/en/function.mb-substr.php
//
// except that:
// * If start or length is invalid, an empty string will be returned.
func Substr(str string, start, length int) string {

	if length == 0 {
		return ""
	}

	runeStr := []rune(str)
	runeLen := len(runeStr)

	if runeLen == 0 {
		return ""
	}

	if start < 0 {
		start = runeLen + start
	}
	if start < 0 {
		start = 0
	}
	if start > runeLen-1 {
		return ""
	}

	end := runeLen

	if length < 0 {
		end = runeLen + length
	} else if length > 0 {
		end = start + length
	}

	if end < 0 || start >= end {
		return ""
	}
	if end > runeLen {
		end = runeLen
	}

	return string(runeStr[start:end])
}

// Take the first or last {limit} characters of a string.
func Take(value string, limit int) string {
	if limit < 0 {
		return Substr(value, limit, Length(value))
	}
	return Substr(value, 0, limit)
}

// Remove all whitespace from both ends of a string.
func Trim(value string) string {
	return strings.TrimSpace(value)
}

// Remove all whitespace from the beginning of a string.
func TrimLeft(value string) string {
	return strings.TrimLeft(value, " ")
}

// Remove all whitespace from the end of a string.
func TrimRight(value string) string {
	return strings.TrimRight(value, " ")
}

// Make a string's first character uppercase.
func Ucfirst(value string) string {
	return Upper(Substr(value, 0, 1)) + Substr(value, 1, Length(value))
}

// Convert the given string to upper-case.
func Upper(value string) string {
	return strings.ToUpper(value)
}

// Convert wildcard to regex pattern
func wildCardToRegexp(pattern string) string {
	components := strings.Split(pattern, "*")
	if len(components) == 1 {
		return "^" + pattern + "$"
	}
	var result strings.Builder
	for i, literal := range components {
		if i > 0 {
			result.WriteString(".*")
		}
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return "^" + result.String() + "$"
}

// Limit the number of words in a string.
func Words(value string, words int) string {

	if len(strings.TrimSpace(value)) == 0 {
		return value
	}

	substrings := strings.Fields(value)

	if len(substrings) <= words {
		return value
	}

	return fmt.Sprintf("%s...", strings.Join(substrings[:words], " "))
}

// Wrap the string with the given strings.
func Wrap(value, before, after string) string {
	if len(after) == 0 {
		after = before
	}
	return fmt.Sprintf("%s%s%s", before, value, after)
}

// Unwrap the string with the given strings.
func Unwrap(value, before, after string) string {

	if StartsWith(value, before) {
		value = value[Length(before):]
	}

	if len(after) == 0 {
		after = before
	}

	if EndsWith(value, after) {
		value = value[:Length(value)-Length(after)]
	}

	return value
}
