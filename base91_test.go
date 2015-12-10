

package base91


import (
  "testing"
  "bytes"
  "crypto/rand"
  "io"
)

func TestBase91(t *testing.T) {
  var buff [10]byte
  io.ReadFull(rand.Reader, buff[:])
  out := Encode(buff[:])
  t.Logf("%q -> %q",buff[:], out)
  buff2 := Decode(out)
  if ! bytes.Equal(buff2, buff[:]) {
    t.Logf("%q vs %q", buff2, buff)
    t.Fail()
  }
}
