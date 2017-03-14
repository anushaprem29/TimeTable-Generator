package main
import "fmt"
import "os"
import "math/rand"
type period struct {
    name string
    nop  int
}
func main(){
  var subjects int
  var days int
  var per int
  var sum int
  var random int
  fmt.Print("Enter number of Subjects: ")
  fmt.Scanln(&subjects)
  fmt.Print("Enter number of Working days: ")
  fmt.Scanln(&days)
  fmt.Print("Enter number of periods for every working day: ")
  fmt.Scanln(&per)
  var timetable = make([]period, subjects)
  for iter := 0; iter < subjects; iter++ {
    fmt.Println("Enter name of Subject",iter+1)
    fmt.Scanln(&timetable[iter].name)
    fmt.Println("Enter number of periods for Subject ",timetable[iter].name)
    fmt.Scanln(&timetable[iter].nop)
    sum += timetable[iter].nop
  }
  if sum != (per*days){
    fmt.Println("Invalid data provided")
    os.Exit(1)
  }
  for iter := 0; iter < sum; {
    random = rand.Intn(subjects)
    if(timetable[random].nop>0){
      if iter % per == 0 {
        fmt.Print("\nDay ",(iter/per)+1)
        fmt.Print("\t\t")
      }
      fmt.Print(timetable[random].name)
      fmt.Print("\t\t")
      timetable[random].nop -=1
      iter++
    }
  }
}
