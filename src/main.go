package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ythosa/pukiclang/src/repl"
)

const PUKICLANG = `
⠀ ⠀⠀⠀⠀⠀⠀   ／＞　 フ
　　　　　| 　_　 _|
　 　　　／'ミ _x 彡
　　 　 /　　　 　 |
　　　 /　 ヽ　　 ﾉ
　／￣|　　 |　|　|
　| (￣ヽ＿_ヽ_)_)
　＼二つ
`

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Print(PUKICLANG)
	fmt.Printf("Hello %s! This is the pukiclang programming language!\n",
		u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
