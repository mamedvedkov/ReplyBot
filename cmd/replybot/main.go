package main

import (
	"flag"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/mamedvedkov/tools/app"

	"github.com/mamedvedkov/ReplyBot/internal"
	"github.com/mamedvedkov/ReplyBot/internal/bot"
	"github.com/mamedvedkov/ReplyBot/internal/storage"
)

type Config struct {
	Bot bot.Config
	PG  storage.Config
}

func main() {
	printConfig()

	_app := app.NewApp()
	cfg := mustConfig()
	pg := storage.Must(cfg.PG)
	svc := internal.NewService(_app.Logger(), pg, pg, pg)
	_bot := bot.Must(cfg.Bot, _app.Logger(), svc)

	_app.
		AddWorkers(_bot.Start).
		Run()
}

func mustConfig() (cfg Config) {
	err := envconfig.Process("", &cfg)
	if err != nil {
		envconfig.Usage("", &cfg)
		log.Fatalln(err)
	}

	return
}

func printConfig() {
	f := flag.Bool("print-config", false, "if need to write file with env usage")
	if f == nil {
		return
	}
	file, err := os.OpenFile("env.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := Config{}
	envconfig.Usagef("", &cfg, file, defaultMdTableFormat)
	log.Fatalln("config printed")
}

const defaultMdTableFormat = `This application is configured via the environment. The following environment
variables can be used:

|KEY|	TYPE|	DEFAULT|	REQUIRED|	DESCRIPTION|
|---|---|---|---|---|
{{range .}}|{{usage_key .}}|	{{usage_type .}}|	{{usage_default .}}|	{{usage_required .}}|	{{usage_description .}}|
{{end}}`
