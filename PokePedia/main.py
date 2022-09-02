import requests

url = "https://pokeapi.co/api/v2/pokemon/1/"
response = requests.get(url)
data = response.json()
print(data['name'])





