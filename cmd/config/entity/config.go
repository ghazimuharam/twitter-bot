package entity

import "time"

// Config store all configuration for the application
type Config struct {
	App    App
	Vendor Vendor
}

// App configuration for main app config
type App struct {
	Consumer                   Consumer
	Access                     Access
	CronList                   CronList
	FeatureFlag                FeatureFlag
	DefaultCountTweetRetriever int
	NumberOfDM                 int
	TriggerWord                string
}

// Consumer configuration for consumer key and secret
type Consumer struct {
	Key    string
	Secret string
}

// Access configuration for access token and secret
type Access struct {
	Token     string
	Secret    string
	AccountID string
}

// CronList configuration for cron timer
type CronList struct {
	ReadMessage string
}

// FeatureFlag configuration for feature toggle
type FeatureFlag struct {
	IsCronWithSeconds bool
}

// Vendor configuration for main app config
type Vendor struct {
	LocalCache LocalCache
}

// LocalCache store local cache configuration
type LocalCache struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}
