package audio

import (
	"io"
	"math/rand/v2"
	"os"
)

// G711
// 160 octets payload (160 bytes)

func ALawParse(audioFile string, handler func(payload []byte, sequenceNumber uint16, timestamp uint32, ssrc uint32)) {

	f, err := os.Open(audioFile)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	sequenceNumber := uint16(1)
	timestamp := uint32(90000)
	ssrc := rand.Uint32()

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

		sequenceNumber++
		timestamp += 3000

		handler(buffer, sequenceNumber, timestamp, ssrc)
	}
}
