# fiber_website
A website using Fiber to work with a list of students

## Files
-[main.go](#main)
-[data.json](#data)
-[views](#views)
-[public](#public)
-[app](#app)

### main.go <div id="main"></div>
Here is written the head of programm. Adding work with HTML, starting the web-server, adding the ability to display static(css) files, working with all links and so on.

### data.json <div id="data"></div>
Here is store all data about students in json format

### views <div id="views"></div>
All html patterns written here

### public <div id="public"></div>
Here is static files, for now it is empty, but I think it will better in the future

### app <div id="app"></div>
-[data.go](#data.go)
-[funcs.go](#funcs.go)
-[model.go](#model.go)
-[routes.go](#routes.go)

#### data.go <div id="data.go"></div>
Here is written functions to READ, WRITE, UPDATE and REMOVE students from data and site.

#### funcs.go <div id="funcs.go"></div>
Here written functions for good working of html files, while we are READ, WRITE, UPDATE and REMOVE students

#### model.go <div id="model.go"></div>
There is one struct, which implements student's characteristics(name, surname, age, group)

#### routes.go <div id="routes.go"></div>
In this file, written GET and POST requests for web-site
