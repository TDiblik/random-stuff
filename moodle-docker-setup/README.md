# How to setup moodle using docker

- Clone this repository
- Go into it from your terminal and run
  ```
  MOODLE_VERSION=3.9.4 #change this to the desired version
  rm -rf moodle*
  git clone https://github.com/moodle/moodle.git
  cd moodle
  git checkout v$MOODLE_VERSION
  cd ..
  ```
- Go into docker-compose.yml and change the [php version based on the version of moodle you selected](https://docs.moodle.org/403/en/PHP#PHP_Versions).
  ```
  eg. from this:
  image: moodlehq/moodle-php-apache:7.4-buster
  to your version:
  image: moodlehq/moodle-php-apache:XXX
  ```
- Run `docker compose up -d`
- You're all done Now, setup your moodle instance by going to `localhost:8080`

## Some usefull info for the installation phase:

- The app will be running on port `8080`
- The app uses `Microsoft SQL Server`
- Your Moodle data will be stored inside a folder named `moodle-data`
- The SQL server hostname is `db`
- Database name is `moodle`
- Database port is `1433`
- The SQL user's name is `SA`
- The SQL user's password is `yourStrong(!)Password`
- Don't panic when the page loads for a long time. It's doing stuff, just wait it out :D
