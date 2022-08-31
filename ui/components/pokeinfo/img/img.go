package img

import (
	"image"

	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	imgascii "github.com/qeesung/image2ascii/convert"
)

type Image struct {
	common common.Common
	img    image.Image
	ascii  string
}

func New(info typdef.PokeResult) Image {
	dat := info.Image

	i := Image{
		img:   dat,
		ascii: util.ImageToASCII(dat, &imgascii.DefaultOptions),
	}
	return i
}

func (i Image) SetSize(width, height int) common.Component {
	i.common.SetSize(width, height)
	// i.ascii = util.ImageToASCII(i.img, &imgascii.Options{
	// 	FixedWidth:  width,
	// 	FixedHeight: height,
	// })
	return i
}

func (i Image) Init() tea.Cmd {
	return nil
}

func (i Image) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i Image) View() string {
	return i.ascii
}
