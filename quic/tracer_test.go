package libp2pquic

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/klauspost/compress/zstd"

	"github.com/lucas-clemente/quic-go/logging"
)

func createLogDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "libp2p-quic-transport-test")
	require.NoError(t, err)
	t.Cleanup(func() { os.RemoveAll(dir) })
	return dir
}

func getFile(t *testing.T, dir string) os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, files, 1)
	return files[0]
}

func TestSaveQlog(t *testing.T) {
	qlogDir := createLogDir(t)
	logger := newQlogger(qlogDir, logging.PerspectiveServer, []byte{0xde, 0xad, 0xbe, 0xef})
	file := getFile(t, qlogDir)
	require.Equal(t, string(file.Name()[0]), ".")
	require.Truef(t, strings.HasSuffix(file.Name(), ".qlog.swp"), "expected %s to have the .qlog.swp file ending", file.Name())
	// close the logger. This should move the file.
	require.NoError(t, logger.Close())
	file = getFile(t, qlogDir)
	require.NotEqual(t, string(file.Name()[0]), ".")
	require.Truef(t, strings.HasSuffix(file.Name(), ".qlog.zst"), "expected %s to have the .qlog.zst file ending", file.Name())
	require.Contains(t, file.Name(), "server")
	require.Contains(t, file.Name(), "deadbeef")
}

func TestQlogBuffering(t *testing.T) {
	qlogDir := createLogDir(t)
	logger := newQlogger(qlogDir, logging.PerspectiveServer, []byte("connid"))
	initialSize := getFile(t, qlogDir).Size()
	// Do a small write.
	// Since the writter is buffered, this should not be written to disk yet.
	logger.Write([]byte("foobar"))
	require.Equal(t, getFile(t, qlogDir).Size(), initialSize)
	// Close the logger. This should flush the buffer to disk.
	require.NoError(t, logger.Close())
	finalSize := getFile(t, qlogDir).Size()
	t.Logf("initial log file size: %d, final log file size: %d\n", initialSize, finalSize)
	require.Greater(t, finalSize, initialSize)
}

func TestQlogCompression(t *testing.T) {
	qlogDir := createLogDir(t)
	logger := newQlogger(qlogDir, logging.PerspectiveServer, []byte("connid"))
	logger.Write([]byte("foobar"))
	require.NoError(t, logger.Close())
	compressed, err := ioutil.ReadFile(qlogDir + "/" + getFile(t, qlogDir).Name())
	require.NoError(t, err)
	require.NotEqual(t, compressed, "foobar")
	c, err := zstd.NewReader(bytes.NewReader(compressed))
	require.NoError(t, err)
	data, err := ioutil.ReadAll(c)
	require.NoError(t, err)
	require.Equal(t, data, []byte("foobar"))
}
