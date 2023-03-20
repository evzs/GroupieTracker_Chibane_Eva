# Groupie Tracker
As a given assignment, the goal of Groupie Tracker was to create a website,
hosted with GoLang, and this latter had to display and depict data gathered
from an API from a given list.

## About the project
I personally chose the Pokémon TCG API. As I used to love collecting Pokémon cards in my early age, I figured it would be a way to reconnect with this part of the past. There is also a handful of possibilities with this API, and the documentation is well supllied.

## How to use the website
As most of the data is handled with a Golang server, in order to use the website and make search search queries you will have to :
- clone the repository
```
https://github.com/evzs/GroupieTracker_Chibane_Eva.git
```
- open the terminal in the src folder's root and run the following command:
```
go run .
```
You should be able to view and use the website on localhost:8080.

## To go a little further
The current queries one can do with this program only consist for now of:
- searching a pokemon card by name of the pokemon
- clicking on the card and getting its details
- clicking on an expansion set and getting all the cards from this expansion set
- searhcing for pokemon cards by type of the pokemon (kind of randomly chosen)

If you'd like to play with it a bit, take a look at the API documentation!!!


## Credits
This website was created by [evzs](https://github.com/evzs) and uses the [Pokémon TCG API](https://pokemontcg.io) to fetch data about Pokémon cards.
