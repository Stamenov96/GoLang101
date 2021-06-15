module github.com/Stamenov96/GoLang101/sort_skeleton/sort/sorting-service

go 1.16

replace github.com/Stamenov96/GoLang101/sort_skeleton/sort/gen => ../gen

require (
	github.com/Stamenov96/GoLang101/sort_skeleton/sort/gen v0.0.0-20210615162828-c9e3f4c42ea6
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.38.0
)
