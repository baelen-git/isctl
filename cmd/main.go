package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	intersight "github.com/cgascoig/intersight-simple-go/intersight"
	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/util"
)

var (
	configFile     string
	jsonPathFilter string
	verbose        bool

	client = &util.IsctlClient{}

	auxCommandsGenerators []commandGenerator
)

type commandGenerator func(*util.IsctlClient) *cobra.Command

const (
	keyIDConfigKey   = "keyID"
	keyFileConfigKey = "keyFile"

	keyOutputConfigKey = "output"

	serverConfigKey   = "server"
	insecureConfigKey = "insecure"

	traceEnvName = "ISCTL_TRACE"
)

func main() {
	if os.Getenv(traceEnvName) != "" {
		log.SetLevel(log.TraceLevel)
	}

	log.Trace("isctl starting")

	cobra.OnInitialize(initConfig)

	log.Trace("Got intersight API client")

	rootCmd := gen.GetCommands(client, resultHandler)
	rootCmd.Use = "isctl"

	log.Trace("Got generated commands")

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.isctl.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")

	rootCmd.PersistentFlags().String(keyIDConfigKey, "", "API Key ID")
	rootCmd.PersistentFlags().String(keyFileConfigKey, "", "API Private Key Filename")

	rootCmd.PersistentFlags().String(serverConfigKey, "intersight.com", "Intersight API Server Address (e.g.\"intersight.com\")")
	rootCmd.PersistentFlags().BoolP(insecureConfigKey, "k", false, "Allow insecure server connections (disable SSL certificate validation)")

	rootCmd.PersistentFlags().StringP(keyOutputConfigKey, "o", "default", `Output format. One of default|yaml|json|table|jsonpath|custom-columns|csv. Examples:
	Get Name attribute from all NTP policies: isctl get ntp policy -o jsonpath="[*].Name"
	Table with just Name and Enabled attributes: isctl get ntp policy -o custom-columns=NAME:.Name,ENABLED:.Enabled
	See [https://isctl.netlify.app/1-basic-queries/#output-customisation]`)
	rootCmd.PersistentFlags().StringVar(&jsonPathFilter, "jsonpath", "", "JSONPath filter to apply to the result (e.g. \"$.Name\")")

	viper.BindPFlag(keyIDConfigKey, rootCmd.PersistentFlags().Lookup(keyIDConfigKey))
	viper.BindPFlag(keyFileConfigKey, rootCmd.PersistentFlags().Lookup(keyFileConfigKey))
	viper.BindPFlag(keyOutputConfigKey, rootCmd.PersistentFlags().Lookup(keyOutputConfigKey))
	viper.BindPFlag(serverConfigKey, rootCmd.PersistentFlags().Lookup(serverConfigKey))
	viper.BindPFlag(insecureConfigKey, rootCmd.PersistentFlags().Lookup(insecureConfigKey))

	configCmd := &cobra.Command{
		Use:               "configure",
		Run:               configure,
		Short:             "Configure the isctl command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}
	rootCmd.AddCommand(configCmd)
	// rootCmd.AddCommand(newCmdApply(client))

	log.Trace("Running auxCommandGenerators")
	for _, cmdGen := range auxCommandsGenerators {
		rootCmd.AddCommand(cmdGen(client))
	}
	log.Trace("Finished auxCommandGenerators")

	rootCmd.PersistentPreRunE = validateFlags

	log.Trace("CLI building complete")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func initConfig() {
	log.Trace("Starting initConfig")
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".isctl")
		configFile = filepath.Join(home, ".isctl.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn("Config file not found, creating empty default.")
			f, err := os.Create(configFile)
			if err != nil {
				log.Fatalf("ERROR while creating empty config file: %v", err)
			}
			f.Close()
		} else {
			// Config file was found but another error was produced
			log.Fatalf("ERROR Invalid config file: %v", err)
		}
	}
	log.Trace("Finished initConfig")
}

func configure(cmd *cobra.Command, args []string) {
	log.Trace("Starting configure")
	scanner := bufio.NewScanner(os.Stdin)

	// configure keyID
	fmt.Printf("%s is currently '%s'\n", keyIDConfigKey, viper.GetString(keyIDConfigKey))
	fmt.Printf("Enter new %s or press Enter to keep existing: ", keyIDConfigKey)
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(keyIDConfigKey, input)
	}

	// configure key file name
	fmt.Printf("key filename is currently '%s'\n", viper.GetString(keyFileConfigKey))
	fmt.Printf("Enter new key file name or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(keyFileConfigKey, input)
	}

	// configure server
	fmt.Printf("Intersight API server is currently '%s'\n", viper.GetString(serverConfigKey))
	fmt.Printf("Enter new Intersight API server or press Enter to keep existing: ")
	scanner.Scan()
	if input := scanner.Text(); input != "" {
		viper.Set(serverConfigKey, input)
	}

	log.Println("Writing config file")
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error occurred writing config file: %v", err)
	}
	log.Trace("Finished configure")
}

