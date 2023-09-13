# Nogosari - NLP for Bahasa
A NLP package for bahasa, based on go-sastrawi (https://github.com/RadhiFadlillah/go-sastrawi),
with several modifications to make it procedural, and additional features for advance functionality.

## Basic Concept
There are two things we need to understand:
- Dictionary as a list of indexed words for reference.
- A word can have more than one function (can be called part/position as well)
depending on the structure and the context of a sentence or phrase. Therefore
it is good to store the information in a `uint` variable that is based on
binary encoding that marks the status of word-function (true/false).

Given the condition, we can make a `dictionary` by using map, with the word
(`string`) as the key and word-functions (`uint16`) as the value.

There are 10 word-functions which can be represented by 10 bit length binary
value, as stated in the following list respectively:
Noun, Pronoun, Verb, Adj, Adverb, Conjunction, Preposition, Interjection,
Numeric, Articula

Note that word-functions are meant to be used in sentence or phrase structure
recognition. As for the basic task like stemming, it will not be used.