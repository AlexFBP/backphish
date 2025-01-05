# Phishing sites impersonating Nequi (alternative 7)

The entrypoint is in the form `https://MIRROR/ini.html`. Please refer [main.go](main.go) file, variable `mirrors`

All mirrors seems to be just subdomains in MS Azure so it might be useless to get any whois information


## Scam workflow

The workflow seems to be going in the following order of *.html pages:

```
ini
neq <-------+
otp         |
loading     |
dinamica    |
load        |
dinamica2   |
finish -----+
```

_(as seen on mirror `nequi.blob.core.windows.net/propulsores`)_
