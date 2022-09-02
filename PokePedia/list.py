import requests
import random

def list():
    randlist = random.sample(range(1, 100), 5)
    for val in randlist:
        url = f"https://pokeapi.co/api/v2/pokemon/{val}"
        response = requests.get(url)
        if response.status_code != 200:
            print(f"Error: the request failed with status code {response.status_code}")

        data = response.json()
        print(data['name'])
    print()
    
