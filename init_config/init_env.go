package init_config

import "os"

func init() {
	os.Setenv("MYSQL_HOST", "120.92.208.211")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PWD", "root")

	os.Setenv("REDIS_CLUSTER", "120.92.208.213:7000,120.92.208.213:7001,120.92.208.213:7002,120.92.208.213:7003,120.92.208.213:7004,120.92.208.213:7005")
}
