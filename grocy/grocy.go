package grocy

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/typositoire/grocy-alerts/utils"
)

type Grocy interface {
	GetDueProduct(days string) (SimpleProductData, error)
}

type grocy struct {
	Logger utils.Logger
	Client *resty.Client
}

func NewClient(url string, apikey string) (Grocy, error) {
	logger, err := utils.NewLogger(os.Stdout, "grocy")
	if err != nil {
		return nil, err
	}

	client := resty.New()
	client.SetHostURL(url)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("GROCY-API-KEY", apikey)

	return &grocy{
		Logger: logger,
		Client: client,
	}, nil
}

func (g grocy) GetDueProduct(days string) (SimpleProductData, error) {
	dueProducts := []SimpleProduct{}
	overdueProducts := []SimpleProduct{}
	expiredProducts := []SimpleProduct{}
	missingProducts := []SimpleProduct{}
	resp, err := g.Client.R().Get("/stock/volatile?due_soon_days=" + days)
	if err != nil {
		g.Logger.Warnf("Cannot get /stock/volatile from %s, error: %s", g.Client.HostURL, err.Error())
		return SimpleProductData{}, err
	}

	var body CurrentVolatileStockResponse

	json.Unmarshal(resp.Body(), &body)

	g.Logger.Debugf("Called %v", resp.Request.RawRequest.URL)

	for _, dueProduct := range body.DueProduct {
		dueProducts = append(dueProducts, SimpleProduct{Name: dueProduct.Product.Name, BestBeforeDate: dueProduct.BestBeforeDate})
	}

	for _, overdueProduct := range body.OverdueProduct {
		overdueProducts = append(overdueProducts, SimpleProduct{Name: overdueProduct.Product.Name, BestBeforeDate: overdueProduct.BestBeforeDate})
	}

	for _, expiredProduct := range body.ExpiredProduct {
		expiredProducts = append(expiredProducts, SimpleProduct{Name: expiredProduct.Product.Name, BestBeforeDate: expiredProduct.BestBeforeDate})
	}

	for _, missingProduct := range body.MissingProduct {
		missingProducts = append(missingProducts, SimpleProduct{Name: missingProduct.Name, AmountMissing: missingProduct.AmountMissing})
	}

	return SimpleProductData{
		DueProduct:     dueProducts,
		OverdueProduct: overdueProducts,
		ExpiredProduct: expiredProducts,
		MissingProduct: missingProducts,
	}, nil
}
