# To Do

**NOTE:** _If you found a new scam site and/or mirror, please read [CONTRIBUTE.md](CONTRIBUTE.md)_

All pending improvements have a rating (Priority-Difficulty)

Priorities: (L)ow / (M)edium / (H)igh
Difficulty or Quantity-of-lines-to-be-changed: (E)asy-few / (M)edium-moderate / (H)ard-huge

- [x] Add huge useless payload
- [x] (M-M) Make it auto-restart instead of stopping/killing, or give a better handling, i.e.: killing only after N failed attempts to same URL within same attack round
- [x] Create CLI arg to define number of processes (and/or threads?)
- [ ] (L-E) Create CLI arg to pass alternative mock server (Implies `-m`)
- [ ] (L-E) Pass attempt number to goroutines for logging purpose
- [ ] (M-H) Create CI-CD workflow, in order to create binaries and/or docker images, i.e., by means of GitHub Actions or Jenkins
- [ ] (L-M) Improve `PrintCookies()` to print all cookies being used without need to give a URL, maybe with a new type extending cookie jar
- [x] (L-E) nq1 target: test if `awaitStatusChange()` is really needed (priorize sending of values)
- [ ] (L-H) Improve calculation of `totalShift` in `AttackRunner()`. Possibly by a **refactor** in which it can be set to each target an average workflow time, or parameters of a range that can be randomly picked by means of any of `RandDelay...()`
- [x] (L-L) Use a [Random User Agent](https://iplogger.org/useragents/) (UA) for each attack, initially from a predefined list of, let's say, 10 UAs; to be updated periodically, monthly perhaps.
- [x] (H-H) Generalyze mirror handling, generalyze what it was done in nq1 for all attacks (By default, a single mirror)
- [x] (L-M) When using mock server, send the original target in a header, "Debug-Original-URL"
- [ ] (L-L) [Discard body reply](https://www.google.com/search?q=golang+http+client+%22discard+OR+drop%22+reply+body) when not passed a object to be filled with reply
- [ ] (L-L) [Build and restart on change](https://www.reddit.com/r/golang/comments/6yap3o/how_do_you_rebuildrestart_your_app_on_file_changes/)
- [ ] (M-H) **Refactor** all targets, in such a manner that for each mirror it could be defined if is alive, so with a CLI flag could be set to attack all active targets... Furthermore, dynamically detect which targets are alive, and attack all of them
- [ ] (M-M) **Refactor** or reuse dp1, to spam with Telegram Bots, [all within fair usage](https://rollout.com/integration-guides/telegram-bot-api/api-essentials) - [oficial](https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this)
- [ ] (L-M) **Refactor** all targets to also show the target at begining
- [ ] (M-L) Retry only if error is not "no such host"
- [ ] (M-H) With concurrency, [clear screen](https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go) and **refactor** so that the status/step of each goroutine can be read peridically, to draw all running goroutines and their states like "Number:Status" separated by "\t"
- [ ] (M-M) Restart current goroutine instead of killing, when there's another goroutine not stuck
- [ ] (M-M) Detect overall i/o timeout on target, so if all goroutines are stuck
- [ ] (L-H) When using menu (and no target selected in CLI), create a sub-menu that allows to save custom defaults in a [.ini file](https://ini.unknwon.io/docs/intro/getting_started), so if the .ini is available, any default values there will take precedence. Also, it mus be an option to restore defaults.
