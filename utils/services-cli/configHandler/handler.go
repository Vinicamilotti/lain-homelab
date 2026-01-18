package confighandler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/models"
)

const configFilePath = "./config.json"

func readServiceDir() map[string]models.ServiceConfig {
	services := make(map[string]models.ServiceConfig)
	folderPath := os.Getenv("SERVICES_FOLDER")
	dir, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		if !d.IsDir() {
			continue
		}

		service := models.ServiceConfig{
			Path:    fmt.Sprintf("%s/%s", folderPath, d.Name()),
			Tracked: false,
		}
		services[d.Name()] = service
	}

	return services
}

func readConfig() models.ConfigFile {

	var configFile models.ConfigFile
	autoStartServices, err := os.ReadFile(configFilePath)
	if err != nil {
		configFile.AutoStart = make(map[string]models.ServiceConfig)
		return configFile
	}

	err = json.Unmarshal(autoStartServices, &configFile)
	if err != nil {
		configFile.AutoStart = make(map[string]models.ServiceConfig)
		log.Default().Print(err.Error())
	}
	return configFile

}

func CreateServicesConfig() models.ConfigFile {
	servicesFromDir := readServiceDir()
	servicesFromConfig := readConfig()

	for service, config := range servicesFromDir {
		if _, exists := servicesFromConfig.AutoStart[service]; exists {
			continue
		}

		servicesFromConfig.AutoStart[service] = config
	}

	SaveConfig(servicesFromConfig)

	return servicesFromConfig
}

func SaveConfig(config models.ConfigFile) {
	os.Remove(configFilePath)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(configFilePath, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