func validateFlags(cmd *cobra.Command, args []string) error {
	log.Trace("Starting validateFlags")

	var httpTransport http.RoundTripper = http.DefaultTransport

	// Setup logging
	if verbose && os.Getenv(traceEnvName) == "" {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging level set to debug(verbose)")
	}

	if logLevel := log.GetLevel(); logLevel == log.DebugLevel || logLevel == log.TraceLevel {
		httpTransport = newLoggingTransport()
	}

	var err error

	keyID := viper.GetString(keyIDConfigKey)
	if keyID == "" {
		return fmt.Errorf("%s is not set", keyIDConfigKey)
	}

	keyFile := viper.GetString(keyFileConfigKey)
	if keyFile == "" {
		return fmt.Errorf("%s is not set", keyFileConfigKey)
	}
	// try doing ~ expansion on the keyFile path
	if expandedKeyFile, err := homedir.Expand(keyFile); err == nil {
		keyFile = expandedKeyFile
	}

	httpsInsecure := viper.GetBool(insecureConfigKey)
	if httpsInsecure {
		log.Trace("Disabled server certificate verification")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	keyData, err := os.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("Unable to key file: %v", err)
	}

	client.IntersightClient, err = intersight.NewClient(intersight.Config{
		KeyID:         keyID,
		KeyData:       string(keyData),
		BaseTransport: httpTransport,
	})
	if err != nil {
		return fmt.Errorf("Unable to setup Intersight API client: %v", err)
	}

	// authCtx = context.WithValue(authCtx, openapi.ContextServerVariables, map[string]string{
	// 	"server": viper.GetString(serverConfigKey),
	// })

	// if os.Getenv(traceEnvName) != "" {
	// 	trace := &httptrace.ClientTrace{
	// 		GetConn: func(hostPort string) {
	// 			log.Printf("Get Conn: %v\n", hostPort)
	// 		},
	// 		GotConn: func(connInfo httptrace.GotConnInfo) {
	// 			log.Printf("Got Conn: %+v\n", connInfo)
	// 		},

	// 		PutIdleConn: func(err error) {
	// 			log.Printf("PutIdleConn: err: %v\n", err)
	// 		},

	// 		GotFirstResponseByte: func() {
	// 			log.Printf("GotFirstResponseByte\n")
	// 		},

	// 		Got100Continue: func() {
	// 			log.Printf("Got100Continue\n")
	// 		},

	// 		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
	// 			log.Printf("Got1xxResponse code: %v header %v\n", code, header)
	// 			return nil
	// 		},

	// 		DNSStart: func(dsi httptrace.DNSStartInfo) {
	// 			log.Printf("DNSStart %v\n", dsi)
	// 		},

	// 		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
	// 			log.Printf("DNS Info: %+v\n", dnsInfo)
	// 		},

	// 		ConnectStart: func(network, addr string) {
	// 			log.Printf("ConnectStart network: %v, addr: %v\n", network, addr)
	// 		},

	// 		ConnectDone: func(network, addr string, err error) {
	// 			log.Printf("ConnectDone network: %v addr %v err %v\n", network, addr, err)
	// 		},

	// 		TLSHandshakeStart: func() {
	// 			log.Printf("TLSHandshakeStart\n")
	// 		},

	// 		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
	// 			log.Printf("TLSHandshakeDone connectionstate %v err %v\n", cs, err)
	// 		},

	// 		WroteHeaderField: func(key string, value []string) {
	// 			log.Printf("WroteHeaderField key %v value %v\n", key, value)
	// 		},

	// 		WroteHeaders: func() {
	// 			log.Printf("WroteHeaders\n")
	// 		},

	// 		Wait100Continue: func() {
	// 			log.Printf("Wait100Continue\n")
	// 		},

	// 		WroteRequest: func(wri httptrace.WroteRequestInfo) {
	// 			log.Printf("WroteRequest %v\n", wri)
	// 		},
	// 	}

	// 	authCtx = httptrace.WithClientTrace(authCtx, trace)
	// }

	log.Trace("Finished validateFlags")

	return nil
}
