package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func originSolution(out io.Writer) {
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	readingBuf := make([]byte, 1) //read file by one letter only

	words := make([]record, 0)
	reader := reader{words: &words} //creating reader

	writingBuf := make([]byte, 0)
	writer := writer{&writingBuf} //writer

	ch := make(chan []byte) //channel that we will use to pass slices of bytes from writer to reader
	//btw reader listens in range of elements that are passed to channel, it will stop working when there are no elements left, so we don't need any wait groups

	go func() {
		for {
			//reading file's letters one by one
			n, err := file.Read(readingBuf)

			if n > 0 {
				byteVal := readingBuf[0]
				if byteVal >= 65 && byteVal <= 90 { //if symbol is uppercase letter

					byteVal = byteVal + 32
					writer.write_to_temp_buf(byteVal) //writing to temporary buffer

				} else if byteVal >= 97 && byteVal <= 122 { //if symbol is lowercase letter

					writer.write_to_temp_buf(byteVal) //writing to temporary buffer

				} else if byteVal == 32 && len(writingBuf) != 0 { //if symbol is [space], and we have letters in our buffer

					writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer

				} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(writingBuf) != 0 { //if symbol is any other than letter or space, and we have letters in our buffer

					writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer

				} else {
					continue
				}
			}

			if err == io.EOF {
				writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer
				break
			}
		}
		close(ch) //close channel, so our that our reader will stop working after there are no elements left, in other case reader will cause deadlock
	}()

	reader.read_from_chan(ch) //reading from channel in range of elements in channel

	reader.get20mostfrequentwords() //getting 20 most frequent words, and write it to rating slice
	reader.print()
}

func MukhenzoSolution1(out io.Writer) {
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	readingBuf := make([]byte, 1) //read file by one letter only

	words := make([]record, 0)
	reader := reader{words: &words} //creating reader

	writingBuf := make([]byte, 0)

	ch := make(chan []byte) //channel that we will use to pass slices of bytes from writer to reader
	//btw reader listens in range of elements that are passed to channel, it will stop working when there are no elements left, so we don't need any wait groups

	r := bufio.NewReader(file)
	go func() {
		for {
			//reading file's letters one by one
			n, err := r.Read(readingBuf)

			if n > 0 {
				byteVal := readingBuf[0]
				if byteVal >= 65 && byteVal <= 90 { //if symbol is uppercase letter

					byteVal = byteVal + 32
					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal >= 97 && byteVal <= 122 { //if symbol is lowercase letter

					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal == 32 && len(writingBuf) != 0 { //if symbol is [space], and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(writingBuf) != 0 { //if symbol is any other than letter or space, and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else {
					continue
				}
			}

			if err == io.EOF {
				ch <- writingBuf
				writingBuf = nil //send temporary buffer content to channel, empty the temporary buffer
				break
			}
		}
		close(ch) //close channel, so our that our reader will stop working after there are no elements left, in other case reader will cause deadlock
	}()

	reader.read_from_chan(ch) //reading from channel in range of elements in channel

	reader.get20mostfrequentwords() //getting 20 most frequent words, and write it to rating slice
	reader.print()
}

func MukhenzoSolution2(out io.Writer) {
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	readingBuf := make([]byte, 1) //read file by one letter only

	words := make([]recordMy, 0)
	reader := readerMy{words: &words} //creating reader

	writingBuf := make([]byte, 0)

	ch := make(chan []byte) //channel that we will use to pass slices of bytes from writer to reader
	//btw reader listens in range of elements that are passed to channel, it will stop working when there are no elements left, so we don't need any wait groups

	r := bufio.NewReader(file)
	go func() {
		for {
			//reading file's letters one by one
			n, err := r.Read(readingBuf)

			if n > 0 {
				byteVal := readingBuf[0]

				if byteVal >= 65 && byteVal <= 90 { //if symbol is uppercase letter

					byteVal = byteVal + 32
					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal >= 97 && byteVal <= 122 { //if symbol is lowercase letter

					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal == 32 && len(writingBuf) != 0 { //if symbol is [space], and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(writingBuf) != 0 { //if symbol is any other than letter or space, and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else {
					continue
				}
			}

			if err == io.EOF {
				ch <- writingBuf
				writingBuf = nil //send temporary buffer content to channel, empty the temporary buffer
				break
			}
		}
		close(ch) //close channel, so our that our reader will stop working after there are no elements left, in other case reader will cause deadlock
	}()

	reader.read_from_chan(ch) //reading from channel in range of elements in channel

	reader.get20mostfrequentwordsMy() //getting 20 most frequent words, and write it to rating slice
}

func MukhenzoSolution3(out io.Writer) {
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	readingBuf := make([]byte, 1) //read file by one letter only

	words := make([]recordMy, 0)
	reader := readerMy{words: &words} //creating reader

	writingBuf := make([]byte, 0)

	ch := make(chan []byte) //channel that we will use to pass slices of bytes from writer to reader
	//btw reader listens in range of elements that are passed to channel, it will stop working when there are no elements left, so we don't need any wait groups

	r := bufio.NewReader(file)
	go func() {
		for {
			//reading file's letters one by one
			n, err := r.Read(readingBuf)

			if n > 0 {
				byteVal := readingBuf[0]

				if byteVal >= 65 && byteVal <= 90 { //if symbol is uppercase letter

					byteVal = byteVal + 32
					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal >= 97 && byteVal <= 122 { //if symbol is lowercase letter

					writingBuf = append(writingBuf, byteVal) //writing to temporary buffer

				} else if byteVal == 32 && len(writingBuf) != 0 { //if symbol is [space], and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(writingBuf) != 0 { //if symbol is any other than letter or space, and we have letters in our buffer

					ch <- writingBuf
					writingBuf = nil

				} else {
					continue
				}
			}

			if err == io.EOF {
				ch <- writingBuf
				writingBuf = nil //send temporary buffer content to channel, empty the temporary buffer
				break
			}
		}
		close(ch) //close channel, so our that our reader will stop working after there are no elements left, in other case reader will cause deadlock
	}()

	reader.read_from_chan(ch) //reading from channel in range of elements in channel

	reader.get20mostfrequentwordsMy() //getting 20 most frequent words, and write it to rating slice
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		originSolution(ioutil.Discard)
	}
}

func BenchmarkMukhenzo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MukhenzoSolution1(ioutil.Discard)
	}
}

func BenchmarkMukhenzo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MukhenzoSolution2(ioutil.Discard)
	}
}

// func BenchmarkMukhenzo3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		MukhenzoSolution3(ioutil.Discard)
// 	}
// }
