
# use natural language toolkit
import nltk
nltk.download('punkt')
#from nltk.stem.lancaster import LancasterStemmer
#from TurkishStemmer import TurkishStemmer
import os
import json
import numpy as np
import sys
import argparse
#stemmer = LancasterStemmer()
#stemmer = TurkishStemmer()

parser = argparse.ArgumentParser(description='Provides sentiment or topic analysis')
parser.add_argument('--sentence')
parser.add_argument('--type')

args = parser.parse_args()

analysis_type = args.type
input_sentence = args.sentence

analysis_type = 'sentiment'

# compute sigmoid nonlinearity
def sigmoid(x):
    output = 1/(1+np.exp(-x))
    return output

# convert output of sigmoid function to its derivative
def sigmoid_output_to_derivative(output):
    return output*(1-output)
 
def clean_up_sentence(sentence):
    # tokenize the pattern
    sentence_words = nltk.word_tokenize(sentence)
    # stem each word
    sentence_words = [word for word in sentence_words]
    return sentence_words

# return bag of words array: 0 or 1 for each word in the bag that exists in the sentence
def bow(sentence, words, show_details=False):
    # tokenize the pattern
    sentence_words = clean_up_sentence(sentence)
    # bag of words
    # print(words)
    bag = [0]*len(words)  
    for s in sentence_words:
        for i,w in enumerate(words):
            if w == s: 
                bag[i] = 1
                if show_details:
                    print ("found in bag: %s" % w)

    return(np.array(bag))


def think(sentence, show_details=False):
    x = bow(sentence.lower(), words, show_details)
    # print(x)
    # input layer is our bag of words
    l0 = x
    # matrix multiplication of input and hidden layer
    l1 = sigmoid(np.dot(l0, synapse_0))
    # output layer
    l2 = sigmoid(np.dot(l1, synapse_1))
    return l2

# probability threshold
ERROR_THRESHOLD = 0.9

# load our calculated synapse values
synapse_file = "sentiment_synapses.json"
if analysis_type == "topic":
    synapse_file = "topic_synapses.json"
    ERROR_THRESHOLD = 0.5
elif analysis_type == "sentiment":
    synapse_file = "sentiment_synapses.json" 

with open(synapse_file) as data_file: 
    synapse = json.load(data_file) 
    words = synapse['words']
    classes = synapse['classes']
    synapse_0 = np.asarray(synapse['synapse0']) 
    synapse_1 = np.asarray(synapse['synapse1'])

def classify(sentence, show_details=False):
    results = think(sentence, show_details)

    results = [[i,r] for i,r in enumerate(results) if r>ERROR_THRESHOLD ] 
    results.sort(key=lambda x: x[1], reverse=True) 
    return_results =[[classes[r[0]],r[1]] for r in results]
    print ("%s \n classification: %s" % (sentence, return_results))
    return return_results

classify(input_sentence, show_details=False)