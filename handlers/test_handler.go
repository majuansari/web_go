package handlers

import (
	"github.com/labstack/echo/v4"
	"web/app"
)

type TestHandler struct {
	*app.App
}

// NewUserHandler ..
func NewTestHandler(server *app.App) TestHandler {
	return TestHandler{server}
}
func (t TestHandler) Test(c echo.Context) error {
	loopAndCreatePointer()
	return c.String(200, "Done")
}

func loopAndCreatePointer() {
	for i := 0; i < 100; i++ {
		_ = CreatePointer()
	}
}

type BigStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}

//go:noinline
func CreatePointer() *BigStruct {
	return &BigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}
