package main

func main() {
	blockchain := NewBlockchain()
	defer blockchain.database.Close()

	cli := CLI{blockchain}
	cli.Run()
}
