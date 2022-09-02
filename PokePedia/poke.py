import requests

def poke(name):
    url = f"https://pokeapi.co/api/v2/pokemon/{name}/"
    response = requests.get(url)

    data = response.json()
    print(f"\n{data['name'].upper()}")
    print("-----------------------------")
    print(f"Weight: {data['weight']}\n")
    print("Abilities:")
    for val in data['abilities']:
        print(f"\t{val['ability']['name']}")
    print("\n")
    
def validate(name):
    url = f"https://pokeapi.co/api/v2/pokemon/{name}/"
    response = requests.get(url)
    
    if response.status_code != 200:
        print("Check your spelling, otherwise the Pokemon you are searching might not exist in the dataset.")
        return False
    return True


