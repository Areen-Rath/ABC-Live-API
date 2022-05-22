import requests
from bs4 import BeautifulSoup

def mc_fetch():
    data = requests.get("https://www.moneycontrol.com/")
    soup = BeautifulSoup(data.content, "html.parser")

    links = []
    titles = []
    descs = []
    imgs = []

    a_data = soup.find(class_ = "sub-col-left")
    a_tags = a_data.find_all("a")
    last_link = ""
    for a in a_tags:
        if (
            a["href"] != last_link
            and a["href"][28:43] == "/news/business/"
            and a["href"][42:51] != "/markets/"
            and a["href"][42:55] != "/commodities/"
        ):
            last_link = a["href"]
            links.append(a["href"])
            titles.append(a["title"])

    for index, link in enumerate(links):
        article = requests.get(link)
        article_soup = BeautifulSoup(article.content, "html.parser")

        try:
            desc = article_soup.find("h2", attrs = {"class", "article_desc"}).text
            descs.append(desc)

            img_data = article_soup.find(class_ = "article_image")
            img = img_data.find_all("img")
            imgs.append(img[0]["data-src"])
        except:
            try:
                links.remove(link)
                titles.pop(index)
            except:
                pass
    
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