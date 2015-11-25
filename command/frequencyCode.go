package command

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/strebul/gogy/component"
	"github.com/strebul/gogy/model"
	"log"
	"time"
)

var FrequencyCode = &cobra.Command{
	Use:   "frequencyCode",
	Short: "Scanner frequency acquirer code",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := component.LoadConfig(fcConfigFile)
		if err != nil {
			return err
		}

		db, err := sql.Open("mysql", "root:root@/gogy?charset=utf8")

		for {
			var startTime time.Time
			var timestamp int
			var reader component.Reader

			db.QueryRow("SELECT `timestamp` FROM `frequencyCode` ORDER BY id DESC LIMIT 1").Scan(&timestamp)
			if timestamp == 0 {
				startTime = time.Now().Add(-time.Duration(24) * time.Hour)
				log.Println("Used default start time")
			} else {
				startTime = time.Unix(int64(timestamp), 0)
			}

			request := model.Request{
				Query:     fmt.Sprintf(`message: "%s"`, config.ReadString("fc.pattern")),
				TimeStart: startTime,
				TimeEnd:   time.Now(),
				Size:      500,
				Order:     "acs",
			}

			decorator := component.Decorator{}
			decorator.DecorateRequest(request)

			client := component.Client{
				Host:     config.ReadString("logstash.host"),
				Login:    config.ReadString("logstash.login"),
				Password: config.ReadString("logstash.password"),
			}
			list := client.FindLogs(request)

			stored := 0
			for _, entity := range list {
				reader = entity.Source

				var exists int

				db.QueryRow("SELECT `id` FROM `frequencyCode` WHERE `unique` = ? LIMIT 1", entity.Id).Scan(&exists)
				if exists != 0 {
					log.Println("Skipped", entity.Id)
					continue
				}

				_, err := db.Exec(
					"INSERT `frequencyCode` SET `unique` = ?, `object` = ?, `timestamp` = ?, `code` = ?, `message` = ?",
					entity.Id,
					entity.Object,
					entity.Time.Unix(),
					reader.ReadString(config.ReadString("fc.code")),
					reader.ReadString(config.ReadString("fc.message")),
				)
				if err != nil {
					panic(err)
				}

				stored++
			}

			if stored == 0 {
				log.Println("Sleep", 60, "seconds")
				time.Sleep(60 * time.Second)
			} else {
				log.Println("Stored", stored, "rows")
			}
		}
	},
}

var fcConfigFile string

func init() {
	FrequencyCode.Flags().StringVarP(&fcConfigFile, "config", "c", component.DefaultConfig, "")
}
