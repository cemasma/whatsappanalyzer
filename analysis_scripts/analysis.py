
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

# sigmoid fonksiyonu
def sigmoid(x):
    output = 1/(1+np.exp(-x))
    return output

# sigmoid fonksiyonu türevi
def sigmoid_output_to_derivative(output):
    return output*(1-output) # s(x)(1- s(x))
 
def clean_up_sentence(sentence):
    # cümleyi tokenize et
    sentence_words = nltk.word_tokenize(sentence)
    # her kelimeyi stem yap
    sentence_words = [word for word in sentence_words]
    return sentence_words

# bag of words: 0 or 1 => cümlenin çantada bulunan kelimeleri için
def bow(sentence, words, show_details=False):
    # tokenize
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
    # giriş katmanımız hazırladığımız bag of words
    l0 = x
    # matrix çarpımı input*hidden layer
    l1 = sigmoid(np.dot(l0, synapse_0))
    # çıkış katmanı
    l2 = sigmoid(np.dot(l1, synapse_1))
    return l2

# olasılık thresholdu
ERROR_THRESHOLD = 0.9

# synapsleri load et analiz tipine göre
synapse_file = "sentiment_synapses.json"
if analysis_type == "topic":
    synapse_file = "topic_synapses.json"
    ERROR_THRESHOLD = 0.4
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