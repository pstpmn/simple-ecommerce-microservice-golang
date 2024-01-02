package migration

type IMigration interface {
	AutoMigrate(user, pass, host, port, dbName string)
}
