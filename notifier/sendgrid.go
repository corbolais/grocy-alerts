package notifier

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/typositoire/grocy-alerts/sendgrid"
	"github.com/typositoire/grocy-alerts/utils"
)

type sendgridNotifier struct {
	SendgridClient sendgrid.Sendgrid
	Logger         utils.Logger
	From           string
	TemplateID     string
}

func newSendgridNotifier(apikey string, from string, templateID string) (Notifier, error) {
	logger, err := utils.NewLogger(os.Stdout, "notifier-sendgrid")
	if err != nil {
		return nil, err
	}

	client, err := sendgrid.NewClient(apikey)
	if err != nil {
		return nil, err
	}
	return sendgridNotifier{
		SendgridClient: client,
		Logger:         logger,
		From:           from,
		TemplateID:     templateID,
	}, nil
}

func (n sendgridNotifier) SendNotification(m interface{}) error {
	return n.SendgridClient.SendEmail(m.(*mail.SGMailV3))
}

func (n sendgridNotifier) BuildNotification(data interface{}, recipient string) (interface{}, error) {
	if !strings.Contains(recipient, "@") {
		return nil, fmt.Errorf("invalid recipient %s", recipient)
	} else if recipient == "" {
		return nil, errors.New("empty recipient list")
	}

	m, err := n.SendgridClient.BuildEmail(n.From, "Grocy Alerts", recipient, n.TemplateID, data)

	if err != nil {
		return nil, err
	}

	return m, nil
}
