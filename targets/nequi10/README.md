# Phishing sites impersonating Nequi (alternative 10)

The entrypoint is in the form `https://MIRROR/propulsor/index.html`. Please refer [mirrors.go](mirrors.go) file, variable `mirrors`

All mirrors seems to be just subdomains in MS infrastructure so it might be useless to get any whois information


## Scam workflow

The workflow seems to be going in the following order of `propulsor/*.html` pages:

```
index
lalo
credi.php (*)
index <---------------+
prestamo              |
load.php              |
cargandodatos         |
error_dinamica.php    |
load2.php             |
error-otp             |
load3 ----------------+
```

(*): From here, prefix is `recargas` instead of `propulsor`
