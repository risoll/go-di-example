package configutil

import (
	"strconv"
	"os"

	"github.com/risoll/tokosidia/constants"
	"github.com/risoll/tokosidia/utils/dbutil"
)

type (
	Config struct {
		DB dbutil.DBConf
		Instance InstanceConf
	}

	InstanceConf struct {
		Port int
	}
)

func Init() (config Config) {

	driverName := os.Getenv(constants.ConfDBDriver)
	dbConn := os.Getenv(constants.ConfDBConn)
	maxOpen, _ := strconv.ParseInt(os.Getenv(constants.ConfDBMaxOpen), 10, 64)
	maxIdle, _ := strconv.ParseInt(os.Getenv(constants.ConfDBMaxIdle), 10, 64)
	dbConf := dbutil.DBConf{
		MaxIdle: int(maxIdle),
		MaxOpen: int(maxOpen),
		Conn: dbConn,
		DriverName: driverName,
	}

	port, _ := strconv.ParseInt(os.Getenv(constants.ConfPort), 10, 64)
	return Config{
		DB: dbConf,
		Instance: InstanceConf{
			Port: int(port),
		},
	}
}