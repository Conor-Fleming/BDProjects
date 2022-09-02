from list import list
from poke import poke, validate

def main():
    print("Welcome to PokePedia, to view a Pokemons stats please input its name.")
    print("If you need some ideas to get started, use the keyword 'list', to view a list of Pokemon")
    print()
    print("use 'quit' to exit PokePedia\n")
    while True:
        command = input()
        if command == "list":
            list()
            continue
        elif command == "quit":
            break
        elif validate(command) == True:
            poke(command)
            continue
        else:
            print("Unknown command - Enter a Pokemon name to view stats, or use 'list' for some sample names")
            print("Use 'quit' to exit.\n")     
    print("Goodbye!")

if __name__ =='__main__':
    main()



