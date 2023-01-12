package main

import (
	"fmt"
	"redisDBService/routers"
)

func main() {

	r := routers.SetRouter()
	fmt.Println(r.Run(":8000"))

}
