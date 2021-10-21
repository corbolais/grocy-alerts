package sendgrid

import (
	"fmt"
	"os"

	sg "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/typositoire/grocy-alerts/grocy"
	"github.com/typositoire/grocy-alerts/utils"
)

type Sendgrid interface {
	BuildEmail(from string, name string, to string, templateID string, data interface{}) (*mail.SGMailV3, error)
	SendEmail(msg *mail.SGMailV3) error
}

type sendgrid struct {
	Logger utils.Logger
	Client *sg.Client
}

func NewClient(apikey string) (Sendgrid, error) {
	logger, err := utils.NewLogger(os.Stdout, "sendgrid")
	if err != nil {
		return nil, err
	}

	sgClient := sg.NewSendClient(apikey)

	return sendgrid{
		Logger: logger,
		Client: sgClient,
	}, nil
}

func (s sendgrid) BuildEmail(from string, name string, to string, templateID string, data interface{}) (*mail.SGMailV3, error) {
	m := mail.NewV3Mail()
	e := mail.NewEmail(name, from)
	m.SetFrom(e)
	m.SetTemplateID(templateID)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Grocy Alerts", to),
	}
	p.AddTos(tos...)
	p.SetDynamicTemplateData("dueProducts", data.(grocy.SimpleProductData).DueProduct)
	p.SetDynamicTemplateData("overdueProducts", data.(grocy.SimpleProductData).OverdueProduct)
	p.SetDynamicTemplateData("expiredProducts", data.(grocy.SimpleProductData).ExpiredProduct)
	p.SetDynamicTemplateData("missingProducts", data.(grocy.SimpleProductData).MissingProduct)

	m.AddPersonalizations(p)

	return m, nil
}

func (s sendgrid) SendEmail(msg *mail.SGMailV3) error {
	resp, err := s.Client.Send(msg)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("something's wrong: %s", resp.Body)
	}

	return err
}
