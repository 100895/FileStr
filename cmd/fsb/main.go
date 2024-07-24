package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const versionString = "3.0.0"

// Configuración de la aplicación
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
		if port == "" {
			port = "8080"
		}
		fmt.Printf("Starting server on port %s\n", port)
		// Inicia el servidor en el puerto especificado
		http.HandleFunc("/", handleRequest)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
			fmt.Println("Failed to start server:", err)
			os.Exit(1)
		}
	},
}

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Manage sessions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Managing sessions...")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&port, "port", "8080", "Port to run the bot on")
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

// Manejador de solicitudes HTTP
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Aquí procesas las solicitudes HTTP
	fmt.Fprintf(w, "Hello, this is your bot responding on port %s!", port)
}
