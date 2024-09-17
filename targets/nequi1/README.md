# Ownership of phishing domains

Phishing pages impersonating [Nequi](https://www.nequi.com.co) hosted in / pointed by :

- `aplicaparahoy.com`
- `aplicaya-neq.com`
- `finanzasaturitmo.com`
- `impuestoscol.com`
- `intelcore.online`
- `n3quionline.com`
- `neq.n3quionline.com`
- `neqwtx.com`
- `newactivalo.com`
- `nqipr0pulsor.com`
- `nqprepropulso.com`
- `nqpropulsa.com`
- `nqpropulsando.com`
- `nuevopropulsor.com`
- `onlineparati.com`
- `perfectoparti.com`
- `parati-nqui.com`
- `prepropulsandonq.com`
- `prepropulneq.com`
- `prepropulnq.com`
- `prestamo-nequi.website`
- `propulahorrosneq.com`
- `propulcolombiano.com`
- `propulideas.com`
- `propulsandoneqpro.com`
- `propulsor-pre.website`
- `propulsoraprobados.website`
- `siperpropcolombia.com`
- `solicitadesdeya.com`
- `51.107.8.147`

Some of them sends a few requests to

- `yousitesureonlineverification.com`

## WHOIS for domains

_[WHOIS queries made for some of the domains previously listed](logs/whois%20domain/README.md)_.


## DNS A records

With `dig +noall +answer @$(dig +short $(dig +short ns DOMAIN | head -n 1)) DOMAIN_OR_SUBDOMAIN` for all (sub)domains and sorted by IP:

```log
prestamo-nequi.website. 1800    IN      A       82.112.247.12
impuestoscol.com.       50      IN      A       84.32.84.32
yousitesureonlineverification.com. 8891 IN A    92.205.169.9
aplicaparahoy.com.      14400   IN      A       108.167.149.240
nqpropulsando.com.      14400   IN      A       108.167.149.241
prepropulnq.com.        14400   IN      A       108.167.149.241
intelcore.online.       1800    IN      A       157.173.209.51
siperpropcolombia.com.  1800    IN      A       157.173.209.216
nqprepropulso.com.      1800    IN      A       157.173.209.251
nqipr0pulsor.com.       14400   IN      A       162.241.2.171
propulahorrosneq.com.   14400   IN      A       162.241.2.171
prepropulneq.com.       6631    IN      A       162.241.2.172
onlineparati.com.       5230    IN      A       162.241.60.254
propulideas.com.        14400   IN      A       162.241.60.254
nqpropulsa.com.         6751    IN      A       162.241.60.255
propulsandoneqpro.com.  14400   IN      A       162.241.60.255
solicitadesdeya.com.    14400   IN      A       162.241.60.255
n3quionline.com.        6243    IN      A       162.241.61.78
neq.n3quionline.com.    14400   IN      A       162.241.61.78
perfectoparti.com.      14400   IN      A       162.241.61.79
parati-nqui.com.        3625    IN      A       162.241.61.244
finanzasaturitmo.com.   2200    IN      A       162.241.203.120
neqwtx.com.             14400   IN      A       162.241.203.120
aplicaya-neq.com.       12389   IN      A       162.241.203.121
propulcolombiano.com.   1800    IN      A       194.164.64.16
nuevopropulsor.com.     1800    IN      A       194.164.64.213
```

## WHOIS of IPs

_[WHOIS queries made for some of the IPs previously listed](logs/whois%20ip/README.md)_.
