package notifier

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/typositoire/grocy-alerts/utils"
)

type stdoutNotifier struct {
	Logger utils.Logger
}

func newStdoutNotifier() (Notifier, error) {
	logger, err := utils.NewLogger(os.Stdout, "notifier-stdout")
	if err != nil {
		return nil, err
	}

	return stdoutNotifier{
		Logger: logger,
	}, nil
}

func (n stdoutNotifier) SendNotification(m interface{}) error {
	jsonString, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonString))
	return nil
}

func (n stdoutNotifier) BuildNotification(data interface{}, recipient string) (interface{}, error) {
	return data, nil
}
