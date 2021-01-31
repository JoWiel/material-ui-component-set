package generator

import {
	"fmt"
	"os"
	"exec"
	"bytes"
	"io"
}

func generator() {
	store := "/componentStore/ls"
    cmd, err := exec.Run(store, []string{store, "-l"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)

    if (err != nil) {
       fmt.Fprintln(os.Stderr, err.String())
       return
    }

    var b bytes.Buffer
    io.Copy(&b, cmd.Stdout)
    fmt.Println(b.String())

    cmd.Close()
}