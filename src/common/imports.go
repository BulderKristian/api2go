package common

func AddImportIfNotExisting(imports []map[string]string, s string) []map[string]string {
	exists := false
	for _, entry := range imports {
		if entry["importValue"] == s {
			exists = true
			break
		}
	}
	if !exists {
		imports = append(imports, map[string]string{"importValue": "time"})
	}
	return imports
}
