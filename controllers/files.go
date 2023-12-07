package controllers

import (
	"io/ioutil"
	"regexp"
)

var Data = &Database{}

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

func GetData() *Database {
	filePath := "config.php"

	configInfo, err := reader(filePath)
	if err != nil {
		panic(err.Error())
	}

	Data.Name = configInfo["DB_DATABASE"]
	Data.Pass = configInfo["DB_PASSWORD"]
	Data.User = configInfo["DB_PASSWORD"]
	Data.Host = configInfo["DB_HOSTNAME"]
	Data.Port = configInfo["DB_PORT"]
	Data.Storage = configInfo["DIR_STORAGE"]

	return Data
}
