package config

var config = configure()

// MyVariable defines a configurable value received from the env vars as MY_ENV_VARIABLE
// then can be accessed with config.MyVariable
// var MyVariable = getenv("MY_ENV_VARIABLE", "default value")
