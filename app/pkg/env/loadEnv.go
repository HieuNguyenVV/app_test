package env

import (
	"os"
	"time"

	"github.com/subosito/gotenv"
)

type EnvType struct {
	SecretKey        string
	AccessTokenTime  time.Duration
	RefreshTokenTime time.Duration
}

func getEnvValue(key string, valueDefault string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return valueDefault
	}
	return value
}

func getEnvParseDuration(key, valueDefault string) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		v, _ := time.ParseDuration(valueDefault)
		return v
	}
	v, _ := time.ParseDuration(value)
	return v
}

func LoadENV() (*EnvType, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, err
	}
	env := &EnvType{
		SecretKey:        getEnvValue("SECRET_KEY", "abcdsdghnsakjgsjklghsajklghsakghsalgasdgsaky8914o14h1ir@897dkgvhsdiov"),
		AccessTokenTime:  getEnvParseDuration("ACCESS_TOKEN_TIMER", "20h"),
		RefreshTokenTime: getEnvParseDuration("REFRESH_TOKEN_TIMER", "10h"),
	}
	return env, nil
}
