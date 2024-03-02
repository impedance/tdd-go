package iteration

const repeatCount = 5

func Repeat(text string, counter int) string {

  if text == "other" {
    return "not count"
  }
  var repeated string
  for i := 0; i < counter; i++ {
    repeated += text
  }
  return repeated
}
