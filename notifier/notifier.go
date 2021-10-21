package notifier

import (
	"os"

	"github.com/typositoire/grocy-alerts/utils"
)

type Notifier interface {
	BuildNotification(data interface{}) (interface{}, error)
	SendNotification(m interface{}) error
}

func NewNotifier(options map[string]string) (Notifier, error) {
	var (
		be  Notifier
		err error
	)

	logger, err := utils.NewLogger(os.Stdout, "notifier")
	if err != nil {
		return nil, err
	}

	switch options["backend"] {
	case "sendgrid":
		be, err = newSendgridNotifier(options["sg_api-key"], options["sg_from"], options["sg_template_id"])
	case "":
		logger.Warnf("No notifier defined, using stdout.")
		be, err = newStdoutNotifier()
	default:
		logger.Warnf("Using default stdout notifier.")
		be, err = newStdoutNotifier()
	}

	return be, err
}
