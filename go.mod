module main

go 1.15

replace local.packages/sample => ./sample

require (
	google.golang.org/grpc v1.35.0 // indirect
	local.packages/sample v0.0.0-00010101000000-000000000000 // indirect
)
