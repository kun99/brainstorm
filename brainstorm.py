import json
import random

def brainstorm():
    try:
        with open('words_dictionary.json') as f:
            dictionary = json.load(f)
        lines = sum(1 for line in dictionary)
        
        words = [None] * lines
        for i, word in enumerate(dictionary):
            words[i] = word
        random_index = random.randint(0, lines-1)
        print(words[random_index])
    except Exception as e:
        print("Couldn't open dictionary json file.")
    
    
brainstorm()