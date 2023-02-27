package config

import (
	"fmt"
	"os"

	"github.com/SkyMack/staledesk/internal/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	fileName      = "conf"
	pathConfigDir = "config"
)

var (
	Config = &Data{}

	ErrCannotPopulateContactsFromConfig = fmt.Errorf("cannot populate contact record from config file")
	ErrCannotProcessConfig              = fmt.Errorf("unable to process config file")

	pathConfigFile1 = fmt.Sprintf("..%c%s%c", os.PathSeparator, pathConfigDir, os.PathSeparator)
	pathConfigFile2 = fmt.Sprintf("%s%c", pathConfigDir, os.PathSeparator)
)

type Data struct {
	Contacts map[int]models.Contact
	Raw      *viper.Viper
}

func SetConfig(conf *Data) {
	Config = conf
}

func GenerateConfigData() (*Data, error) {
	var err error
	confData := &Data{
		Contacts: map[int]models.Contact{},
	}
	if err = confData.processConfigFile(); err != nil {
		return &Data{}, err
	}
	if err = confData.populateContacts(); err != nil {
		return &Data{}, err
	}
	return confData, nil
}

func (cd *Data) processConfigFile() error {
	cd.Raw = viper.New()
	cd.Raw.SetConfigType("json")
	cd.Raw.SetConfigName(fileName)
	cd.Raw.AddConfigPath(pathConfigFile1)
	cd.Raw.AddConfigPath(pathConfigFile2)

	err := cd.Raw.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"file.name":  fileName,
			"file.path1": pathConfigFile1,
			"file.path2": pathConfigFile2,
			"error":      err.Error(),
		}).Fatal("unable to read viperConfig file")
		return ErrCannotProcessConfig
	}
	return nil
}

func (cd *Data) populateContacts() error {
	// Populate existing contacts on start-up
	var defaultContacts []models.Contact
	if err := cd.Raw.UnmarshalKey("data.contacts", &defaultContacts); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("cannot read contacts from config file")
		return ErrCannotPopulateContactsFromConfig
	}

	for _, contact := range defaultContacts {
		log.WithFields(log.Fields{
			"contact.active":     contact.Active,
			"contact.address":    contact.Address,
			"contact.created_at": contact.CreatedAt,
			"contact.email":      contact.Email,
			"contact.id":         contact.ID,
			"contact.language":   contact.Language,
			"contact.mobile":     contact.Mobile,
			"contact.name":       contact.Name,
			"contact.people_id":  contact.PeopleID,
			"contact.phone":      contact.Phone,
			"contact.timezone":   contact.TimeZone,
			"contact.twitter_id": contact.TwitterID,
		}).Trace("populating default contact")
		cd.Contacts[contact.ID] = contact
	}
	return nil
}

func addConfigShowRawData(cmd *cobra.Command) {
	showRaw := &cobra.Command{
		Use:   "show-raw",
		Short: "Prints out the raw config data value as read in by viper",
		RunE: func(cmd *cobra.Command, args []string) error {
			//fmt.Println(Config.Raw.Get("data"))
			return nil
		},
	}

	cmd.AddCommand(showRaw)
}

func AddConfigCmd(cmd *cobra.Command) {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Commands related to the config file",
	}

	addConfigShowRawData(configCmd)
	cmd.AddCommand(configCmd)
}
