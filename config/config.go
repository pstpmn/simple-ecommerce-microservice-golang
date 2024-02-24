package config

type (
	Config struct {
		App        App        `mapstructure:"App"`
		MongoDb    MongoDb    `mapstructure:"mongoDb"`
		PostgresDb PostgresDb `mapstructure:"PostgresDb"`
		Redis      Redis      `mapstructure:"Redis"`
		Grpc       Grpc       `mapstructure:"Grpc"`
	}

	App struct {
		Name     string `mapstructure:"name"`
		HttpProt string `mapstructure:"httpProt"`
		GrpcProt string `mapstructure:"grpcProt"`
	}
	MongoDb struct {
		Uri string `mapstructure:"uri"`
	}

	PostgresDb struct {
		Uri string `mapstructure:"Uri"`
	}

	Grpc struct {
		Product  string `mapstructure:"product"`
		Customer string `mapstructure:"customer"`
		Auth     string `mapstructure:"auth"`
	}

	Redis struct {
		Uri string `mapstructure:"Uri"`
	}
)
