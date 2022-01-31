module testProject

go 1.15

require (
	example.com/testProject/newModule v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/jinzhu/gorm v1.9.16 // indirect
)

replace example.com/testProject/newModule => ./newModule
