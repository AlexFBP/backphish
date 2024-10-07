# Phishing sites impersonating Nequi (alternative 3)

Will only be listed here DNS A records, sorted by IP:

```
obtentupropulsivo.online. 1800  IN      A       157.173.214.100
```

It is required to use DevTools with Android/iPhone emulation, otherwise it won't allow the button to continue and/or will redirect to another page.

Once DevTools configured to do so, the entrypoint of the scam is `https://LISTED_DOMAIN/?p`. **BUT** this redirects to official page. By checking source codes, the real entry point is `https://LISTED_DOMAIN/bdigital/login.html` or ended in `inicio.php`
