package data

// type AppData struct {
// 	MonitoredAppList []MonitoredApp `json:"app_list"`
// }

// // AddMonitoredApp
// func (a *AppData) AddMonitoredApp(m MonitoredApp) {
// 	a.MonitoredAppList = append(a.MonitoredAppList, m)

// 	// saveAppData(a)
// }

// DeleteMonitoredApp
// func (a *AppData) DeleteMonitoredApp(m MonitoredApp) {
// uuidToDelete := m.Uuid

// for i, monitoredApp := range a.MonitoredAppList {
// 	if monitoredApp.Uuid == uuidToDelete {
// 		a.MonitoredAppList[i] = a.MonitoredAppList[len(a.MonitoredAppList)-1]
// 		a.MonitoredAppList[len(a.MonitoredAppList)-1] = MonitoredApp{}
// 		a.MonitoredAppList = a.MonitoredAppList[:len(a.MonitoredAppList)-1]

// 		break
// 	}
// }

// saveAppData(a)
// }

// UpdateMonitoredApp
// func (a *AppData) UpdateMonitoredApp(m MonitoredApp) {
// uuidToDelete := m.Uuid

// for i, monitoredApp := range a.MonitoredAppList {
// 	if monitoredApp.Uuid == uuidToDelete {
// 		a.MonitoredAppList[i] = m

// 		break
// 	}
// }

// saveAppData(a)
// }

// // LoadAppData is use to load a json file from a given path and return a AppData struct
// func LoadAppData() AppData {

// 	if _, err := os.Stat(appdataPath); err != nil {
// 		newAppData := AppData{}

// 		saveAppData(&newAppData)
// 	}

// 	byteFile, _ := ioutil.ReadFile(appdataPath)

// 	var appdata AppData

// 	json.Unmarshal(byteFile, &appdata)

// 	return appdata
// }

// // saveAppData is use to save an AppData struct into a json file
// func saveAppData(appdata *AppData) {

// 	file, _ := json.MarshalIndent(appdata, "", "    ")

// 	_ = ioutil.WriteFile(appdataPath, file, 0644)
// }
