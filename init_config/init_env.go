package init_config

import "os"

func init() {
	os.Setenv("MYSQL_HOST", "x")
	os.Setenv("MYSQL_PORT", "x")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PWD", "root")
}
