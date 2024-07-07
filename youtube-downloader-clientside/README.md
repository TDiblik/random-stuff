After careful consideration, I've decided to NOT finish this project.

I've wanted to get youtube video metadata with tor socks5 and then create a cloudflare worker proxy which downloads it. Unfortunatelly, this aproach takes shit tone of time and the cloudflare worker proxy is VERY unreliable for larger downloads / byte streams. This is likely caused by some undocumented free tier limitations. I've tried coding the proxy in multiple languages (including Rust), but I've always ended up hitting the same problems. Therefore, my whole plan goes to shit. Also, I cannot download the video directly (using fetch) from the provided url.

This folder could be usefull as an example on how to work with processes, environment variables, fiber, cloudflare workers and RAM caches inside golang.

This project was ended in VERY early stages, so there's little to no documentation.

# Setup

`air`

https://stackoverflow.com/a/61306130/16638833
