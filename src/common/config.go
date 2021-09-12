package common

// ServiceConfig describes service wide config parameters
type ServiceConfig struct {
	UserserviceUrl  string
	UserservicePort string
	Mode            string
}

// IsDevMode returns true if application is in devmode
func (s *ServiceConfig) IsDevMode() bool {
	return s.Mode == "development"
}
