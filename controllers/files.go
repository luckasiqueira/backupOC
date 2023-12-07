package controllers

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func reader(filePath string) (map[string]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	patterns := map[string]string{
		"DB_PORT":     `define\('DB_PORT', '([^']+)'\);`,
		"DB_PASSWORD": `define\('DB_PASSWORD', '([^']+)'\);`,
		"DB_DATABASE": `define\('DB_DATABASE', '([^']+)'\);`,
		"DB_USERNAME": `define\('DB_USERNAME', '([^']+)'\);`,
		"DB_HOSTNAME": `define\('DB_HOSTNAME', '([^']+)'\);`,
		"DIR_STORAGE": `define\('DIR_STORAGE', '([^']+)'\);`,
	}

	configInfo := make(map[string]string)

	for key, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(string(content))

		if len(match) >= 2 {
			configInfo[key] = match[1]
		}
	}

	return configInfo, nil
}

func GetData() {
	filePath := "config.php"

	configInfo, err := reader(filePath)
	if err != nil {
		panic(err.Error())
		return
	}

	for key, value := range configInfo {
		fmt.Printf("%s: %s\n", key, value)
	}
}
