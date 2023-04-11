package cmd

import (
	"fmt"
	"os"

	"github.com/SuperJourney/gopen/cmd/bootstrap"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dbFile string

var MigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage database migrations",
	Long:  `Create, insert or update database migrations`,
}

var initDB = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(dbFile)
		if os.IsNotExist(err) {
			f, err := os.Create(dbFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.Close()

			db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
			if err != nil {
				panic("failed to connect database:%v")
			}

			// Create the users table
			db.AutoMigrate(bootstrap.GetTables()...)

			fmt.Println("Database initialized successfully")
		}
	},
}

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

var genModel = &cobra.Command{
	Use:   "gen",
	Short: "Generates code for the models",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gen.NewGenerator(gen.Config{
			OutPath: "./repo/query",
			Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})

		db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
		if err != nil {
			panic("failed to connect database:%v")
		}
		g.UseDB(db)

		g.ApplyBasic(bootstrap.GetTables()...)
		g.ApplyInterface(func(querier Querier) {
			// Define your dynamic SQL methods here
			querier.FilterWithNameAndRole("", "")
		})

		g.Execute()

		return nil
	},
}

var mock = &cobra.Command{
	Use:   "mock",
	Short: "Generates mock data",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gen.NewGenerator(gen.Config{
			OutPath: "./repo/query",
			Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})

		db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
		if err != nil {
			panic("failed to connect database:%v")
		}
		g.UseDB(db)

		g.ApplyBasic(bootstrap.GetTables()...)
		g.ApplyInterface(func(querier Querier) {
			// Define your dynamic SQL methods here
			querier.FilterWithNameAndRole("", "")
		})

		g.Execute()

		return nil
	},
}

func init() {
	initDB.PersistentFlags().StringVar(&dbFile, "dbfile", "/data/db", "Path to SQLite3 database file")
	MigrationCmd.AddCommand(initDB)
	MigrationCmd.AddCommand(genModel)
	MigrationCmd.AddCommand(mock)
}
