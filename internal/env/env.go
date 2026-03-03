package env

import (
	"os"
	"strconv"
)

func GetEnvString(key string, fallBack string) string {
	res, ok := os.LookupEnv(key)
	if !ok {
		return fallBack
	}
	return res
}


func GetEnvInt(key string, fallBack int) int{
	res, ok := os.LookupEnv(key)
	if !ok{
		return  fallBack
	}
	intVal, err := strconv.Atoi(res)
	if err != nil{
		return  fallBack
	}
	return  intVal
}