package scraper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Tüm transform zincirini uygular
func ApplyTransform(value interface{}, sel *goquery.Selection, transforms Transform) interface{} {
	strVal := normalizeToString(value)
	for _, t := range transforms {
		strVal = applySingleTransform(strVal, sel, t)
	}
	return strVal
}

// Tek bir transform adımını uygular
func applySingleTransform(val string, sel *goquery.Selection, t string) string {
	switch {
	case t == "trim":
		return strings.TrimSpace(val)

	case t == "date":
		parts := strings.Split(val, "-")
		if len(parts) == 3 {
			return fmt.Sprintf("%s.%s.%s", parts[2], parts[1], parts[0])
		}
		return val

	case strings.HasPrefix(t, "attr(") && strings.HasSuffix(t, ")"):
		attr := t[5 : len(t)-1]
		if attrVal, exists := sel.Attr(attr); exists {
			return attrVal
		}
	}
	return val
}

// Gelen her türlü veriyi stringe dönüştürür (object, array, primitive)
func normalizeToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case bool:
		return strconv.FormatBool(v)
	default:
		// Geri kalan tüm map, slice, struct vs. için JSON.stringify
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return string(b)
	}
}

func NormalizeTransform(i interface{}) Transform {
	switch v := i.(type) {
	case string:
		return Transform{v}
	case []interface{}:
		var result Transform
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case []string:
		return v
	default:
		return nil
	}
}
