package scraper

func NormalizeTransform(i interface{}) Transform {
	switch v := i.(type) {
	case string: //transform: "trim"
		return Transform{v}
	case []interface{}:
		result := make(Transform, 0, len(v))
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case []string: //transform: ["trim", "attr(src)"]
		return v
	default:
		return nil
	}
}
