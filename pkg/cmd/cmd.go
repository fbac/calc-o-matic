package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fbac/calc-o-matic/pkg/eval"
	"github.com/fbac/calc-o-matic/pkg/parser"
)

// Init returns an interactive cmd and reads the input as string until \n
// CTRL + C ends the interactive shell
func Init() {
	ctx := newCancelableContext()

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("calc-o-matic /> ")
		inputExpr, err := r.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		inputExpr = strings.TrimSuffix(inputExpr, "\n")
		tokenStack, err := parser.ParseExpr(inputExpr)
		if err != nil {
			log.Printf("error: %v", err)
		}

		ast, err := parser.CreateAST(tokenStack)
		if err != nil {
			log.Printf("error: %v", err)
		}

		fmt.Printf("%v\n", eval.Calculate(ast))
	}

	<-ctx.Done()
}

// newCancelableContext returns a context that gets canceled by a SIGINT/SIGTERM
func newCancelableContext() context.Context {
	doneCh := make(chan os.Signal, 1)
	signal.Notify(doneCh, os.Interrupt, syscall.SIGTERM)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-doneCh
		fmt.Printf("\nexiting %v", filepath.Base(os.Args[0]))
		cancel()
		os.Exit(1)
	}()

	return ctx
}
