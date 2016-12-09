package main

import (
	. "fmt"

	"github.com/go-vgo/robotgo"
)

func screencap(filename string) {

	bit_map := robotgo.CaptureScreen(0, 0, 1430, 900)
	Println("...", bit_map)
	fx, fy := robotgo.FindBitmap(bit_map)
	Println("FindBitmap------", fx, fy)
	robotgo.SaveBitmap(bit_map, filename)
	robotgo.SetMouseDelay(1000)

}

func clickthumbnail(x int, y int) {
	robotgo.MoveMouse(x, y)//product row 1 left corner//
	Println(robotgo.GetPixelColor(x, y))
	robotgo.MouseClick()
	robotgo.MouseClick()
	robotgo.MouseClick()
	robotgo.SetMouseDelay(1000)
}

func exitdetailview() {
	robotgo.MoveMouse(179, 547)//exit detail view//
	robotgo.MouseClick()
	robotgo.MouseClick()
	robotgo.MouseClick()
	robotgo.SetMouseDelay(1000)
}


func clickingstuff() {

	clickthumbnail(183, 567)/*click shop now cta*/


  clickthumbnail(53, 267)
	screencap("test1.png")
  exitdetailview()

  robotgo.SetMouseDelay(1000)
  robotgo.MoveMouse(179, 547)
	//Println("pos:", x, y)


	clickthumbnail(140, 267)/*product in row 1 middle left */
	screencap("test2.png")
  exitdetailview()

	// robotgo.SetMouseDelay(1000)
	// robotgo.MoveMouse(179, 547)
	// Println("pos:", x, y)
	//
	// exitdetailview()



}


func main() {
	//for i := 1; i <= 10; i++ {
		clickingstuff()
	//}
}
