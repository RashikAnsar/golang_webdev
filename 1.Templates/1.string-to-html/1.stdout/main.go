// The following command generates html file
// go run main.go > index.html
package main

import (
	"fmt"
)

func main() {
	name := "Rashik Ansar"

	tpl := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Hello world</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`

	fmt.Println(tpl)
}
