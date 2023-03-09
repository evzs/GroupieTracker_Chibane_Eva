# "DEV JOURNAL"
Daily logs of my progression on the project and contributions to it.

## DETAILED REPORT
More or less
### 2023-03-06
- API choice
- Repository setup

### 2023-03-07
Morning :
- cleaned the folder tree as well as I could until further edits
- brainstormed on some design ideas for the front, still undecided
- set up two "playground" branches, one for the front and possible ideas, another for handling the API with Go to try and understand better how it works before diving in

Afternoon :
- tried to tackle on the Go, struggling a bit with how everything works. I do have a clear visualization of what needs to be done and which elements are required for me to get started.
- getting a better grasp on how to link structures and external data, do need to polish up on how to handle API requests
- wrapping my head around templates too, trying to put everything together
- might start doing this with the actual API (there is actual documentation and I made something work, but at the moment I do have a lot of brain fog and cannot get myself into it without physically straining my neurons)

Evening :
- managed to make the templates + server + API requests work together
- I now have an idea on how to present data, I need to sort out which will be relevant to the content of the website
- thinking about implementing a search bar/toggle filter system to search for specific cards, I do not know at the moment what I would like to present

### 2023-03-08
Morning :
- finished up on basic handling of the data
- html and go are properly linked together
- the structs are kind of a big dump, I will have to sort out which ones I will definitely keep for usage
- for now, you can actually search a pokemon name and get its cards listed: I used the documentation of the API for Golang to make the requests since they have their own package
- on that, I will have to learn all its little quirks to understand better what I'm going to do with it

Afternoon :
- decided to get back on the design, truly hesitating on a dark or light theme, as well as wondering if I should do something clear and concise or if I want to pimp it a bit and make it extra for no reason, I do not know if it corresponds to the criteria for these kind of websites so I'm still undecided...
- started laying out a bit where things should go, no real content nor responsive yet (header, body, footer... still wondering if I should have other pages other than the ones required)



Not a very productive day overall and I might stop there, will gain more energy for tomorrow.

### 2023-03-09
Morning :
- added some changes to the front, trying to fill up the navbar with content to navigate through

Afternoon :
- finally made search queries work independently from each other, or together
- tested on pokemon names on the expansion names, so if you look for example pikachu you'll get all the pikachu cards (within the set limit for now) and if you look for a specific expansion pack it will show all the cards from this one. Managed to make it work when you look for a pokemon and a specific expansion pack, if card(s) exist the data is returned.
- trying to make everything work together little by little, there are probably more than a thousands of possibilities of search queries to be made so I would like to decide on how to present it

Evening :
- added the select type functionality, you can check boxes on the page to select it



Overall not much done, I do undestand things better gradually however, will read more of the API documentation before stopping.


## TIME TRACKED
Approximations of time spent on the project each day (usually tracked)
| Date  | Time spent |
| ----- | ------------ |
| 2023-03-06 | Around 3 hours |
| 2023-03-07 | A bit more than 8 hours |
| 2023-03-08 | A bit more than 3 hours |
| 2023-03-09 | Around 5 hours |
