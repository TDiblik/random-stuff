###
### iDNES had a campaign where they hid "coins" into random articles. 
### This is a python script to find out which articles contain these "coins" and "collects" them.
### Worked during the whole 2023 campaign.
### 

import time
from bs4 import BeautifulSoup
import requests
from selenium import webdriver
from selenium.webdriver.common.by import By

locs = []
month = input("Which month's articles you want to go through: ") 
root = "https://www.idnes.cz/sitemap.xml"
soup = BeautifulSoup(requests.get(root).text)
sitemap_tags = soup.find_all("sitemap")
for sitemap in sitemap_tags:
    loc = sitemap.findNext("loc").text.replace("?type=sitemap", "")
    locs.append(f"{loc}?type=clanky-2023_{month}")
    
browser = webdriver.Edge() # if you want to use different browser, feel free to change
browser.get("https://www.idnes.cz/")

print("Login to your account. You have 20 seconds.")
time.sleep(20) # Time for you to login to your account

for line in locs:
    try:
        soup = BeautifulSoup(requests.get(line.strip()).text)
        url_tags = soup.find_all("url")
    except:
        print(f"Skipping {line}")

    for url_tag in url_tags:
        try:
            loc = url_tag.findNext("loc").text
            browser.get(loc)
            try:
                browser.find_element(By.ID, "didomi-notice-agree-button").click()
            except:
                print("Unable to agree to cookies :(")
            elem = browser.find_element(By.XPATH, '//a[.//img[@alt="Seber minci pro Megahru"]]')
            time.sleep(0.5)
            elem.click()
            print("Collected a coin :)")
            time.sleep(2)
        except Exception as e:
            print("Does not contain a coin")