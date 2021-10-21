# grocy-alerts

Grocy alerts was made in an attempt to give more visibility to expiring soon products in Grocy.

## Usage

```
Fetch products from grocy api and can notify users using multiple backend

Currently supported backend are:
- Sendgrid Dynamic Templates
- Stdout
- More to come.

Usage:
  grocy-alerts watch [flags]

Flags:
      --grocy_api-key string        Grocy API Key to gather products due soon. (default "APIKEY")
      --grocy_api-url /api          Grocy url without /api. (default "http://grocy.example.com")
      --grocy_due-soon-max string   How far due to fetch. (default "5")
  -h, --help                        help for watch
      --interval -1                 Interval to check. If -1, will run once. (default -1)
      --notifier-backend string     Notifier backend. (default "stdout")
      --sg_api-key string           Sendgrid api key. (default "APIKEY")
      --sg_from-email string        Dynamic template ID for notification. (default "grocy-alerts@example.com")
      --sg_recipients string        Comma seperated list of recipient for sendgrid backend.
      --sg_template-id string       Dynamic template ID for notification. (default "d-dXXXXXXXXXXXXXXX")
  -v, --version                     version for watch

Global Flags:
      --config string      config file (default is $HOME/.grocy-alerts.yaml)
  -l, --log-level string   Extra output (default "info")
```

### Sendgrid Example

```shell
grocy-alerts watch \
  --grocy_api-url https://grocy.example.com \
  --grocy_api-key APIKEY \
  --grocy_due-soon-max 1 \ ## Only products due in the next day. (Defaults to 5)
  --notifier-backend sendgrid \ ## Send notification through Sendgrid
  --sg_template-id d-XXXXXXXXXX \ ## Sendgrid dynamic template ID
  --sg_from-email grocy-alerts@grocy.example.com \ ## Sendgrid FROM
  --sg_api-key SENDGRID_TOKEN \ ## Sendgrid token
  --sg_recipients "me@example.com,you@example.com" ## Comma separated list of recipients
```

### STDOUT Example

```shell
grocy-alerts watch \
  --grocy_api-url https://grocy.example.com \
  --grocy_api-key APIKEY \
  --grocy_due-soon-max 1 \ ## Only products due in the next day. (Defaults to 5)
  --notifier-backend stdout
```

### STDOUT Example (loop)

```shell
grocy-alerts watch \
  --grocy_api-url https://grocy.example.com \
  --grocy_api-key APIKEY \
  --grocy_due-soon-max 1 \ ## Only products due in the next day. (Defaults to 5)
  --notifier-backend stdout \
  --interval 60
```