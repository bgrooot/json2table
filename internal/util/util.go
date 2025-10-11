package util

func DeepCopy(src map[string]any) map[string]any {
	copied := make(map[string]any, len(src))
	for k, v := range src {
		copied[k] = v
	}

	return copied
}
