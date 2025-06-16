package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ApplyTransform(value interface{}, sel *goquery.Selection, transforms Transform) interface{} {
	strVal := normalizeToString(value)

	for _, t := range transforms {
		strVal = applySingleTransform(strVal, sel, t)
	}

	return strVal
}

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
		return fmt.Sprintf("%v", v)
	}
}

func NormalizeTransform(i interface{}) Transform {
	switch v := i.(type) {
	case string:
		return Transform{v}
	case []interface{}:
		result := make(Transform, 0, len(v))
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
