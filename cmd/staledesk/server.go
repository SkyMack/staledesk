package main

import (
	"fmt"

	"github.com/SkyMack/clibase"
	"github.com/SkyMack/staledesk/config"
	"github.com/SkyMack/staledesk/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	apiPathBase          = "/api/v2"
	flagNameAuthRequired = "require-auth"
	flagNameListenHost   = "listen-host"
	flagNameListenPort   = "listen-port"
)

type Server struct {
	Config  config.Data
	Options *ServeOptions
	Router  *gin.Engine
}

type ServeOptions struct {
	ListenHost   string
	ListenPort   string
	AuthRequired bool
}

func genServerOptionsFromFlags(flags *pflag.FlagSet) (ServeOptions, error) {
	listenHost, err := flags.GetString(flagNameListenHost)
	if err != nil {
		return ServeOptions{}, err
	}
	listenPort, err := flags.GetString(flagNameListenPort)
	if err != nil {
		return ServeOptions{}, err
	}
	authReq, err := flags.GetBool(flagNameAuthRequired)
	if err != nil {
		return ServeOptions{}, err
	}

	return ServeOptions{
		ListenHost:   listenHost,
		ListenPort:   listenPort,
		AuthRequired: authReq,
	}, nil
}

func NewServer() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	contacts := controller.NewContactsController()

	apiBase := router.Group(apiPathBase)
	{
		contactsGroup := apiBase.Group("contacts")
		{
			// Requests ending in "contacts" or "contacts/"
			contactsGroup.GET("", contacts.GetAll)
			contactsGroup.GET("/", contacts.GetAll)
			contactsGroup.POST("", contacts.Add)
			contactsGroup.POST("/", contacts.Add)

			// Requests ending in "contacts/autocomplete"
			contactsGroup.GET("/autocomplete", contacts.Search)

			// Requests ending in "contacts/export"
			contactsGroup.GET(fmt.Sprintf("/export/%s", controller.ParamNameContactID), contacts.ExportStatus)
			contactsGroup.POST("/export", contacts.ExportStart)

			// Requests ending in "contacts/ID_NUMBER"
			contactsGroup.DELETE(fmt.Sprintf("/:%s", controller.ParamNameContactID), contacts.Delete)
			contactsGroup.GET(fmt.Sprintf("/:%s", controller.ParamNameContactID), contacts.GetByID)
			contactsGroup.PUT(fmt.Sprintf("/:%s", controller.ParamNameContactID), contacts.Update)
		}

		searchGroup := apiBase.Group("search")
		{
			searchGroup.GET("/contacts", contacts.Filter)
		}
	}

	return router
}

func Serve(opts ServeOptions) error {
	server := NewServer()
	serverBindHost := fmt.Sprintf("%s:%s", opts.ListenHost, opts.ListenPort)
	return server.Run(serverBindHost)
}

func addServerFlags(flags *pflag.FlagSet) {
	serveFlags := &pflag.FlagSet{}
	serveFlags.Bool(flagNameAuthRequired, true, "Whether or not valid client authentication credentials must be used")
	serveFlags.String(flagNameListenHost, "localhost", "The hostname/IP interface the server will bind to (usually localhost or 0.0.0.0")
	serveFlags.String(flagNameListenPort, "5000", "The port the server will listen on")

	clibase.SetFlagsFromEnv(flagPrefix, serveFlags)
	flags.AddFlagSet(serveFlags)
}

func addServeCmd(cmd *cobra.Command) {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "starts the REST API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts, err := genServerOptionsFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			return Serve(opts)
		},
	}
	addServerFlags(serveCmd.Flags())

	cmd.AddCommand(serveCmd)
}
