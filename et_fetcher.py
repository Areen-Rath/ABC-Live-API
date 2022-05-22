import requests
from bs4 import BeautifulSoup

def et_fetch():
    data = requests.get("https://economictimes.indiatimes.com/markets")
    soup = BeautifulSoup(data.content, "html.parser")

    links = []
    titles = []
    descs = []
    imgs = []

    a_data = soup.find(class_ = "btm_border")
    a_tags = a_data.find_all("a")
    for a in a_tags:
        if a["href"][:21] == "/markets/stocks/news/":
            links.append(a["href"])
            titles.append(a.text)

    links = links[:10]
    titles = titles[:10]
    
    for index, link in enumerate(links):
        article = requests.get(f"https://economictimes.indiatimes.com{link}")
        article_soup = BeautifulSoup(article.content, "html.parser")

        try:
            desc = article_soup.find_all("h2", attrs = {"class", "summary"})[0].text
            descs.append(desc)

            img = article_soup.find_all("img")
            imgs.append(img[3]["src"])
        except:
            links.remove(link)
            titles.pop(index)

    if len(links) > 5:
        links = links[:5]
        titles = titles[:5]

    data = []
    for index, link in enumerate(links):
        data.append({
            "title": titles[index],
            "desc": descs[index],
            "link": link,
            "img": imgs[index]
        })
    
    return data