import requests

def poke(name):
    url = f"https://pokeapi.co/api/v2/pokemon/{name}/"
    response = requests.get(url)
    
    if response.status_code != 200:
        print(f"Error: the request failed with status code {response.status_code}")
        print("Check your spelling, otherwise the Pokemon you are searching might not exist in the dataset.")

    data = response.json()
    print(data['name'])
    print(data['weight'])
    print(data['abilities'])
