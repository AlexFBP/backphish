# To Do

**NOTE:** _If you found a new scam site and/or mirror, please read [CONTRIBUTE.md](CONTRIBUTE.md)_

All pending improvements have a rating (Priority-Difficulty):
- Priorities: (L)ow / (M)edium / (H)igh
- Difficulty or Quantity-of-lines-to-be-changed: (E)asy-few / (M)edium-moderate / (H)ard-huge

- [ ] **1.** (L-H) Network - Add huge useless payload (when no body required or as additional params in body) - In other words, allow a DDoS with all frontend objects (html, css, js, ...) - maybe with help of a scapper like [colly](https://github.com/gocolly/colly)
- [ ] **2.** (L-E) Create CLI arg to pass alternative mock server (Implies `-m`)
- [ ] **3.** (M-L) Logging - Pass attempt number to goroutines (attackRunner?)
- [ ] **4.** Logging - Create a statistics record. Ideas:
      - 4.1 It shall allow to record last state returned by a goroutine
      - maybe a fixed lenght array with mutex.
      - It shall have a statistics of last statuses
      - On full buffer, should add 1 to the incoming type, and decrease 1 to oldest one.
- [ ] **5.** (M-H) CICD - Create CI-CD workflow, in order to create binaries and/or docker images, i.e., by means of GitHub Actions, GitLab CI or Jenkins
- [ ] **6.** (L-M) Debug/Devel - Improve `PrintCookies()` to print all cookies being used without need to give a URL, maybe with a new type extending cookie jar
- [ ] **7.** (L-H) Behavior - With concurrency, improve the spread time between attacks according to consequent goroutines
      - implement calculus of `totalShift` in `AttackRunner()`
      - Possibly a **refactor** in which each mirror has a way to calculate the average attack time
- [ ] **8.** (L-L) Performance - [Discard body reply](https://www.google.com/search?q=golang+http+client+%22discard+OR+drop%22+reply+body) when not passed an object to be filled with that reply
- [ ] **9.** (L-L) Debug/Devel - [Build and restart on change](https://www.reddit.com/r/golang/comments/6yap3o/how_do_you_rebuildrestart_your_app_on_file_changes/)
- [ ] **10.** (H-M) Performance - Allow to test that a target/mirror is still alive
      - It might be by adding a feature to all mirrors/targets to test HTTP 2xx on each entry point
      - Alternatively, it might be by adding to the runner/"attack" of all mirrors, a parameter number to indicate until which point should it reach, as example:
        - With `-1`, test all steps
        - With `0`, only run the "alive" test
- [ ] **11.** (H-M) Performance - Save/read "alive/down" state of a target/mirror, on a [.ini file](https://ini.unknwon.io/docs/intro/getting_started)
      - Manually, either by cli or menu option
      - Auto (depends on **10.**)
- [ ] **12.** (M-M) Performance - Update status of all targets/mirrors (depends on **10.** and **11.**)
      - This could be through a CLI flag and/or a menu entry
- [ ] **13.** (M-M) Performance - Attack all alive targets/mirrrors (depends on **10.** and **11.**)
      - This could be through a CLI flag
- [ ] **14.** (L-M) Performance - Retry only if error is not "no such host" (depends on **18.**)
- [ ] **15.** (L-H) Appearance - With concurrency (and no limit?), show a statistic/sum of final status of last N attacks
      - 15.1 Update every time an attack/goroutine finishes
      - 15.2 [clear screen](https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go) on every update
      - 15.3 Update screen only after a fixed period of time, e.g., every 2 seconds. (Improvement of 15.1)
      - It could be a row per each status code, like: `Status: NNN - Qty: NN`
      - show a matrix/table of running attacks (maybe with something like an array of -status- chans... maybe with last status code)
- [ ] **16.** (L-H) When using menu (and no target selected in CLI), create a sub-menu that allows to save custom defaults in a [.ini file](https://ini.unknwon.io/docs/intro/getting_started), so if the .ini is available, any default values there will take precedence. Also, it must be an option to restore defaults. (Might reuse **11.**)
- [x] **17.** ~~(H-M) Debug/Devel - Normalize all attacks: Make all of them to work with mirrors, at least 1 per target~~
- [ ] **18.** (M-M) Metrics/Maintenance - Once attack finished, return metrics like: status after each attack, and ellapsed time on success of all queries of the attack
- [ ] **19.** (L-E) Debug/Devel/Behavior - In verbose mode, restrict a single runner. If any value given to "-p" flag, bypass and warn
- [ ] **20.** Once detected Ctrl+C or attacks finished, update [.ini file](https://ini.unknwon.io/docs/intro/getting_started) (depends on **11.** or **12.**):
      - ask if update all targets
      - optionally, also update last average time of last sucessful attacks
- [ ] **21.** (L-M) Feature - Script to generate a json/yaml report of phishing sites (no matter if being scammed-back or not). (depends on **11.**)
      - Include the last status (alive/down) and if it was "backphished" or not
      - The report shall be git ignored and tracked only as a build file, as example, with GitHub actions, to be included in a release
- [ ] **22.** (L-M) Debug - Prepare "-version" option, initially to show the commit under it was built the binary
- [ ] **23.** (L-M) Devel. - Create utility to test URL patterns, as example, `url-NNN` where `NNN` is dynamic
- [ ] **24.** (M-H) Save/read each mirror/target parameters [in a yaml file](https://medium.com/better-programming/parsing-and-creating-yaml-in-go-crash-course-2ec10b7db850) (maybe related with **21.**)
