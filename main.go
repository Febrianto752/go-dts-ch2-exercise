package main

import "tmp_latihan/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
