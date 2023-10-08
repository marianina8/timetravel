package speedometer

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

func New() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	speedDisplay, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}

	mphtxt, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	if err := mphtxt.Write([]*segmentdisplay.TextChunk{
		segmentdisplay.NewChunk("MPH"),
	}); err != nil {
		panic(err)
	}

	speedDonut, err := donut.New(donut.CellOpts(cell.FgColor(cell.ColorNumber(33))))
	if err != nil {
		panic(err)
	}

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("Speedometer - Press Q to quit"),
		container.SplitVertical(
			container.Left(
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(speedDisplay),
					),
					container.Right(
						container.PlaceWidget(mphtxt),
					),
					container.SplitPercent(70),
				),
			),
			container.Right(
				container.PlaceWidget(speedDonut),
			),
			container.SplitPercent(75),
		),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	go func() {
		for speed := 0.0; speed <= 88.8; speed += 0.1 {
			if err := speedDisplay.Write([]*segmentdisplay.TextChunk{
				segmentdisplay.NewChunk(fmt.Sprintf("%.1f", speed)),
			}); err != nil {
				panic(err)
			}

			percentage := math.Ceil((speed / 88.8) * 100)
			speedDonut.Percent(int(percentage))

			time.Sleep(5 * time.Millisecond)
		}
		cancel()
	}()

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}
}
