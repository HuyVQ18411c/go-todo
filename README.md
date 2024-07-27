# Todo APIs with Golang
A pet project to get used to Go (v1.22.2)

## Dependencies
- GORM: it was initially added when I have no knowledge about how to access database with Go and it is the first result pop up when I search for Go ORM. After implementing it, I feel like we can Go with built-in package and install the target driver instead, which will help us to swap to any database that we want. In the scope of this ticket, I'm using SQLite. 

There are a controversy around using "ORM" (in any languages), and Go has `sqlc` which is also very popular, however since this is a learning project, I just want to go through it as fast as possible. 

## Reflection model 
- Go reflection model is cool, comparing to how we can achieve the same thing in Python (which I mainly work with). It's very interesting how it can keep the data still strongly type, but still map to the correct field at runtime.

## Built-in template language
- Though it's not a complete one, but still nice to see Go provides me an option. Although I'm not using it in this project, I gave it a few tries.

## HTTP package
- I love `net/http` package, it's easy to use, nice features, even better with new versions. I feel like, I don't even need 3rd party framework to bulid such a simple API like this. Still, I haven't exposed to Go enough to feel all the possible pains. 

## Testing
- I wrote a test (which is not related to API's functionalities) but still see a lot of similarities between them and Python/C#/JavaScript. I haven't had many chances to work with procedure language before, therefore I'm not sure if this a standard way to write them.

## Modulize Go project
- I'm still not so sure about how can I organize a Go project, so I just went for the one that I'm most comfortable with. I previewed a few projects, and to be honest, it's a little bit overwhelmed for me and for such a small project like this.

In conclusion, I love Go.
