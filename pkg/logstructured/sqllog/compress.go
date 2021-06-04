package sqllog

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
)

var toCompress = true

func compressGzip(data []byte) []byte {

	if toCompress {
		start := time.Now()

		var b bytes.Buffer
		gz := gzip.NewWriter(&b)

		// If there is any error in compressing the data, we return the data
		if _, err := gz.Write(data); err != nil {
			logrus.Tracef("COMPRESS ERROR %s", err)
			return data
		}

		// If there is any error in compressing the data, we return the data
		if err := gz.Close(); err != nil {
			logrus.Tracef("COMPRESS ERROR %s", err)
			return data
		}

		elapsed := time.Since(start)
		logrus.Tracef("COMPRESS OK [Original Size: %d] [Compressed Size: %d] took %s", len(data), len(b.Bytes()), elapsed)

		return b.Bytes()
	}
	return data

}

func uncompressGzip(data []byte) []byte {
	if toCompress {
		if data != nil {
			b := bytes.NewReader(data)
			reader, err := gzip.NewReader(b)
			if err != nil {
				logrus.Tracef("COMPRESS Invalid Reader", err)
				return data
			}
			uncompressed, err := ioutil.ReadAll(reader)

			if err != nil {
				logrus.Tracef("COMPRESS Invalid decompression", err)
				return data
			}

			return uncompressed
		}
	}

	return data
}
