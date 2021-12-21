package config

func configExampleLoad() []configParams {
	cfg := []configParams{
		{
			Name:     "server with jump",
			User:     "appuser",
			Host:     "192.168.8.35",
			Port:     "22",
			Password: "123456",
			Keypath:  "",
			Jump: []jump{
				{
					User: "appuser",
					Host: "192.168.8.36",
					Port: "2222",
				},
			},
		},
	}
	return cfg
}

func configExampleLoadJump() []configParams {
	cfg := []configParams{
		{
			Name:     "server with jump",
			User:     "appuser",
			Host:     "192.168.8.35",
			Port:     "22",
			Password: "123456",
			Keypath:  "",
			Jump: []jump{
				{
					User: "appuser",
					Host: "192.168.8.36",
					Port: "2222",
				},
				{
					User: "appuser",
					Host: "192.168.8.36",
					Port: "2222",
				},
			},
		},
	}
	return cfg
}
