module github.com/AlejandroArcoPu/go/create_go_module/hello

go 1.26.4

require example.com/greetings v0.0.0-00010101000000-000000000000

replace example.com/greetings => ../greetings_local
