from list import list
from poke import poke

def main():
    print("Welcome to PokePedia, to view and Pokemons information please input its name.")
    print("If you need some ideas to get started, use the keyword 'list', to view a list of Pokemon")
    print()
    print("use 'quit' to exit PokePedia")
    quit = False
    while quit == False:
        command = input()
        if command == "list":
            list()
        if command == "quit":
            quit = True
        else:
            poke(command)
    print("Goodbye!")

main()



