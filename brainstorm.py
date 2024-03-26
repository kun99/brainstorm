import json



def brainstorm():
    try:
        with open('words_dictionary.json') as f:
            dictionary = json.load(f)
            lines = sum(1 for line in dictionary)
            print(lines)
        
    except Exception as e:
        print("Couldn't open dictionary json file.")
    
    
brainstorm()