package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "serve",
		Usage: "run aaa server",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "port", Value: 8080, EnvVars: []string{"PORT"}, Usage: "Which port to run the server on"},
		},
		Action: Serve,
	}

	if err := app.Run(os.Args); err != nil {
		log.Err(err).Msg("cli error")
	}
}

func Serve(c *cli.Context) error {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		alliterate := strings.ToLower(r.URL.Query().Get("alliterate")) == "true"

		sep := r.URL.Query().Get("sep")

		if sep == "" {
			sep = "-"
		}
		numStr := r.URL.Query().Get("num")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			num = 2
		}
		words := aaa.Generate(num, &aaa.Options{Alliterate: alliterate})

		w.Write([]byte(strings.Join(words, sep)))
	})

	portNum := c.Int("port")
	log.Info().Int("port", portNum).Msg("run server")
	return http.ListenAndServe(fmt.Sprintf(":%d", portNum), m)
}
