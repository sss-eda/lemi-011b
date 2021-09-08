package main

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

var (
	serialPort *serial.Port
	err        error
)

var (
	serialName string = "/dev/ttyUSB0"
	serialBaud int    = 115200
)

func init() {
	serialPort, err = serial.OpenPort(
		&serial.Config{
			Name: serialName,
			Baud: serialBaud,
		},
	)
	if err != nil {
		log.Error("Could not open serial port.")
	}
	time.Sleep(time.Second)
}

func main() {
	defer serialPort.Close()

	//ctx, cancel := context.WithTimeout(
	//    context.Background(),
	//    time.Second*10,
	//)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := readSerial(
		ctx,
		serialPort,
	)

	// logSerial(
	//     ctx,
	//     ch1,
	// )

	writeFile(
		ctx,
		ch,
	)

	<-ctx.Done()
}

func readSerial(
	ctx context.Context,
	sp *serial.Port,
) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		err := sp.Flush()
		if err != nil {
			log.Error("Could not flush serial port.")
		}
		buf := make([]byte, 128)
		result := []byte{}
		for {
			n, err := sp.Read(buf)
			if err != nil {
				log.Error("Could not read from serial port.")
			}
			// Since the result will be used in other goroutines, we can't
			// send a pointer the buffer directly, since the buffer might be
			// changed in the meantime.
			for i := 0; i < n; i++ {
				switch buf[i] {
				case '\r':
				case '\n':
					s := time.Now().UTC().Format("2006-01-02 15:04:05.000000") +
						", " + string(result)
					select {
					case <-ctx.Done():
					case out <- []byte(s):
						result = []byte{}
					}
				default:
					result = append(result, buf[i])
				}
			}
		}
	}()

	return out
}

func timestampSerial(
	ctx context.Context,
	in <-chan []byte,
) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for data := range orDone(ctx, in) {
			timestamp := time.Now().UTC().Format("2006-01-02 15:04:05") + ", "
			out <- append(
				[]byte(timestamp),
				data...,
			)
		}
	}()

	return out
}

// func logSerial(
//     ctx context.Context,
//     in <-chan []byte,
// ) {
//     go func() {
//         for s := range orDone(ctx, in) {
//             log.WithFields(log.Fields{
//                 "timestamp":
//             })

//             log.Info(string(s))
//         }
//     }()
// }

func getBase(t time.Time) string {
	return ("marlem1_" +
		t.Format("2006-01-02") +
		".dat")
}

func getPath(t time.Time) string {
	return ("/data/MARL111/R/" +
		t.Format("2006/01/02/"))
}

func writeFile(
	ctx context.Context,
	in <-chan []byte,
) {
	// TODO: Get last saved file and parse filename to get last file.
	go func() {
		// prev := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
		for s := range orDone(ctx, in) {
			now := time.Now().UTC()

			path := getPath(now)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err := os.MkdirAll(
					path,
					0755,
				)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
						"path":  path,
					}).Error("Could not create new path.")
				} else {
					log.WithFields(log.Fields{
						"path": path,
					}).Info("New directory created.")
				}
			}

			base := getBase(now)
			if _, err := os.Stat(path + base); os.IsNotExist(err) {
				log.WithFields(log.Fields{
					"path": path,
					"base": base,
				}).Info("Creating a new file.")
			}

			// prev = now
			file, err := os.OpenFile(
				path+base,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY,
				0755,
			)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"path":  path,
					"base":  base,
				}).Error("Could not open file.")
				continue
			}
			_, err = file.WriteString(string(s) + "\n")
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"data":  string(s),
				}).Error("Could not write to file.")
			} else {
				log.WithFields(log.Fields{
					"data": string(s),
				}).Info("Data written to file.")
			}
			file.Close()
		}
	}()
}

func orDone(
	ctx context.Context,
	in <-chan []byte,
) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
				}
			}
		}
	}()

	return out
}

func tee(
	ctx context.Context,
	in <-chan []byte,
) (<-chan []byte, <-chan []byte) {
	out1 := make(chan []byte)
	out2 := make(chan []byte)
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(ctx, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
