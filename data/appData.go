package data

import (
	"encoding/json"
	"log"
	"os"

	"github.com/atarte/cli-cd/config"
)

type AppData struct {
	MonitoredAppList []MonitoredApp `json:"app_list"`
}

// AddMonitoredApp
func (a *AppData) AddMonitoredApp(m MonitoredApp) AppData {
	a.MonitoredAppList = append(a.MonitoredAppList, m)

	a.UpdateCliCdConfigFile()

	return *a
}

// DeleteMonitoredApp
func (a *AppData) DeleteMonitoredApp() {
	// uuidToDelete := m.Uuid

	// for i, monitoredApp := range a.MonitoredAppList {
	// 	if monitoredApp.Uuid == uuidToDelete {
	// 		a.MonitoredAppList[i] = a.MonitoredAppList[len(a.MonitoredAppList)-1]
	// 		a.MonitoredAppList[len(a.MonitoredAppList)-1] = MonitoredApp{}
	// 		a.MonitoredAppList = a.MonitoredAppList[:len(a.MonitoredAppList)-1]

	// 		break
	// 	}
	// }

	a.UpdateCliCdConfigFile()
}

// UpdateMonitoredApp
func (a *AppData) UpdateMonitoredApp() {
	// uuidToDelete := m.Uuid

	// for i, monitoredApp := range a.MonitoredAppList {
	// 	if monitoredApp.Uuid == uuidToDelete {
	// 		a.MonitoredAppList[i] = m

	// 		break
	// 	}
	// }

	// config.UpdateCliCdConfigFile(a)
}

// UpdateCliCdConfigFile
func (a *AppData) UpdateCliCdConfigFile() {
	if !config.IsCliCdConfigFileExist() {
		log.Fatal("Config file should exist at this point. Cannot 'UpdateCliCdConfigFile'. Relauch the program.")
	}

	dataJson, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	filePath, err := config.GetCliCdFilePath()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filePath, dataJson, 0644); err != nil {
		log.Fatal(err)
	}
}

// LoadCkiCdConfigFile
func LoadCliCdConfigFile() AppData {
	if !config.IsCliCdConfigFileExist() {
		config.CreateCliCdConfigFile()

		return AppData{}
		// log.Fatal("Config file should exist at this point. Cannot 'LoadCliCdConfigFile'. Relauch the program.")
	}

	filePath, err := config.GetCliCdFilePath()
	if err != nil {
		log.Fatal(err)
	}

	byteFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var appdata AppData
	if err := json.Unmarshal(byteFile, &appdata); err != nil {
		log.Fatal(err)
	}

	return appdata
}
