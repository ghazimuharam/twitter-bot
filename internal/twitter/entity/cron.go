package entity

type CronList struct {
	CronFunction []CronFunction
}

type CronFunction struct {
	Function func()
	Crontab  string
}
