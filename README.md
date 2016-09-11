# similar-word-finder

A small web service for printing similar words in the English language.

## Run app

* make run

## test app

* make test

## Algorithm description

1. Load data to internal DB.
 * Read each word
 * Assign a unique key - the word sorted by A,B,C...
 * Save word to db under the key.

1. When a word is requested
 * Sort word to get key.
 * Use key to get all words in group.
 * Extract the word from the group.
 * Return request.


1. Big O calculation:
 * Since db is implemented by a Golang hash map, Retrieval is O(1).