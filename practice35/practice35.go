package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var (
	moveX int
	moveY int
	oneLineWord =[]rune{'网', '易','云','音','乐','欢','迎','您'}
	twoLineWord =[]rune{'1','.','获','取','榜','单','.'}
	threeLineWord =[]rune{'2','.','获','取','歌','曲','热','评','.'}
)


func initDraw(moveX ,moveY int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j:=0
	for i:=30;i<46 ;i+=2  {
		termbox.SetCell(i,5,oneLineWord[j],termbox.ColorDefault,termbox.ColorBlue)
		j++
	}

	termbox.SetCell(moveX,moveY,'>',termbox.ColorDefault,termbox.ColorRed)

	k:=0
	for i:=32;i<46 ;i+=2  {
		termbox.SetCell(i,8,twoLineWord[k],termbox.ColorDefault,termbox.ColorBlack)
		k++
	}
	l:=0
	for i:=32;i<50 ;i+=2  {
		termbox.SetCell(i,11,threeLineWord[l],termbox.ColorDefault,termbox.ColorBlack)
		l++
	}

	termbox.Flush()
}

func drawSongLists(){

}

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	moveX=30
	moveY=8
	initDraw(30,8)
loop:
	for ev := range eventQueue {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp :
				moveY-=3
			case termbox.KeyArrowDown:
				moveY+=3
			case termbox.KeyEnter:
				//....
			case termbox.KeyEsc:
				break loop
			}
		}
		if moveY>11 {
			moveY=11
		}else if moveY<8 {
			moveY=8
		}
		initDraw(moveX, moveY)
	}
}
