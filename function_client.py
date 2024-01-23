import requests
from requests import Response, RequestException
import json
from openai import OpenAI
from dotenv import load_dotenv
import os
import sys



def get_job_offers(name_job: str, pgt: int, skills: list = None):
    scrapper_url_params = f"http://localhost:3000/getJobs?myJob={name_job}&skills=skill1&pages={pgt}"
    try:
        res: Response = requests.get(scrapper_url_params)
        res.raise_for_status()
        return json.dumps(res.json()[5]['content'])
    except RequestException as e:
        return f"Error: {e}"


functionSpec = {
            "name": "get_job_offers",
            "description": "Get the job offer for the given job and page to scrappe",
            "parameters": {
                "type": "object",
                "properties": {
                    "name_job": {
                        "type": "string",
                        "description": "The name of the job",
                    },
                    "pgt": {
                        "type": "integer",
                        "description": "Number pages to scrap",
                    },
                },
                "required": ["name_job", "pgt"],
            },
        }

if __name__ == '__main__':

    argmts_jobs = sys.argv[1:]

    # Check id the number of arguments is correct
    if len(argmts_jobs) != 2:
        print("Error: Invalid number of arguments")
        sys.exit()
    # Get the arguments
    name_job, pgt = argmts_jobs

    # Open AI client
    load_dotenv("./.env")
    api_key = os.environ.get('API_KEY_TOKEN')
    client = OpenAI(api_key=api_key)
    messages=[
            {"role": "system", "content": "Keep answer short"},
            {"role": "user", "content": f"Can you generate short intro for the job offer {name_job}, please scrappe page {pgt}. Given my name is Camilo and Im a Computer science "}
            ]
    # Call Openai API
    completion = client.chat.completions.create(
        model="gpt-3.5-turbo-0613",
        messages=messages,
        functions=[functionSpec],
    )
    reply_content = completion.choices[0].message
    messages.append(reply_content)

    if reply_content.function_call.name == "get_job_offers":
        args = json.loads(reply_content.function_call.arguments)
        print("GPT asked for job offers")
        res = get_job_offers(**args)
        print(f"GPT received job offers {res}")
        messages.append({"role": "user", "content": "Can you generate short intro as you were Camilo and were applying for this job position.you are Computer Scientist graduated at NCKU"})
        messages.append({"role": "function", "name": "get_job_offers", "content":res})
        print("===============Second request================")
        print(messages)

        completion2 = client.chat.completions.create(
                model="gpt-3.5-turbo-0613",
                messages=messages,
                functions=[functionSpec],
            )
        print(completion2.choices[0].message)



    # print(get_current_jobs("Software Engineer", 1)['content'])



