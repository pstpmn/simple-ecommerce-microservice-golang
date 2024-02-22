package config

type (
	Config struct {
		App        App        `mapstructure:"App"`
		MongoDb    MongoDb    `mapstructure:"MongoDb"`
		PostgresDb PostgresDb `mapstructure:"PostgresDb"`
		Redis      Redis      `mapstructure:"Redis"`
	}

	App struct {
		Name     string `mapstructure:"Name"`
		HttpProt string `mapstructure:"httpProt"`
		GrpcProt string `mapstructure:"grpcProt"`
	}
	MongoDb struct {
		Uri string `mapstructure:"Uri"`
	}

	PostgresDb struct {
		Uri string `mapstructure:"Uri"`
	}

	Redis struct {
		Uri string `mapstructure:"Uri"`
	}
)
