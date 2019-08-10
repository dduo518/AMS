package constant

import "AMS/config"

func GetTokenSecret() string  {
	return config.Conf.TokenSecret
}


