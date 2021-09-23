# English - Gopher Dictionary

## Description

The project represents a service written in GoLang that exposes an API for translating english words and sentences to the ancient gopher language. It uses several translation rules:

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

## Installation

The service can be started manually from the terminal by passing a single argument `--port 1234`. In case `--port` argument is not found the program will use the default port `8899`. There is also support for `docker-compose` command which will create a container and run the application inside it. `docker-compose up` will run the service on port `1234` (within the container) which will be mapped to the default port `8899` on the host.

## External packages and libraries

For this project no external libraries have been used. I decided to use plain vanilla Go language with the stuff provided in the standard library. It was hell of a fun!

## License

Please refer to the LICENSE file in the root directory of the project!