package generator

import (
	"fmt"
	"os"
	"os/exec"
)

// SetGenerator generates custom component sets from the selected sets
func SetGenerator(path string) {
	// store := "/componentStore/ls"
	// cmd, err := exec.Run(store, []string{store, "-l"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)

	// if (err != nil) {
	//    fmt.Fprintln(os.Stderr, err.String())
	//    return
	// }

	// var b bytes.Buffer
	// io.Copy(&b, cmd.Stdout)
	// fmt.Println(b.String())

	// cmd.Close()
	// command := exec.Command("../app/node/npm","mkdir newset && cd newset && yarn && npx webpack --config webpack.config.js")
	projectRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
    cmdLine := projectRoot + "/generator/generator.sh"
	// nodeExecPath, err := exec.LookPath("node")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// line := "cd " + projectRoot + " && pwd"
	command := exec.Command("/bin/sh", cmdLine, path)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Run the command
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("command succesfully ran:" + cmdLine)
}

// StaticGenerator copies the static files
func StaticGenerator(path string) {
	projectRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// nodeExecPath, err := exec.LookPath("node")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	cmdLine := projectRoot + "/generator/static.sh"

	// line := "cd " + projectRoot + " && pwd"
	command := exec.Command("/bin/sh", cmdLine, path)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Run the command
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("command succesfully ran:")
}
