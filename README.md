# English - Gopher Dictionary

## Description

The project represents a service that exposes an API for translating english words and sentences to the ancient gopher language. It uses several translation rules:

The language that the gophers speak is a modified version of English and has a few simple rules:

* If a word starts with a vowel letter, add prefix g to the word:
    * `apple -> gapple`
    * `ear -> gear`
    * `oak -> goak`
* If a word starts with the consonant letters xr , add the prefix ge to the beginning of the word:
    * `xray -> gexray`
* If a word starts with a consonant sound, move it to the end of the word and then add ogo suffix to the word. Consonant sounds can be made up of multiple
consonants i.e., a consonant cluster:
    * `chair -> airchogo`
* If a word starts with a consonant sound followed by qu , move it to the end of the word, and then add ogo suffix to the word:
    * `square -> aresquogo`

## Endpoints

The API consists of three endpoints:

* `/word`     for translating a single word
* `/sentence` for translating a whole sentence
* `/history`  for retrieving a history of all translations (sorted alphabetically by the english word/sentence)
