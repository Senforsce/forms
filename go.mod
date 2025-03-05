module github.com/senforsce/forms

go 1.23.3

replace github.com/senforsce/tndr0cean => ../tndr0cean

replace github.com/senforsce/tndr => ../tndr

replace github.com/senforsce/forms => ./

replace github.com/senforsce/o => ../o

require (
	github.com/senforsce/o v0.0.0-00010101000000-000000000000
	github.com/senforsce/tndr v0.0.6
	github.com/senforsce/tndr0cean v0.0.3
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
)
