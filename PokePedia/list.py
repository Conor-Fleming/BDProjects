import requests

def list():
    url = "https://pokeapi.co/api/v2/pokemon?offset=10&limit=10/"
    response = requests.get(url)
    if response.status_code != 200:
        print(f"Error: the request failed with status code {response.status_code}")

    data = response.json()
    for val in data['results']:
        print(val['name'])
