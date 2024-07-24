package main

import (
	"EverythingSuckz/fsb/config"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const versionString = "3.0.0"

var port string

var rootCmd = &cobra.Command{
	Use:               "fsb [command]",
	Short:             "Telegram File Stream Bot",
	Long:              "Telegram Bot to generate direct streamable links for telegram media.",
	Example:           "fsb run --port 8080",
	Version:           versionString,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the bot",
	Run: func(cmd *cobra.Command, args []string) {
		// Aquí se configura el servidor HTTP
		http.HandleFunc("/generate-url", func(w http.ResponseWriter, r *http.Request) {
			// Aquí procesas el hash y generas la URL
			hash := r.URL.Query().Get("hash")
			if hash == "" {
				http.Error(w, "Hash not provided", http.StatusBadRequest)
				return
			}

			url := fmt.Sprintf("http://example.com/%s", hash)
			fmt.Fprintf(w, "Generated URL: %s", url)
		})

		// Inicia el servidor HTTP
		fmt.Printf("Starting server on port %s\n", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	runCmd.Flags().StringVar(&port, "port", "8080", "Port to run the server on")
	config.SetFlagsFromConfig(runCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(sessionCmd)
	rootCmd.SetVersionTemplate(fmt.Sprintf(`Telegram File Stream Bot version %s`, versionString))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
