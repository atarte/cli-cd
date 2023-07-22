package data

type MonitoredApp struct {
	// Uuid             string `json:"uuid"`
	Name string `json:"name"`
	Path string `json:"path"`
	// AutomaticDeploy  bool   `json:"automatic_deploy"`
	// RefreshFrequence uint   `json:"refresh_frequence"`
	// DockerfilePath   string `json:"dockerfile_path"`
	// Kubernetes       bool   `json:"Kubernetes"`
	// Checksum         string `json:"checksum"`
}
