package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"moviesDB/modele"
	"moviesDB/utils"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var moviesDBDir = "/media/veracrypt60/MoviesDB/"
var dbCollecsName = moviesDBDir + "Collections.sqlite"

var (
	sashPos1 float32 = 200
	sashPos2 float32 = 200
	sashPos3 float32 = 200
	sashPos4 float32 = 100
)
var font *giu.FontInfo
var moviesList []modele.MovieList
var movie modele.Movie
var texture *giu.Texture

type appUI struct {
	collec  string
	dbName  string
	dbType  string
	pathImg string
}

func buildRows() []*giu.TableRowWidget {
	moviesList = modele.GetMoviesList()

	rows := make([]*giu.TableRowWidget, len(moviesList))

	for i := range rows {
		seen := moviesList[i].Seen
		id := moviesList[i].ID
		title := moviesList[i].Title
		rows[i] = giu.TableRow(
			giu.Selectable(title).OnClick(func() {
				movie = modele.GetMoviesByID(id)
				picture := moviesDBDir + "Movies/" + movie.Picture
				fmt.Println(picture)
				//pictureImg, _ = utils.LoadImage(picture)
				pictureImg, _ := utils.LoadImage("gopher.jpg")
				img := giu.ImageToRgba(pictureImg)
				giu.NewTextureFromRgba(img, func(tex *giu.Texture) {
					texture = tex
				})
			}),
			giu.Checkbox("", &seen),
		)
	}
	return rows
}

func loop() {
	fontPushed := false
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			fontPushed = giu.PushFont(font)
		}),

		giu.SplitLayout(giu.DirectionHorizontal, sashPos1,
			giu.Layout{
				giu.Table().Freeze(0, 1).
					Columns(
						giu.TableColumn("Titre"),
						giu.TableColumn("Vu").Flags(giu.TableColumnFlagsWidthFixed).InnerWidthOrWeight(30),
					).
					Rows(buildRows()...).
					Flags(giu.TableFlags(giu.TableColumnFlagsDefaultSort | giu.TableColumnFlagsIsSorted | giu.TableColumnFlags(giu.TableFlagsSortMulti))),
			},
			giu.Layout{
				giu.Column(
					giu.Row(
						giu.Button(utils.I64toA(movie.ID)).Disabled(true),
						giu.Label(movie.Title),
						giu.AlignManually(giu.AlignRight, giu.Button("Télécharger"), 100, true),
					),
					giu.Row(
						giu.Custom(func() {
							canvas := giu.GetCanvas()
							if texture != nil {
								canvas.AddImage(texture, image.Pt(0, 0), image.Pt(600, 800))
							}

						}),
					),
				),
			},
		),

		giu.Custom(func() {
			if fontPushed {
				giu.PopFont()
			}
		}),
	)
}

func main() {
	myUI := &appUI{}
	if utils.Exists(dbCollecsName) {
		modele.OpenDBCollec(dbCollecsName)
		collection := modele.GetCollecByID(1)
		myUI.collec = collection.Name
		myUI.dbType = collection.Type
		myUI.dbName = moviesDBDir + collection.Name + ".sqlite"
		myUI.pathImg = moviesDBDir + collection.Name + "/"
	} else {
		modele.OpenDBCollec(dbCollecsName)
		modele.CreateCollecTable()
	}
	modele.OpenDB(myUI.dbName)

	wnd := giu.NewMasterWindow("MoviesDB", 1280, 1000, 0)
	imgui.StyleColorsLight()
	// imgui.PushStyleColor(imgui.StyleColorWindowBg, giu.ToVec4Color(color.White))
	// imgui.PushStyleColor(imgui.StyleColorButton, giu.ToVec4Color(color.White))
	// imgui.PushStyleColor(imgui.StyleColorButton, giu.ToVec4Color(color.White))
	// imgui.PushStyleColor(imgui.StyleColorText, giu.ToVec4Color(color.Black))
	//
	// //fmt.Println(giu.AddFont("Arial.ttf", 50))
	// giu.GetDefaultFonts()[0].SetSize(50)
	// //giu.Style().SetFont(giu.AddFont("Arial.ttf", 50))
	// //giu.SetDefaultFontSize(50)
	// giu.Style().SetFontSize(60).Build()

	//imgui.Style.ScaleAllSizes(imgui.CurrentStyle(), 2)
	font = giu.AddFont("Arial.ttf", 20)
	fmt.Println(font)

	pictureImg, _ := utils.LoadImage("/media/veracrypt60/MoviesDStar/Movies/_Chef_0.jpg")
	img := giu.ImageToRgba(pictureImg)
	//img, _ := g.LoadImage("gopher.png")
	giu.EnqueueNewTextureFromRgba(img, func(tex *giu.Texture) {
		texture = tex
	})

	wnd.Run(loop)
}
