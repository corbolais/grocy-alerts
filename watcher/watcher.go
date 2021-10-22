package watcher

import (
	"os"

	"github.com/spf13/viper"
	"github.com/typositoire/grocy-alerts/grocy"
	"github.com/typositoire/grocy-alerts/notifier"
	"github.com/typositoire/grocy-alerts/utils"
)

type Watcher interface {
	Run()
}

type watcher struct {
	Logger      utils.Logger
	GrocyClient grocy.Grocy
	Notifier    notifier.Notifier
}

func New(grocy_url string, grocy_apikey string, sg_apikey string, notifier_backend string) (Watcher, error) {
	logger, err := utils.NewLogger(os.Stdout, "watcher")
	if err != nil {
		return nil, err
	}

	notifierOptions := map[string]string{
		"backend":        notifier_backend,
		"sg_from":        viper.GetString("sg_from-email"),
		"sg_template_id": viper.GetString("sg_template-id"),
		"sg_api-key":     viper.GetString("sg_api-key"),
	}

	n, err := notifier.NewNotifier(notifierOptions)
	if err != nil {
		return nil, err
	}

	gClient, err := grocy.NewClient(grocy_url, grocy_apikey)
	if err != nil {
		return nil, err
	}
	return &watcher{
		Logger:      logger,
		GrocyClient: gClient,
		Notifier:    n,
	}, nil
}

func (w watcher) Run() {
	days := viper.GetString("grocy_due-soon-max")
	tos := viper.GetString("sg_recipients")

	dp, err := w.GrocyClient.GetDueProduct(days)
	if err != nil {
		w.Logger.Warnf("Error: %s", err.Error())
	}

	m, err := w.Notifier.BuildNotification(dp, tos)

	if err != nil {
		w.Logger.Fatalf("Error: %s", err.Error())
	}

	err = w.Notifier.SendNotification(m)

	if err != nil {
		w.Logger.Fatalf("Error: %s", err.Error())
	}
}
