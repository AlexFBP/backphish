# How to contribute

Unless otherwise noted, all contributes to this project might be classified in one of the following categories:

- New phishing sites and/or mirrors, to be reported
- New functionalities

Please read below according each category

## Reporting new mirrors and/or phishing websites

<details>
<summary>(Old phillosophy)</summary>

**_TLDR_**: The most important thing is to report it as a security issue, as described in next paragraph. Other than that is additional information.

**If you find a new scam, please DO NOT OPEN a issue nor a pull request!**. Please report it (with, if any, mirrors you identified)
[as a Security Vulnerability](https://docs.github.com/en/code-security/security-advisories/guidance-on-reporting-and-writing-information-about-vulnerabilities/privately-reporting-a-security-vulnerability#privately-reporting-a-security-vulnerability) instead. This allows to not only track new phishing pages and their  specific measure, but also if you want to collaborate in source code, you can start a temporary private fork with it.

</details>

If you have one or more, suspicious or confirmed **"phishing/scam/fraud" page(s)/URL(s)/link(s)** (hereinafter **"scam(s)"**), no matter if it came in a SMS message, email, ads in any social network, etc:

Please first [go to Issues](https://github.com/AlexFBP/backphish/issues?q=is%3Aissue%20) to see if it might had already been reported.

If not reported, please [create a New Issue](https://github.com/AlexFBP/backphish/issues/new), including one or more URLs and include the most possible context/details you can.

If possible, please include only **scams** impersonating the same entity, and do not mix URLs impersonating different entites. As example:
- If you found 20 **scams**, all of them appearing to be "Bank1", is enough to put all of them in a single Issue.
- If you found 10 **scams** where 3 of them are impersonating "Bank1" and the remaining 7 **scams** are impersonating "Bank2", please create one Issue with all **scams** impersonating "Bank1" and other for "Bank2"

### Additional steps

If you have time and access to a PC, it would also be helpful if you can [get a HAR capture](https://toolbox.googleapps.com/apps/har_analyzer/) and attach it.

Take care of the following precautions first:

1. If you received the link in your smartphone, please write it down and/or copy it, so you can continue in a PC.
   - As example, if you use Whatsapp web, you can "share" it to yourself to your private/personal "me" chat.
   - Or save it in a OneNote with a list of any scams you found.
   - If the scam appeared in a integrated browser, it should be similar: "share" it to your chat or another place that you can easily open in a PC.
2. On your PC, please open a private/incognito session of your preferred web browser. Alternatively, you can use [Tor browser](https://www.torproject.org/) if it makes you feel safer.
3. Please DO NOT OPEN IT YET, first create a report by means of any available way to do it in your browser. If no available, you can [report it to Google Safe Browsing (GSB)](https://safebrowsing.google.com/safebrowsing/report_phish/). It will take some time for that report to be a big red full screen warning in browsers that integrates GSB, as Firefox or Chrome.
4. In your browser, try to just open the URL and see if it loads normally.

   Note: Some scams are focused to a specific device (iPhone/Android/PC...) thus once they detect your device type is not the one they expect, they might forward you to another page or just show a blank page. So, if when opening the scam site on your PC, you end being forwarded to another website or it just don't show anything, please do not discard it immediately.

5. [Do a HAR capture](https://toolbox.googleapps.com/apps/har_analyzer/), while following scam site instructions but using fake data.

The purpose is to get a HAR file in which it will be recorded all pertinent HTTP requests made to whatever backend they're using, that will be used to replicate the same phishing schema but with random data, reduce the quantity of requests need to be made

## New features

Please do not hesitate in opening a pull request, any help will be appreciated. You can find some ideas in the [TODO.md](TODO.md) file.
