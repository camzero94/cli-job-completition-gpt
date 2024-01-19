from bs4 import BeautifulSoup
import sys
import requests
import json

def scrapper(url:str):
    print("Url:",url)
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        content = soup.find('p', class_='mb-5 r3 job-description__content text-break')
        requirements = soup.find('p', class_='m-0 r3 w-100')

        if content and not requirements:
            requirements = ""
            content = content.text
        elif requirements and not content:
            content = ""
            requirements = requirements.text
        else:
            content = content.text
            requirements = requirements.text

        data = {
            'content': content,
            'requirements': requirements,
        }
        json_data = json.dumps(data, ensure_ascii=False)
        print(json_data)
    except Exception as e:
        print(e)

if __name__ =="__main__":
    url = sys.argv[1]
    scrapper(url)

