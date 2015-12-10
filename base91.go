package base91

var enctab = []string{
  "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
  "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
  "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
  "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
  "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "#", "$",
  "%", "&", "(", ")", "*", "+", ",", ".", "/", ":", ";", "<", "=",
  ">", "?", "@", "[", "]", "^", "_", "`", "{", "|", "}", "~", "\"",
}
var dectab = []byte{
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 62, 90, 63, 64, 65, 66, 91, 67, 68, 69, 70, 71, 91, 72, 73,
  52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 74, 75, 76, 77, 78, 79,
  80,  0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14,
  15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 81, 91, 82, 83, 84,
  85, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
  41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 86, 87, 88, 89, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
  91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
}

//
// encode bytes to base91'd string
//
func Encode(data []byte) (out string) {
  nbits := uint32(0)
  queue := uint64(0)
  for _, b := range data {
    queue |= uint64(b) << nbits
    nbits += 8
    if nbits > 13 {/* enough bits in queue */
      val := uint32(queue & 8191)
      
      if val > 88 {
        queue >>= 13
        nbits -= 13
      } else {/* we can take 14 bits */
        val = uint32(queue & 16383)
        queue >>= 14
        nbits -= 14
      }
      out += enctab[val % 91] + enctab[val / 91]
    }
  }
  if nbits != 0 {
    out += enctab[queue % 91]
    if nbits > 7 || queue > 90 {
      out +=  enctab[queue / 91]
    }
  }
  return
}

// decode base91'd string to bytes
// strips invalid characters
func Decode(data string) (out []byte) {
  d := uint32(0)
  queue := uint64(0)
  nbits := uint32(0)
  val := -1
  for _, c := range data {
    d = uint32(dectab[c])
    if  d == 91 {
      continue;/* ignore non-alphabet chars */
    }
    if val == -1 {
      val = int(d)
    } else {
      val += int(d) * 91
      queue |= uint64(val << nbits)
      if (val & 8191) > 88 {
        nbits += 13
      } else {
        nbits += 14
      }
      for {
        out = append(out, byte(queue))
        queue >>= 8;
        nbits -= 8;
        if nbits <= 7 {
          break
        }
      }
      val = -1
    }
  }
  if val != -1 {
    out = append(out, byte(queue | uint64(val << nbits)))
  }
  return
}
