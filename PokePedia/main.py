from list import list
from poke import poke, validate
import os

def main():
    welcome()
    while True:
        command = input()
        if command == "list":
            list()
            #continue
        elif command == "quit":
            break
        elif command == "clear":
            os.system('clear')
            welcome()
        elif validate(command) == True:
            poke(command)
            #continue
        else:
            print("\nEnter a Pokemon name to view stats, or use 'list' for some sample names")
            print("Use 'quit' to exit.\n")     
    print("Goodbye!")

def welcome():
    print("Welcome to PokePedia, to view a Pokemon's stats please input its name.")
    print("If you need some ideas to get started, use the command 'list', to view a list of Pokemon.")
    print("Use 'quit' to exit PokePedia.\n")


if __name__ =='__main__':
    main()



