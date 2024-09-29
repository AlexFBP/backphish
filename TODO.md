# TODO

## Load new targets

**If you find a new target, please DO NOT OPEN a issue nor a pull request!**. Please report it (with, if any, mirrors you identified)
[as a Security Vulnerability](https://docs.github.com/en/code-security/security-advisories/guidance-on-reporting-and-writing-information-about-vulnerabilities/privately-reporting-a-security-vulnerability#privately-reporting-a-security-vulnerability) instead. This allows to not only track new phishing pages and their  specific measure, but also if you want to collaborate in source code, you can start a temporary private fork with it.

Once identified a **"phishing/scam/fraud page"** (hereinafter **"target"**), no matter if it came in a SMS message, email, ads in any social network, etc.

1. If in smartphone, please write down the URL of the target and/or copy it, so you can continue in a PC.
   - As example, if you use Whatsapp web, you can "share" it to yourself to your private/personal "me" chat.
   - Or save it in a OneNote with a list of any target you found.
   - If the target appeared in a integrated browser, still the same: "share" it to your chat or another place that you can easily open on a PC.
2. On your PC, please open a private session of your browser. Alternatively, you can use Tor browser if it makes you feel comfortable.
3. Please DO NOT OPEN IT YET, first [report it to Google Safe Browsing](https://safebrowsing.google.com/safebrowsing/report_phish/) (if your browser does not have an option to report a fraud/phising site) - It will take some time for that report to be a big red full screen warning in browsers that integrates it, Firefox or Chrome as example.
4. In your browser (private session or Tor), paste the URL and try to just open it and see if it loads normally.

   Note: Some targets are focused on iPhone/Android/...smartphone users and/or viceversa, focused only to PC users. Due to that, please do not discard if when opening the target, you end forwarded to Google or another site way different to the scam page you saw.

5. Open a new tab and then, open DevTools (If not sure how to do this, it all depends on your browser, so please "google" `how to open devtools in BROWSER` with the `BROWSER` you use instead)
6. In the DevTools, go to "Network" tab. Things to check:
   - If there is not a message like "Perform a request or reload", there is some traffic. Please clear it, typically with some icon at top-left (bin in Firefox, forbidden sign in Chrome)
   - Please mark the checkbox "Disable Cache"
   - Please mark the checkbox "Preserve Log" (On Firefox, you might need first to open "Settings icon" at top right)
7. Paste again the URL and browse the target the most you can, "safely": When asked, use random but coherent names, addresses, PO codes (Suggestion: use nearest Police department contact data :v haha). As long as you browse the target, the Network traffic will be recorded in the Network tab.

   Note: Similarly as on point 4, some targets detect when DevTools is open, inhibits its loading, or in case you had previously open it, the target attempts to redirect you to another page or rollback your history, so please do not discard those targets, is still possible to work with them.

8. Once finished (target struck in "loading" or ended redirecting you to the real page or another page not directly related to the scam), save all that traffic as a HAR file. As example, something like
   - Chrome: Icon "Export HAR..." (a tray with an arrow down)
   - Firefox: Settings --> Save all as HAR

The purpose is to get a HAR file in which it will be recorded all pertinent HTTP requests made to whatever backend they're using, replicating the same phishing schema but with random data, reduce the quantity of requests need to be made

## Source Code Improvements

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
- [ ] (L-L) nq1 target: test if `awaitStatusChange()` is really needed (priorize sending of values)
- [ ] (L-H) Improve calculation of `totalShift` in `AttackRunner()`. Possibly by a refactor in which it can be set to each target an average workflow time, or parameters of a range that can be randomly picked by means of any of `RandDelay...()`
- [x] (L-L) Use a [Random User Agent](https://iplogger.org/useragents/) (UA) for each attack, initially from a predefined list of, let's say, 10 UAs; to be updated periodically, monthly perhaps.
- [x] (H-H) Generalyze mirror handling, generalyze what it was done in nq1 for all attacks (By default, a single mirror)
- [x] (L-M) When using mock server, send the original target in a header, "Debug-Original-URL"
- [ ] (L-L) [Discard body reply](https://www.google.com/search?q=golang+http+client+%22discard+OR+drop%22+reply+body) when not passed a object to be filled with reply
- [ ] (L-L) [Build and restart on change](https://www.reddit.com/r/golang/comments/6yap3o/how_do_you_rebuildrestart_your_app_on_file_changes/)
