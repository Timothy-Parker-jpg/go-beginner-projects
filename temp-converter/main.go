package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n\n\n*************Temp Converter (type 'exit' to end)*************")
		fmt.Println("\tType your temperature value and unit Ex: 25 c")

		fmt.Print(">> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Stdin read fail.\nTry again.")
			continue
		}
		line = strings.TrimSpace(line)

		if strings.ToLower(line) == "exit" {
			os.Exit(0)
		}
		args := strings.Fields(line)

		if len(args) < 2 {
			fmt.Println("Missing Arguments. Make sure to use {temp} {unit} positional arguments.")
			continue
		}

		arg1, err := strconv.ParseFloat(args[0], 64)

		if err != nil {
			fmt.Printf("Failed to parse '%v' to float. Make sure to input a digit\n", args[0])
			continue
		}

		arg2 := strings.ToLower(args[1])

		c, f, k, err := Convert(arg1, arg2)

		if err != nil {
			fmt.Printf("Convert function failed. %v\n", err)
			continue
		}

		format := "%.2f°C = %.2f°F = %.2fK\n"

		fmt.Printf(format, c, f, k)
	}

}
