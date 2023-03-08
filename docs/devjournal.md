# "DEV JOURNAL"
Daily logs of my progression on the project and contributions to it.

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
- On that, I will have to learn all its little quirks to understand better what I'm going to do with it
