package main

func main() {
	app := newApp()
	if err := app.Run(":8081"); err != nil {
		panic("failed to run app: " + err.Error())
	}
}
