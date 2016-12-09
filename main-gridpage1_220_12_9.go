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

 	clickthumbnail(140, 267)/*product in row 1 middle left */
 	screencap("test2.png")
  exitdetailview()

	clickthumbnail(220, 270)/*product in row 1 middle right*/
	screencap("test3.png")
	exitdetailview()

	clickthumbnail(302, 267)
	screencap("test4.png")
	exitdetailview()

	clickthumbnail(50, 401)
	screencap("test5.png")
	exitdetailview()

	clickthumbnail(137, 401)
	screencap("test6.png")
	exitdetailview()

	clickthumbnail(220, 401)
	screencap("test7.png")
	exitdetailview()

	clickthumbnail(306, 401)
	screencap("test8.png")
	exitdetailview()

	clickthumbnail(52, 527)
	screencap("test9.png")
	exitdetailview()

	clickthumbnail(138, 527)
	screencap("test10.png")
	exitdetailview()

	clickthumbnail(220, 527)
	screencap("test11.png")
	exitdetailview()

	clickthumbnail(303, 527)
	screencap("test12.png")
	exitdetailview()


}


func main() {
	//for i := 1; i <= 10; i++ {
		clickingstuff()
	//}
}
