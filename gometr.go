package main

type HealthCheck struct {
	status    string
	ServiceID string
}

type GoMetrClient struct {
	url     string
	timeout int
}

func CreateGoMetrClient(url string, timeout int) *GoMetrClient {
	return &GoMetrClient{url: url, timeout: timeout}
}

func CreateHealthCheck(status string, ServiceID string) *HealthCheck {
	return &HealthCheck{status: status, ServiceID: ServiceID}
}

func (gmc *GoMetrClient) Ping() error {
	return nil
}

func (gmc *GoMetrClient) GetID() string {
	return gmc.url
}

func (gmc *GoMetrClient) Health() bool {
	return gmc.getHealth().status == PassStatus
}

func (gmc *GoMetrClient) getHealth() HealthCheck {
	return HealthCheck{}
}
