package dashboard

import (
	"context"
	"fmt"
	"time"

	"github.com/marianina8/timetravel/utils"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

func New(travelTime string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	destinationSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	presentSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	departedSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}

	if err := destinationSD.Write([]*segmentdisplay.TextChunk{
		segmentdisplay.NewChunk(travelTime),
	}); err != nil {
		panic(err)
	}

	if err := presentSD.Write([]*segmentdisplay.TextChunk{
		segmentdisplay.NewChunk(utils.FormatDestination(time.Now())),
	}); err != nil {
		panic(err)
	}

	if err := departedSD.Write([]*segmentdisplay.TextChunk{
		segmentdisplay.NewChunk(utils.FormatDestination(time.Now().Add(-4 * time.Minute))),
	}); err != nil {
		panic(err)
	}

	travelBtn, err := button.New("Blast Off!",
		func() error {
			cancel()
			return nil
		},
		button.FillColor(cell.ColorGreen),
	)
	if err != nil {
		panic(err)
	}

	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("DASHBOARD - Press Q to quit"),
		container.SplitHorizontal(
			container.Top(
				container.SplitHorizontal(
					container.Top(
						container.Border(linestyle.Light),
						container.BorderTitle("DESTINATION TIME"),
						container.BorderTitleAlignCenter(),
						container.PlaceWidget(destinationSD),
					),
					container.Bottom(
						container.SplitHorizontal(
							container.Top(
								container.Border(linestyle.Light),
								container.BorderTitle("PRESENT TIME"),
								container.BorderTitleAlignCenter(),
								container.PlaceWidget(presentSD),
							),
							container.Bottom(
								container.Border(linestyle.Light),
								container.BorderTitle("LAST DEPARTED TIME"),
								container.BorderTitleAlignCenter(),
								container.PlaceWidget(departedSD),
							),
							container.SplitPercent(50),
						)),
					container.SplitPercent(33),
				),
			),
			container.Bottom(
				container.PlaceWidget(travelBtn),
			),
			container.SplitPercent(90),
		),
	)
	if err != nil {
		panic(err)
	}

	quit := false
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			quit = true
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}

	return !quit
}
