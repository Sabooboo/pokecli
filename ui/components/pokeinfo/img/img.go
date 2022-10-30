package img

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	imgascii "github.com/qeesung/image2ascii/convert"
	"image"
)

var showShiny = false

type Image struct {
	Common  common.Common
	resizer imgascii.ResizeHandler
	imgs    typdef.ShinyToggleable[image.Image]
	asciis  typdef.ShinyToggleable[string]
}

func New(info typdef.PokeResult, width, height int) Image {
	normalDat := info.Images.Normal.Img
	shinyDat := info.Images.Shiny.Img
	size := util.Min(width, height)

	i := Image{
		Common: common.Common{
			Width:  width,
			Height: height,
		},
		imgs: typdef.ShinyToggleable[image.Image]{
			Normal: normalDat,
			Shiny:  shinyDat,
		},
		asciis: typdef.ShinyToggleable[string]{
			Normal: util.ImageToASCII(normalDat, size, size, false),
			Shiny:  util.ImageToASCII(shinyDat, size, size, false),
		},
	}
	return i
}

func resizeImg(img image.Image, width, height int) string {
	size := util.Min(width, height)
	ascii := util.ImageToASCII(img, size, size, false)
	return ascii
}

func (i Image) SetSize(width, height int) common.Component {
	i.Common.SetSize(width, height)

	// SetSize can be called before img is ever loaded.
	if i.imgs.Normal != nil {
		i.asciis.Normal = resizeImg(i.imgs.Normal, width, height)
	}
	if i.imgs.Shiny != nil {
		i.asciis.Shiny = resizeImg(i.imgs.Shiny, width, height)
	}
	return i
}

func (i Image) Init() tea.Cmd {
	return nil
}

func (i Image) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "s": // Shiny view toggle
			showShiny = !showShiny
			return i, nil
		}

	}
	return i, nil
}

func (i Image) View() string {
	ascii := i.asciis.Normal
	if showShiny {
		ascii = i.asciis.Shiny
	}
	return ascii
}
