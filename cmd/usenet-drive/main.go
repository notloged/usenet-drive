package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/javi11/usenet-drive/internal/domain"
	"github.com/javi11/usenet-drive/internal/webdav"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "usenet-drive",
	Short: "A WebDAV server for Usenet",
	Run: func(_ *cobra.Command, _ []string) {
		// Read the config file
		configData, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		// Parse the config file
		var config domain.Config
		err = yaml.Unmarshal(configData, &config)
		if err != nil {
			log.Fatalf("Failed to parse config file: %v", err)
		}

		// Call the handler function with the config
		srv, err := webdav.StartServer(config)
		if err != nil {
			log.Fatalf("Failed to handle config: %v", err)
		}

		log.Printf("Server started at http://localhost:%v", config.ServerPort)

		srv.ListenAndServe()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "path to YAML config file")
	rootCmd.MarkPersistentFlagRequired("config")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
