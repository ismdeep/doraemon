package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ismdeep/doraemon/pkg"
)

// DBWait wait
func DBWait() *cobra.Command {
	var dialect string
	var dsn string
	var timeout int

	cmd := &cobra.Command{
		Use:          "wait",
		Short:        "wait",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch dialect {
			case "mysql":
				return pkg.ConnectToMySQL(dsn, timeout)
			default:
				return fmt.Errorf("unsupported dialect: %v", dialect)
			}
		},
	}

	cmd.Flags().StringVar(&dialect, "dialect", "", "db dialect. e.g. mysql")
	_ = cmd.MarkFlagRequired("dialect")

	cmd.Flags().StringVar(&dsn, "dsn", "", "db dsn.")
	_ = cmd.MarkFlagRequired("dsn")

	cmd.Flags().IntVar(&timeout, "timeout", 60, "timeout (s)")

	return cmd
}

// DBCreate create
func DBCreate() *cobra.Command {
	v := viper.New()

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create",
		Run: func(cmd *cobra.Command, args []string) {
			dialect := v.GetString("dialect")
			dsn := v.GetString("dsn")
			dbName := v.GetString("db")
			auths := v.GetStringSlice("addition-auth")
			switch dialect {
			case "mysql":
				if err := pkg.CreateDBOnMySQL(dsn, dbName, auths); err != nil {
					fmt.Println("Error:", err.Error())
					return
				}
			default:
				fmt.Println("Error:", fmt.Sprintf("%v is not supported", dialect))
			}

			fmt.Println("OK")
		},
	}

	cmd.Flags().String("dialect", "", "db dialect. e.g. mysql")
	_ = cmd.MarkFlagRequired("dialect")
	cmd.Flags().String("dsn", "", "db dsn")
	_ = cmd.MarkFlagRequired("dsn")
	cmd.Flags().String("db", "", "db name")
	_ = cmd.MarkFlagRequired("db")
	cmd.Flags().StringSlice("addition-auth", []string{}, "addition authentications, format: <username>:<password>. e.g. username1:password1")
	_ = v.BindPFlags(cmd.Flags())

	return cmd
}

// DB cli
func DB() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "db",
	}

	cmd.AddCommand(DBWait())
	cmd.AddCommand(DBCreate())

	return cmd
}
