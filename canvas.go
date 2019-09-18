package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 마우스로 선 그리기
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     1200,
		Height:    1200,
		FrameRate: 60,
		Title:     "Canvas",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Black)
		ctx.SetLineWidth(10)
	})

	c.Draw(func(ctx *canvas.Context) {
	ctx.Push()
		if ctx.IsMouseDragged {
			ctx.DrawLine(
				ctx.Mouse.X,
				ctx.Mouse.Y,
				ctx.PMouse.X,
				ctx.PMouse.Y,
			)
			ctx.Stroke()
		}
		ctx.Pop()

		if ctx.IsKeyPressed(pixelgl.KeySpace) {
			ctx.Push()
			ctx.SetColor(colornames.White)
			ctx.Clear()
			ctx.SetColor(colornames.Black)
			ctx.Pop()
		}
	})
}
