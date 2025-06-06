# PhishPayback: Phishing Spammer

The idea is to combat phishing with spam, this is a hacking repo (needless to add that "hacking" IS "ethical" by its sole [origin](https://youtu.be/XMm0HsmOTFI?si=mBFLIIfae_TRcQDD). **Please stop confusing hacking with cracking, there's no need to add "ethical" after "hacking"!**).

Inspired in

- GitHub - [v3lip/PhishingSpammer](https://github.com/v3lip/PhishingSpammer/)
- YouTube - [@ScammerPayback](https://www.youtube.com/@ScammerPayback)

_**Dear Cracker, a note for you if you dare to feel "affected" by this repo:**: First of all: You're not a hacker, you're just a cracker. Be a real hacker: some companies offer [bounties for finding bugs and other vulnerabilities](https://www.google.com/search?q=bounties+for+finding+softwar+evulnerabilities) in their systems. I trust you are able to apply to one of them, look: I'm just doing this for fun in my free time and I'm not even an expert in what you surely can do with your knowledge. Wether you get a bounty, or you find a job in which you feel comfortable, once you have free a little of your precious time, please come back here and [give me a hand](TODO.md) with this repo. I've learnt a lot by just developing this repo_

## Types

1. Send random data to phishing backends
2. TODO: Create DDOs attacks for frontends
3. TODO: Create DDOs attacks for backends

## Targets

For safety reasons, no URL will be in this readme, please check source code for details.

And also, for non-safety reasons, Why should I care about any exposed token in this repository? ¯\\\_(ツ)\_/¯

## Usage

Requirements:

- This repo cloned. All commands here assumes that you are within your cloned copy of this repo
- [Golang installed](https://go.dev/doc/install) (For now, perhaps in future i'd upload binaries...)
- (Optional for development) [docker installed](https://docs.docker.com/engine/install/) (To run a local mockserver)

### Init/Install

```sh
go get
```

### Build

```sh
go build
```

### Run

```sh
./backphish
```

If no options (as above), will list the available phishing targets.

Available CLI options:

```
  -l int
        Log Level. Values: 0-Verbose/Debug 1-Normal 2-Quiet/Severe (default 1)
  -m value
        Use mock server in HTTP requests. Must be a valid URL or empty string "" to take default. ("http://localhost:1080")
  -n string
        Attack name
  -p int
        Simultaneous processes/threads (default 1)
  -q int
        Total quantity of attacks. Set 0 for unlimited (default 4)
```

Note: `--help` option returns the above list of options

### (Miscellaneous)

#### Local mock server (for development only)

Runs a mock/dummy HTTP server in your local 1080 port - Based in [mockserver](https://www.mock-server.com/). With docker:

```sh
docker run -it --rm -p 1080:1080 mockserver/mockserver
```

And then [go to the dashboard](http://localhost:1080/mockserver/dashboard)

## Refereces or related links

The following is a recopilation of technologies/providers used in this repo and/or by phishing pages

- [VCC Generator](https://www.vccgenerator.org/) - Generates cards for testing purposes
- [Replit deployments](https://replit.com/deployments)
- [devtools-detect](https://github.com/sindresorhus/devtools-detect) - Trigger actions depending on detection of devtools/inspector usage
- [Chrome DevTools - Network requests: Test your site by blocking network requests](https://developer.chrome.com/docs/devtools/network-request-blocking) (To disable loading of utilities that mess with devtools/debugger)
- [Reverse IP Lookup](https://reverseip.domaintools.com/)
- [Engintron: Nginx on cPanel/WHM server](https://github.com/engintron/engintron)
- [cf-clearance (old?)](https://github.com/vvanglro/cf-clearance) - [cf-clearance (active?)](https://github.com/ModestasBar/cf-clearance) - _"make a cloudflare v2 challenge pass successfully"_
- [HTTrack](https://github.com/xroche/httrack) - website copier
- [Webflow](https://webflow.com/) - Website builder
- [clevertap(.min.js)](https://developer.clevertap.com/docs/web-quickstart-guide) - ???
- [r/dataisbeautiful - Most common 4 digit PIN numbers](https://www.reddit.com/r/dataisbeautiful/comments/1cn7l7r/oc_most_common_4_digit_pin_numbers_from_an/#lightbox)
- [DNS Propagation Checker](https://www.whatsmydns.net/#NS/)

### Abuse reporting

- [Cloudflare](https://abuse.cloudflare.com/phishing)
- [openprovider](https://www.openprovider.com/company/contact-us/report-abuse)
