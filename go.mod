module example/main

go 1.16

replace example/hello => ./hello

replace example/goodbye => ./goodbye

require (
	example/goodbye v0.0.0-00010101000000-000000000000
	example/hello v0.0.0-00010101000000-000000000000
)
