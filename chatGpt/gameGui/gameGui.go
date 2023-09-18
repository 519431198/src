package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strconv"
)

type Game struct {
	i uint8
}

func (g *Game) Update() error {
	// 在这里添加你的游戏逻辑
	return nil
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.i++
	if g.i < 255 {
		screen.Fill(Hex2RGB("#0dceda", g.i))
	} else {
		g.i = 0
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello,ebiten!\nTPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480 // 设置图形界面的尺寸，你可以根据需求进行修改
}

func main() {
	ebiten.SetWindowSize(640, 480) // 设置窗口大小
	ebiten.SetWindowTitle("游戏")    // 设置窗口标题
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
