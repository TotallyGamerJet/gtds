package main

import (
	"fmt"
	"golang.org/x/image/math/f32"
	"letsgo"
	mtl "letsgo/internal/metal"
	"log"
	"time"
	"unsafe"
)

func main() {
	if err := gtds.Run(run); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	window := gtds.CreateWindow(gtds.WindowConfig{
		Title:  "Window Title Here",
		Style:  gtds.Titled | gtds.Closable,
		Width:  640,
		Height: 480,
	})
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		return err
	}

	var windowSize = [2]int32{640, 480}
	var pos [2]float32

	lib, err := device.MakeLibrary(source, mtl.CompileOptions{})
	if err != nil {
		return err
	}
	vs, err := lib.MakeFunction("VertexShader")
	if err != nil {
		return err
	}
	fs, err := lib.MakeFunction("FragmentShader")
	if err != nil {
		return err
	}
	var rpld mtl.RenderPipelineDescriptor
	rpld.VertexFunction = vs
	rpld.FragmentFunction = fs
	rpld.ColorAttachments[0].PixelFormat = window.PixelFormat()
	rps, err := device.MakeRenderPipelineState(rpld)
	if err != nil {
		return err
	}

	// Create a vertex buffer.
	type Vertex struct {
		Position f32.Vec4
		Color    f32.Vec4
	}
	vertexData := [...]Vertex{
		{f32.Vec4{0, 0, 0, 1}, f32.Vec4{1, 0, 0, 1}},
		{f32.Vec4{300, 100, 0, 1}, f32.Vec4{0, 1, 0, 1}},
		{f32.Vec4{0, 100, 0, 1}, f32.Vec4{0, 0, 1, 1}},
	}
	vertexBuffer := device.MakeBuffer(unsafe.Pointer(&vertexData[0]), unsafe.Sizeof(vertexData), mtl.ResourceStorageModeManaged)

	cq := device.MakeCommandQueue()

	frame := startFPSCounter()

	for !window.ShouldClose() {
		gtds.PollEvents()
		// Create a drawable to render into.
		drawable, err := window.NextDrawable()
		if err != nil {
			return err
		}

		cb := cq.MakeCommandBuffer()

		// Encode all render commands.
		var rpd mtl.RenderPassDescriptor
		rpd.ColorAttachments[0].LoadAction = mtl.LoadActionClear
		rpd.ColorAttachments[0].StoreAction = mtl.StoreActionStore
		rpd.ColorAttachments[0].ClearColor = mtl.ClearColor{Red: 0.35, Green: 0.65, Blue: 0.85, Alpha: 1}
		rpd.ColorAttachments[0].Texture = drawable.Texture()
		rce := cb.MakeRenderCommandEncoder(rpd)
		rce.SetRenderPipelineState(rps)
		rce.SetVertexBuffer(vertexBuffer, 0, 0)
		rce.SetVertexBytes(unsafe.Pointer(&windowSize[0]), unsafe.Sizeof(windowSize), 1)
		rce.SetVertexBytes(unsafe.Pointer(&pos[0]), unsafe.Sizeof(pos), 2)
		rce.DrawPrimitives(mtl.PrimitiveTypeTriangle, 0, 3)
		rce.EndEncoding()

		cb.PresentDrawable(drawable)
		cb.Commit()

		frame <- struct{}{}
	}
	return nil
}

func startFPSCounter() chan struct{} {
	frame := make(chan struct{}, 4)
	go func() {
		second := time.Tick(time.Second)
		frames := 0
		for {
			select {
			case <-second:
				fmt.Println("fps:", frames)
				frames = 0
			case <-frame:
				frames++
			}
		}
	}()
	return frame
}
