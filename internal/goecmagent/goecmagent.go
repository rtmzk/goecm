package goecmagent

import (
	"go-ecm/internal/goecmagent/config"
	"go-ecm/internal/goecmagent/options"
	"go-ecm/pkg/app"
	"go-ecm/pkg/log"
)

const commandDesc = `The Go ECM agent`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("Go ECM agent",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
