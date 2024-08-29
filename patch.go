package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
)

var (
	expectedBytes = []byte{0x39, 0x8E, 0xE3, 0x3F}
	patchBytes    = []byte{0xCD, 0xCC, 0xCC, 0x3F}
)

var errNotFound = errors.New("expected bytes were not found")

func scanFile(ctx context.Context, rs io.ReadSeeker, search []byte) (int64, error) {
	br := bufio.NewReader(rs)

	idx := 0
	offset := int64(0)

	for idx < len(search) {
		select {
		case <-ctx.Done():
			return -1, ctx.Err()
		default:
		}

		b, err := br.ReadByte()
		if errors.Is(err, io.EOF) {
			return -1, errNotFound
		}
		if err != nil {
			return -1, fmt.Errorf("read byte: %w", err)
		}

		if search[idx] == b {
			idx++
		} else {
			idx = 0
		}

		offset++
	}

	return offset - int64(len(search)), nil
}
