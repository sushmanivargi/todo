package main

import (
	"encoding/json"
	"os"

	"app/route"
	"app/shared/database"
	"app/shared/jsonconfig"
	"app/shared/server"
)

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Connect to database
	database.Connect(config.Database)

	// Start the listener
	server.Run(route.NewRouter(), config.Server)

}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// configuration contains the application settings
type configuration struct {
	Database database.MySQLInfo "json:'Database'"
	Server   server.Server      `json:"Server"`
}

// config the settings variable
var config = &configuration{}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
