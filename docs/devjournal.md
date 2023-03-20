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

### 2023-03-10
Morning :
- progressed on the basic design and layout of the website, needed elements are starting to have their place

Afternoon :
- stylized things such as the search bar, inputs...


Once I make everything work properly with the queries and such, and the front-end is responsive, I will start linking everything together, possibly merge the two "playground" branches.

### 2023-03-13
Morning :
After the decision to restart the front design from scratch, I got started in class, what I worked on:
- logo and icons
- colors used
- navigation bar
- cards display

Afternoon :
- finalized the brand design
- laid out more things such as the footer
- decided on having different pages for each type of search to fill the navbar a bit

### 2023-03-14
Morning :
- preshot the handling of routes that I will focus on later, before going to class

Afternoon and evening:
- finished up the footer
- progressed in the responsive: handling the navbars, the sizes...
- cleaned up the code slightly, will complete that later, putting this here as a reminder

### 2023-03-15
Evening :
- started laying out the card details page
- handling added for it

### 2023-03-16
Afternoon :
- finally merged most of the front with the go code
- added more css for the card details page
- did some responsive

### 2023-03-18
Afternoon/evening :
- route handling for most of the pages added
- changed the redirections in html to adapt to the routes
- set up the type search page
- did some more responsive for added features

### 2023-03-19
Afternoon/evening :
- search query by clicking on set image and redirecting to cards of selected set
- types checkbox filter done, basic
- stylized some buttons
- card details more furnished, responsive and easy on the eye compared to before
- changed folder structure

### 2023-03-20
Morning :
- filled the homepage with API information, not done at all, could not take the time for it before turning everything in
- still has some kind of layout, would like to appropriate it more one day, add toggles
- checked if everything worked fine enough
- cleaned up html and css
- brief commenting
- cleaned up unused elements in assets

## TIME TRACKED
Approximations of time spent on the project each day (usually tracked)
| Date  | Time spent |
| ----- | ------------ |
| 2023-03-06 | Around 3 hours |
| 2023-03-07 | A bit more than 8 hours |
| 2023-03-08 | A bit more than 3 hours |
| 2023-03-09 | Around 5 hours |
| 2023-03-10 | Around 4 hours |
| 2023-03-13 | Around 5 hours |
| 2023-03-14 | Around 5 hours |
| 2023-03-15 | Around 2 hours |
| 2023-03-16 | Around 2 hours |
| 2023-03-18 | A bit more than 5 hours |
| 2023-03-19 | Around 4 hours |
| 2023-03-20 | Around 4 hours |

