package main

import (
    "os"

    "hashcode2021/m/v2/src/parser"
)

func main() {
    algo := os.Args[1]
    name := os.Args[2]

    input := parser.LoadInput("./data", name)

    println(algo, name, input.String())

}
