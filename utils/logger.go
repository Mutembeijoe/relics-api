package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"log"
	"os"
	"strings"
	"time"
)

func InitLogger(){
	path := os.Getenv("LOG_FOLDER")

	env := strings.ToLower(os.Getenv("APP_ENV"))
	//appName:= os.Getenv("APP_NAME")

	if env == "dev" {
		cwd, err := os.Getwd()
		if err == nil {
			path = fmt.Sprintf("%s/logs/", cwd)
		}
	}

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s old.UTC.%s",path,"%Y-%m-%d"),
		rotatelogs.WithLinkName(path+"current"),
		rotatelogs.WithMaxAge(time.Hour *24*7 ),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Fatalln("Failed to Initiate log file :", err)
	}

	log.SetOutput(writer)
}
