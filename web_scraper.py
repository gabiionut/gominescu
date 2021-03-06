import requests
import re
from bs4 import BeautifulSoup
headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36'}

BASE_URL = 'http://www.romanianvoice.com/poezii/'

URL = BASE_URL + 'poeti/eminescu.php'
page = requests.get(URL, headers=headers)

soup = BeautifulSoup(page.content, 'html.parser')

anchros = soup.find_all("a")
anchros.pop(0)
anchros.pop(0)
for a in anchros:
    href = a.get('href')

    if href.startswith('../'):
        poemUrl = BASE_URL + href.replace('../', '')
        poemPage = requests.get(poemUrl, headers=headers)
        poemSoup = BeautifulSoup(poemPage.content, 'html.parser')
        poemResults = poemSoup.find_all("td", width="100%", align="left", valign="top")
        text = list(filter(None, poemResults[0].text.replace('Mihai Eminescu', '').split('\n\n\n')))
        if len(text) == 2:
            title = text[0]
            content = text[1]
            body = {'title': title, 'content': content}
            r = requests.post('http://localhost:8080/poem', json = body)
            print(r.json())
        else:
            print(text)
