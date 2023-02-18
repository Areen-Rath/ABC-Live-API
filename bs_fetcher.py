import time
import requests
from bs4 import BeautifulSoup
import lxml
from concurrent.futures import ThreadPoolExecutor

session = requests.Session()
def bs_fetch():
    data = session.get("https://www.business-standard.com/")
    soup = BeautifulSoup(data.content, "lxml")
    time.sleep(0.01)

    links = []
    titles = []
    descs = []
    imgs = []

    a_data = soup.find(class_ = "coutent-panel bs-new-top-story-image-block")
    a_tags = a_data.find_all("a")
    topics = ["economy-policy", "companies"]
    for a in a_tags:
        try:
            topic = a["href"].split("/")[2]
        except:
            continue

        if (
            a["href"][:17] == "/article/finance/" 
            or a["href"][:17] == "/article/economy-"
            or a["href"][:17] == "/article/companie"
        ) and a["href"][-5:] == ".html" and topic in topics:
            links.append(f'https://www.business-standard.com{a["href"]}')
            titles.append(a.text)

    for index, link in enumerate(links):
        if titles[index] == "\n\n":
            links.remove(link)
            titles.pop(index)
            continue
    
    with ThreadPoolExecutor(max_workers = len(links)) as p:
        future = list(p.submit(scrape_more, link).result() for link in links)

    for i in future:
        descs.append(i[0])
        imgs.append(i[1])

    data = []
    for index, link in enumerate(links):
        if descs[index] and imgs[index]:
            data.append({
                "title": titles[index],
                "desc": descs[index],
                "link": links[index],
                "img": imgs[index]
            })

    return data

def scrape_more(link):
    desc_img = []

    article = session.get(link)
    article_soup = BeautifulSoup(article.content, "lxml")

    desc = article_soup.find("h2", attrs = {"class", "subHeadingClass"})
    if not desc:
        desc_img.append(None)
    else:
        desc_img.append(desc.text)

    img = article_soup.find("img", attrs = {"class", "imgCont"})
    if not img:
        desc_img.append(None)
    else:
        desc_img.append(img["src"])

    return desc_img