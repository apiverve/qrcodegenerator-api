// Package qrcodegenerator provides a Go client for the QR Code Generator API.
//
// For more information, visit: https://apiverve.com/marketplace/qrcodegenerator?utm_source=go&utm_medium=readme
package qrcodegenerator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// ValidationRule defines validation constraints for a parameter.
type ValidationRule struct {
	Type      string
	Required  bool
	Min       *float64
	Max       *float64
	MinLength *int
	MaxLength *int
	Format    string
	Enum      []string
}

// ValidationError represents a parameter validation error.
type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return "Validation failed: " + strings.Join(e.Errors, "; ")
}

// Helper functions for pointers
func float64Ptr(v float64) *float64 { return &v }
func intPtr(v int) *int             { return &v }

// Format validation patterns
var formatPatterns = map[string]*regexp.Regexp{
	"email":    regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`),
	"url":      regexp.MustCompile(`^https?://.+`),
	"ip":       regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`),
	"date":     regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
	"hexColor": regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`),
}

// Request contains the parameters for the QR Code Generator API.
//
// Parameters:
//   - value (required): string - The text or data to encode in the QR code
//   - type: string - The type of data being encoded. Advanced types (wifi, vcard) are premium.
//   - format: string - Output format. Vector formats (svg, webp) are premium.
//   - size: number - Size of the QR code in pixels (50-2048)
//   - margin: number - Margin around the QR code in pixels (0-100)
//   - color: string - Foreground color as hex code (e.g., #000000)
//   - backgroundColor: string - Background color as hex code (e.g., #ffffff)
//   - dotStyle: string - Style of QR code dots
//   - cornerSquareStyle: string - Style of corner squares
//   - cornerDotStyle: string - Style of corner dots
//   - gradient: object - Gradient configuration with type (linear, radial) and colorStops array
//   - logo: string - URL of logo image to place in center of QR code [format: url]
//   - logoSize: number - Size of logo relative to QR code (0.1-0.5)
//   - logoMargin: number - Margin around logo in pixels
type Request struct {
	Value string `json:"value"` // Required
	Type string `json:"type,omitempty"` // Optional
	Format string `json:"format,omitempty"` // Optional
	Margin string `json:"margin,omitempty"` // Optional
}

// ToQueryParams converts the request struct to a map of query parameters.
// Only non-zero values are included.
func (r *Request) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if r == nil {
		return params
	}

	v := reflect.ValueOf(*r)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Get the json tag for the field name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		// Handle tags like `json:"name,omitempty"`
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "-" {
			continue
		}

		// Skip zero values
		if field.IsZero() {
			continue
		}

		// Convert to string
		params[jsonName] = fmt.Sprintf("%v", field.Interface())
	}

	return params
}

// Validate checks the request parameters against validation rules.
// Returns a ValidationError if validation fails, nil otherwise.
func (r *Request) Validate() error {
	rules := map[string]ValidationRule{
		"value": {Type: "string", Required: true},
		"type": {Type: "string", Required: false},
		"format": {Type: "string", Required: false},
		"size": {Type: "number", Required: false},
		"margin": {Type: "number", Required: false},
		"color": {Type: "string", Required: false},
		"backgroundColor": {Type: "string", Required: false},
		"dotStyle": {Type: "string", Required: false},
		"cornerSquareStyle": {Type: "string", Required: false},
		"cornerDotStyle": {Type: "string", Required: false},
		"gradient": {Type: "object", Required: false},
		"logo": {Type: "string", Required: false, Format: "url"},
		"logoSize": {Type: "number", Required: false},
		"logoMargin": {Type: "number", Required: false},
	}

	if len(rules) == 0 {
		return nil
	}

	var errors []string
	v := reflect.ValueOf(*r)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		jsonName := strings.Split(jsonTag, ",")[0]

		rule, exists := rules[jsonName]
		if !exists {
			continue
		}

		// Check required
		if rule.Required && field.IsZero() {
			errors = append(errors, fmt.Sprintf("Required parameter [%s] is missing", jsonName))
			continue
		}

		if field.IsZero() {
			continue
		}

		// Type-specific validation
		switch rule.Type {
		case "integer", "number":
			var numVal float64
			switch field.Kind() {
			case reflect.Int, reflect.Int64:
				numVal = float64(field.Int())
			case reflect.Float64:
				numVal = field.Float()
			}
			if rule.Min != nil && numVal < *rule.Min {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at least %v", jsonName, *rule.Min))
			}
			if rule.Max != nil && numVal > *rule.Max {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at most %v", jsonName, *rule.Max))
			}

		case "string":
			strVal := field.String()
			if rule.MinLength != nil && len(strVal) < *rule.MinLength {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at least %d characters", jsonName, *rule.MinLength))
			}
			if rule.MaxLength != nil && len(strVal) > *rule.MaxLength {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at most %d characters", jsonName, *rule.MaxLength))
			}
			if rule.Format != "" {
				if pattern, ok := formatPatterns[rule.Format]; ok {
					if !pattern.MatchString(strVal) {
						errors = append(errors, fmt.Sprintf("Parameter [%s] must be a valid %s", jsonName, rule.Format))
					}
				}
			}
		}

		// Enum validation
		if len(rule.Enum) > 0 {
			strVal := fmt.Sprintf("%v", field.Interface())
			found := false
			for _, enumVal := range rule.Enum {
				if strVal == enumVal {
					found = true
					break
				}
			}
			if !found {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be one of: %s", jsonName, strings.Join(rule.Enum, ", ")))
			}
		}
	}

	if len(errors) > 0 {
		return &ValidationError{Errors: errors}
	}
	return nil
}

// ResponseData contains the data returned by the QR Code Generator API.
type ResponseData struct {
	Id string `json:"id"`
	Format string `json:"format"`
	Type string `json:"type"`
	Correction string `json:"correction"`
	Size int `json:"size"`
	Margin int `json:"margin"`
	Expires int `json:"expires"`
	DownloadURL string `json:"downloadURL"`
}

