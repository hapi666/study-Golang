package main

import (
	"fmt"
	"reflect"

	"github.com/nsf/termbox-go"
)

var (
	eventQueue = make(chan termbox.Event)
	//moveX int
	moveY                   int
	oneLineWord             = []rune{'网', '易', '云', '音', '乐', '欢', '迎', '您'}
	twoLineWord             = []rune{'1', '.', '获', '取', '榜', '单', '.'}
	threeLineWord           = []rune{'2', '.', '获', '取', '歌', '曲', '热', '评', '.'}
	NewSongList             = []rune{'1', '.', '云', '音', '乐', '新', '歌', '榜', '.'}
	HotSongList             = []rune{'2', '.', '云', '音', '乐', '热', '歌', '榜', '.'}
	NeteaseOriginalSongList = []rune{'3', '.', '网', '易', '原', '创', '歌', '曲', '榜', '.'}
	CloudMusicSoared        = []rune{'4', '.', '云', '音', '乐', '飙', '升', '榜', '.'}
)

func initDraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 30+2*len(oneLineWord); i += 2 {
		termbox.SetCell(i, 5, oneLineWord[j], termbox.ColorDefault, termbox.ColorBlue)
		j++
	}

	termbox.SetCell(28, moveY, '>', termbox.ColorDefault, termbox.ColorRed)

	k := 0
	for i := 30; i < 30+2*len(twoLineWord); i += 2 {
		termbox.SetCell(i, 8, twoLineWord[k], termbox.ColorDefault, termbox.ColorBlack)
		k++
	}
	l := 0
	for i := 30; i < 30+2*len(threeLineWord); i += 2 {
		termbox.SetCell(i, 11, threeLineWord[l], termbox.ColorDefault, termbox.ColorBlack)
		l++
	}

	termbox.Flush()
}

func enterSongList() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 46; i += 2 {
		termbox.SetCell(i, 5, oneLineWord[j], termbox.ColorDefault, termbox.ColorBlue)
		j++
	}
	termbox.SetCell(28, moveY, '>', termbox.ColorDefault, termbox.ColorRed)
	k := 0
	for i := 30; i < 30+2*len(NewSongList); i += 2 {
		termbox.SetCell(i, 8, NewSongList[k], termbox.ColorDefault, termbox.ColorBlack)
		k++
	}
	l := 0
	for i := 30; i < 30+2*len(HotSongList); i += 2 {
		termbox.SetCell(i, 11, HotSongList[l], termbox.ColorDefault, termbox.ColorBlack)
		l++
	}
	n := 0
	for i := 30; i < 30+2*len(NeteaseOriginalSongList); i += 2 {
		termbox.SetCell(i, 14, NeteaseOriginalSongList[n], termbox.ColorDefault, termbox.ColorBlack)
		n++
	}
	p := 0
	for i := 30; i < 30+2*len(CloudMusicSoared); i += 2 {
		termbox.SetCell(i, 17, CloudMusicSoared[p], termbox.ColorDefault, termbox.ColorBlack)
		p++
	}
	termbox.Flush()
}

func enterSongHotComment() {
	//termbox.KeyInsert
}

func flush(minY, maxY int, f func()) {
	//moveX = 28
	moveY = 8
	f()
loop:
	for ev := range eventQueue {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				moveY -= 3
			case termbox.KeyArrowDown:
				moveY += 3
			case termbox.KeyEnter:
				//....
				var b bool
				switch moveY {
				case 8:
					switch b {
					case reflect.DeepEqual(f, initDraw):
						flush(8, 17, enterSongList)
					case reflect.DeepEqual(f, enterSongList):
						//call a SetCell function.
					}

				case 11:
					//flush(8,,)
					switch b {
					case reflect.DeepEqual(f, initDraw):
						flush(8, 11, enterSongHotComment)
					case reflect.DeepEqual(f, enterSongHotComment):
						//call a SetCell function.
					}
				}
				break loop
			case termbox.KeyEsc:
				break loop
			}
		}
		if moveY > maxY {
			moveY = maxY
		} else if moveY < minY {
			moveY = minY
		}
		f()
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer termbox.Close()

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
	flush(8, 11, initDraw)
}
