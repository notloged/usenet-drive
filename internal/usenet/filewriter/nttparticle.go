package filewriter

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io"

	"github.com/javi11/usenet-drive/pkg/yenc"
)

type ArticleData struct {
	partNum   int64
	partTotal int64
	partSize  int64
	partBegin int64
	partEnd   int64
	fileNum   int
	fileTotal int
	fileSize  int64
	fileName  string
	poster    string
	group     string
	msgId     string
}

func ArticleToReader(p []byte, data ArticleData) (io.Reader, error) {
	subj := fmt.Sprintf(
		"[%d/%d] - \"%s\" yEnc (%d/%d)",
		data.fileNum,
		data.fileTotal,
		data.fileName,
		data.partNum,
		data.partTotal,
	)
	header := fmt.Sprintf("From: %s\r\nNewsgroups: %s\r\nMessage-ID: <%s>\r\nX-Newsposter: UsenetDrive\r\nSubject: %s\r\n\r\n=ybegin part=%d total=%d line=128 size=%d name=%s\r\n=ypart begin=%d end=%d\r\n",
		data.poster,
		data.group,
		data.msgId,
		subj,
		data.partNum,
		data.partTotal,
		data.fileSize,
		data.fileName,
		data.partBegin+1,
		data.partEnd,
	)

	// yEnc end line
	h := crc32.NewIEEE()
	_, err := h.Write(p)
	if err != nil {
		return nil, err
	}
	footer := fmt.Sprintf("=yend size=%d part=%d pcrc32=%08X\r\n", data.partSize, data.partNum, h.Sum32())

	buf := bytes.NewBuffer(make([]byte, 0))

	_, err = buf.WriteString(header)
	if err != nil {
		return nil, err
	}
	// Encoded data
	err = yenc.Encode(p, buf)
	if err != nil {
		return nil, err
	}
	_, err = buf.WriteString(footer)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
