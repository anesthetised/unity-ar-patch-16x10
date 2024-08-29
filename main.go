package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func run(ctx context.Context, conf *Config) error {
	var err error

	if err = conf.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	f, err := os.OpenFile(conf.Filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func() { _ = f.Close() }()

	offset, err := scanFile(ctx, f, expectedBytes)
	if err != nil {
		return fmt.Errorf("scan file: %w", err)
	}

	fmt.Printf("found expected bytes offset: %x\n", offset)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if _, err = f.WriteAt(patchBytes, offset); err != nil {
		return fmt.Errorf("write patch bytes to file: %w", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}

	return nil
}

func main() {
	var conf Config

	flag.StringVar(&conf.Filename, "f", "", "filename")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := run(ctx, &conf)
	if errors.Is(err, context.Canceled) {
		fmt.Println("canceled")
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf("failed to %w", err))
		return
	}

	fmt.Println("done")
}
