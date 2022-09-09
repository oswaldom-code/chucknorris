package api

import (
	"github.com/oswaldom-code/chucknorris/pkg/log"
	"github.com/oswaldom-code/chucknorris/src/adapters/rest/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(serveCmdNew)
}

var serveCmdNew = &cobra.Command{
	Use:   "server",
	Short: "Spin up the webserver that hosts the API",
	Long:  `The web server hosts the API, and manages the authentication middleware`,
	Run: func(cmd *cobra.Command, args []string) {
		NewServer()
	},
}

func NewServer() {
	//setup routes
	r := routes.SetupRouter()
	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.ErrorWithFields("Error starting server: ", log.Fields{
			"error": err,
		})
	}
}
