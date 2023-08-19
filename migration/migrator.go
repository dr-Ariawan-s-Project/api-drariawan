package migrator

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
)

// Migrate ...
func Migrate(cfg config.AppConfig, arg string) {

	var migratePath string

	if len(cfg.APP_PATH) > 0 {
		migratePath = fmt.Sprintf("%s/migration/", cfg.APP_PATH)
	} else {
		migratePath = "migration/"
	}

	rawCMD := "echo \"y\" | migrate -database \"mysql://%s:%s@tcp(%s:%d)/%s\" -path %s %s"
	command := fmt.Sprintf(rawCMD,
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOSTNAME,
		cfg.DB_PORT,
		cfg.DB_NAME,
		migratePath,
		arg)

	fmt.Println(command)

	cmd := exec.Command("bash", "-c", command)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, _ := cmd.CombinedOutput()
	log.Printf("%s", out)
}
