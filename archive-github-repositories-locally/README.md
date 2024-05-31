Just a little script that downloads every single repository you have access to on Github and saves it on your server every day.

## Setup

1. Run following to setup env: `sudo apt install jq libbz2-dev bzip2`
   - Optionally install `dos2unix`, and run the script throught it, if you `scp` the file from windows system.
2. Generate new github api personal access token (set expiration as you wish)
3. Copy the script on your server (`scp`) and make sure it's runable (`chmod`)
4. Add the following cron to your jobs `0 0 * * * /usr/bin/bash -c "cd /path/to/the/folder && GITHUB_API_TOKEN=<INSERT_YOUR_API_PERSONAL_ACCESS_TOKEN> /path/to/the/script.sh >> /path/to/the/log.txt"`
