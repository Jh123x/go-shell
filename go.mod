module github.com/Jh123x/go-shell

go 1.19

replace commands => ./commands

replace github.com/Jh123x/go-shell/ => ./

require commands v0.0.0-00010101000000-000000000000

require github.com/Jh123x/go-validate v0.0.0-20231021183158-57da40742a50 // indirect
