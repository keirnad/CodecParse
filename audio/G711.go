package audio

import (
	"io"
	"os"
)

// G711
// 160 octets payload (160 bytes)

func ALawParse(audioFile string, handler func(payload []byte)) {

	f, err := os.Open(audioFile)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	buffer := make([]byte, 160)

	for {
		sample, err := f.Read(buffer)

		if err != nil {
			panic(err)
		}

		if err != nil && err != io.EOF {
			panic(err)
		}

		if sample == 0 {
			break
		}

		handler(buffer)
	}
}
