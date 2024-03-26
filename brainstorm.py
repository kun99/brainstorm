import json



def brainstorm():
    try:
        with open('words_dictionary.json') as f:
            dictionary = json.load(f)
        lines = sum(1 for line in dictionary)
        
        words = [None] * lines
        for i, word in enumerate(dictionary):
            words[i] = word
        print(words[0])
    except Exception as e:
        print("Couldn't open dictionary json file.")
    
    
brainstorm()