package pkg

import (
	"context"
	"io"
	"os"
)

func CreateMigrationCtx(ctx context.Context, log *Log, fullPath string) string {
	f, err := os.Create(fullPath)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return ""
	}
	defer f.Close()

	err = writeTemplate(f)
	if err != nil {
		log.Printf("Failed to write template into file: %v", err)
		return ""
	}

	log.Printf("Migration file created: %s", fullPath)

	return fullPath
}

func writeTemplate(f io.Writer) error {
	bytesToWrite := [][]byte{
		stringWithPrefix("f-start"),
		[]byte("\n"),
		stringWithPrefix("up-start"),
		[]byte("SELECT 'up SQL query'; \n"),
		stringWithPrefix("up-end"),
		[]byte("\n"),
		stringWithPrefix("down-start"),
		[]byte("SELECT 'down SQL query'; \n"),
		stringWithPrefix("down-end"),
		[]byte("\n"),
		stringWithPrefix("f-end"),
	}

	for _, b := range bytesToWrite {
		_, err := f.Write(b)
		if err != nil {
			return err
		}
	}

	return nil
}

func stringWithPrefix(s string) []byte {
	return bytesWithPrefix([]byte(string(s + "\n")))
}

func bytesWithPrefix(b []byte) []byte {
	res := make([]byte, len(b)+len(Prefix))
	copy(res, []byte(Prefix))
	copy(res[len(Prefix):], b)
	return res
}
