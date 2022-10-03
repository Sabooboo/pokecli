package img

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	imgascii "github.com/qeesung/image2ascii/convert"
	"image"
)

// var options = imgascii.Options{
// 	FixedWidth:  -1,
// 	FixedHeight: -1,
// }

type Image struct {
	Common  common.Common
	resizer imgascii.ResizeHandler
	img     image.Image
	ascii   string
}

func New(info typdef.PokeResult, width, height int) Image {
	dat := info.Image
	size := util.Min(width, height)

	i := Image{
		Common: common.Common{
			Width:  width,
			Height: height,
		},
		img:   dat,
		ascii: util.ImageToASCII(dat, size, size, false),
	}
	return i
}

func (i Image) SetSize(width, height int) common.Component {
	i.Common.SetSize(width, height)

	// SetSize can be called before img is ever loaded.
	if i.img != nil {
		size := util.Min(width, height)
		i.ascii = util.ImageToASCII(i.img, size, size, false)
	}
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
