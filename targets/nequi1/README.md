# Ownership of phishing domains

Phishing pages impersonating [Nequi](https://www.nequi.com.co) hosted in / pointed by :

- `aplicaparahoy.com` (*1)
- `aplicaya-neq.com` (*1)
- `co.nqicolmbia.com` (*1)
- `colmbianeq.website`
- `credialinstante.com`
- `cuztco.com/NEQUI`
- `dasbord.online` (*1)
- `finanzasaturitmo.com` (*1)
- `impuestoscol.com`
- `impulsatunq.com`
- `impulsornequi.github.io` ([GitHub repo](https://github.com/impulsornequi/impulsornequi.github.io) abuse!)
- `intelcore.online` (*1)
- `n3quionline.com` (*1)
- `neq.n3quionline.com` (*1)
- `neqwtx.com`
- `newactivalo.com` (*1)
- `nq-col.website`
- `nq-colombiaonline.website` (*1)
- `nqi-pr0pls0r.com`
- `nqicolmbia.com/NEQUI`
- `nqipr0pulsor.com` (*1)
- `nqprepropulso.com` (*1)
- `nqpropulsa.com`
- `nqpropulsando.com`
- `nuevopropulsor.com`
- `onlineparati.com`
- `parati-nqui.com` (*1)
- `perfectoparti.com`
- `preadelanto.online` (*1)
- `prepropulsandonq.com` (*1)
- `prepropulneq.com` (*1)
- `prepropulnq.com` (*1)
- `prestamo-nequi.website` (*1)
- `propulahorrosneq.com`
- `propulcolombiano.com` (*1)
- `propulideas.com` (*1)
- `propulsandoneqpro.com` (*1)
- `propulsor-pre.website` (*1)
- `propulsoraprobados.website` (*1)
- `siperpropcolombia.com`
- `solicitadesdeya.com`
- `51.107.8.147`

_Technical detail: some of them sends (or used to send) a few additional requests to `yousitesureonlineverification.com`_

All domains above listed are impersonating Nequi and its product ["Crédito Propulsor"](https://www.nequi.com.co/personas/credito/propulsor) to stole username (phone number) and password (4 digit pin), and even 2FA (6 digit time-based token).

The entrypoint of the scam is `https://LISTED_DOMAIN/NEQUI/3d/propulsor/nequi/n.html` or ended in `neq.php` instead of `n.html`

**Notes (*N)**

1. Seems to be currently down (Otherwise, please open a Issue!)

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
colmbianeq.website.     14400   IN      A       162.241.60.254
onlineparati.com.       5230    IN      A       162.241.60.254
propulideas.com.        14400   IN      A       162.241.60.254
nq-col.website.         14400   IN      A       162.241.60.255
nqpropulsa.com.         6751    IN      A       162.241.60.255
propulsandoneqpro.com.  14400   IN      A       162.241.60.255
solicitadesdeya.com.    14400   IN      A       162.241.60.255
n3quionline.com.        6243    IN      A       162.241.61.78
neq.n3quionline.com.    14400   IN      A       162.241.61.78
perfectoparti.com.      14400   IN      A       162.241.61.79
credialinstante.com.    14400   IN      A       162.241.61.138
impulsatunq.com.        14400   IN      A       162.241.61.138
co.nqicolmbia.com.      14400   IN      A       162.241.61.244
parati-nqui.com.        3625    IN      A       162.241.61.244
finanzasaturitmo.com.   2200    IN      A       162.241.203.120
neqwtx.com.             14400   IN      A       162.241.203.120
aplicaya-neq.com.       12389   IN      A       162.241.203.121
propulcolombiano.com.   1800    IN      A       194.164.64.16
nuevopropulsor.com.     1800    IN      A       194.164.64.213
```

Non relevant DNS A records listed for the following ones:

|           Domain          |       Reason         |
|            ---            |         ---          |
| `cuztco.com`              | Through Cloudflare   |
| `impulsornequi.github.io` | Through GitHub Pages |

## WHOIS of IPs

_[WHOIS queries made for some of the IPs previously listed](logs/whois%20ip/README.md)_.
