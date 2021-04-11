package core

func ParseItem(item interface{}) string {
	if item == nil {
		return ""
	}
	return item.(string)
}
