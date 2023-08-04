package snake

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func LoadStartMenu(g *Game) *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(backgroundColor)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: ScreenHeight / 2.5}),

			widget.RowLayoutOpts.Spacing(20),
		),
		),
	)
	// startButtonContainer := newContainer()
	startButton := newButton("Start", func(*widget.ButtonClickedEventArgs) { g.reset(); g.nextState(GameState) })
	// startButtonContainer.AddChild(startButton)
	// exitButtonContainer := newContainer()
	exitButton := newButton("Exit", func(*widget.ButtonClickedEventArgs) { g.exit() })
	// exitButtonContainer.AddChild(exitButton)
	rootContainer.AddChild(startButton)
	rootContainer.AddChild(exitButton)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}
	return ui
}

func LoadPausedMenu(g *Game) *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(backgroundColor)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: ScreenHeight / 2.5}),

			widget.RowLayoutOpts.Spacing(20),
		),
		),
	)
	// startButtonContainer := newContainer()
	resumeButton := newButton("Resume", func(*widget.ButtonClickedEventArgs) { g.nextState(GameState) })
	// startButtonContainer.AddChild(startButton)
	// exitButtonContainer := newContainer()
	restartButton := newButton("Restart", func(*widget.ButtonClickedEventArgs) { g.reset(); g.nextState(GameState) })
	// exitButtonContainer.AddChild(exitButton)
	rootContainer.AddChild(resumeButton)
	rootContainer.AddChild(restartButton)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}
	return ui
}

func LoadEndMenu(g *Game) *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(backgroundColor)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: ScreenHeight / 2.5}),

			widget.RowLayoutOpts.Spacing(20),
		),
		),
	)
	scoreButton := newButton(fmt.Sprintf("Your Score: %v", g.score), func(*widget.ButtonClickedEventArgs) {})
	scoreImage := loadButtonImage()
	scoreImage.Idle = image.NewNineSliceColor(backgroundColor)
	scoreImage.Hover = image.NewNineSliceColor(backgroundColor)
	scoreImage.Pressed = image.NewNineSliceColor(backgroundColor)
	scoreButton.Configure(
		widget.ButtonOpts.Image(
			scoreImage,
		),
	)
	restartButton := newButton("Restart", func(*widget.ButtonClickedEventArgs) { g.reset(); g.nextState(GameState) })
	backButton := newButton("Back", func(*widget.ButtonClickedEventArgs) { g.reset() })
	rootContainer.AddChild(scoreButton)
	rootContainer.AddChild(restartButton)
	rootContainer.AddChild(backButton)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}
	return ui
}

func newButton(text string, onClick func(*widget.ButtonClickedEventArgs)) *widget.Button {
	buttonImage := loadButtonImage()
	face, _ := loadFont(30)
	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
				MaxWidth: ScreenWidth / 4,
			}),
		),

		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.Text(text, face, &widget.ButtonTextColor{
			Idle: textColor,
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),

		widget.ButtonOpts.ClickedHandler(onClick),
	)
	return button
}

func loadButtonImage() *widget.ButtonImage {
	idle := image.NewNineSliceColor(buttonColor)
	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})
	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
