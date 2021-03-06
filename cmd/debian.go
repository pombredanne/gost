package cmd

import (
	"github.com/knqyf263/gost/db"
	"github.com/knqyf263/gost/fetcher"
	"github.com/knqyf263/gost/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debianCmd represents the debian command
var debianCmd = &cobra.Command{
	Use:   "debian",
	Short: "Fetch the CVE information from Debian",
	Long:  `Fetch the CVE information from Debian`,
	RunE:  fetchDebian,
}

func init() {
	fetchCmd.AddCommand(debianCmd)
}

func fetchDebian(cmd *cobra.Command, args []string) (err error) {
	log.Info("Fetched all CVEs from Debian")
	cves, err := fetcher.RetrieveDebianCveDetails()

	log.Info("Initialize Database")
	driver, err := db.InitDB(viper.GetString("dbtype"), viper.GetString("dbpath"), viper.GetBool("debug-sql"))
	if err != nil {
		return err
	}

	log.Infof("Insert Debian CVEs into DB (%s)", driver.Name())
	if err := driver.InsertDebian(cves); err != nil {
		log.Errorf("Failed to insert. dbpath: %s. err: %s",
			viper.GetString("dbpath"), err)
		return err
	}

	return nil
}
