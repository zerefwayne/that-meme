package main

func main() {

	var connections Connections

	connections.ConnectDatabase()
	connections.ConnectCache()

	defer connections.Cache.Close()

}
