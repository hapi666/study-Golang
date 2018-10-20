package main

import (
	"fmt"
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

func flush(minY, maxY int, f func(),flag int) {
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
				switch moveY {
				case 8:
					if flag == 1 {
						flush(8, 17, enterSongList, 1)
					}
					if flag == 2 { //在第二个页面点击第一个的enter键
						//results := crawl.TopList("云音乐新歌榜")
						//setListCell(results)
					}
					if flag == 3 { //在第三个页面点击第一个的enter键
						//Call HotComment function.
					}
				case 11:
					if flag == 1 { //在第一个页面点击第二个的enter键
						//flush(8, 11, enterSongHotComment, 1)
					}
					if flag == 2 { //在第二个页面点击第二个的enter键
						//call a SetCell function.
						//results := crawl.TopList("云音乐热歌榜")
						//setListCell(results)
					}
				case 14:
					if flag == 2 {
						//results := crawl.TopList("网易原创歌曲榜")
						//setListCell(results)
					}
					//...
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
	flag:=1
	flush(8, 11, initDraw,flag)
}
