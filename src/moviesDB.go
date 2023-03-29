package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"moviesDB/modele"
	"moviesDB/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

const appDevPath = "/media/jluc20mx/WorkSpace/GO/moviesDB_giu/"

var (
	appRep        string
	err           error
	moviesDBDir   string
	dbCollecsName string
	moviesList    []modele.MovieList
	movie         modele.Movie

	explorerWidth float32 = 200
	font          *giu.FontInfo
	bigFont       *giu.FontInfo
	texture       *giu.Texture
	idxSelectable int64
)

type appUI struct {
	collec  string
	dbName  string
	dbType  string
	pathImg string
}

func (ui *appUI) buildRows() []*giu.TableRowWidget {
	moviesList = modele.GetMoviesList()

	rows := make([]*giu.TableRowWidget, len(moviesList))

	for i := range rows {
		seen := moviesList[i].Seen
		id := moviesList[i].ID
		title := moviesList[i].Title
		titleSelectable := giu.Selectable(title)
		titleSelectable.OnClick(func() {
			idxSelectable = id
			movie = modele.GetMoviesByID(id)
			pictureImg, _ := utils.LoadImage(filepath.Join(ui.pathImg, movie.Picture))
			img := giu.ImageToRgba(pictureImg)
			giu.NewTextureFromRgba(img, func(tex *giu.Texture) {
				texture = tex
			})
		})
		if id == idxSelectable {
			titleSelectable.Selected(true)
		}
		rows[i] = giu.TableRow(
			titleSelectable,
			giu.Checkbox("", &seen),
		)
	}
	if idxSelectable == 0 {
		id := moviesList[0].ID
		idxSelectable = id
		movie = modele.GetMoviesByID(id)
		pictureImg, _ := utils.LoadImage(filepath.Join(ui.pathImg, movie.Picture))
		img := giu.ImageToRgba(pictureImg)
		giu.NewTextureFromRgba(img, func(tex *giu.Texture) {
			texture = tex
		})
	}
	return rows
}

func (ui *appUI) loop() {
	fontPushed := false
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			fontPushed = giu.PushFont(bigFont)
		}),

		giu.SplitLayout(giu.DirectionHorizontal, explorerWidth,
			giu.Layout{
				giu.Style().
					SetFont(font).
					To(
						giu.Table().Freeze(0, 1).
							Columns(
								giu.TableColumn("Titre"),
								giu.TableColumn("Vu").Flags(giu.TableColumnFlagsWidthFixed).InnerWidthOrWeight(30),
							).
							Rows(ui.buildRows()...).
							Flags(giu.TableFlags(giu.TableColumnFlagsDefaultSort | giu.TableColumnFlagsIsSorted | giu.TableColumnFlags(giu.TableFlagsSortMulti))),
					),
			},
			giu.Layout{
				giu.Column(
					giu.Row(
						giu.Style().
							SetColor(giu.StyleColorButton, color.White).
							SetStyle(giu.StyleVarFramePadding, 5, 5).
							To(giu.Button(utils.I64toA(movie.ID)).Disabled(true)),
						giu.Style().
							SetStyle(giu.StyleVarFramePadding, 5, 5).
							To(giu.InputText(&movie.Title)),
						giu.AlignManually(
							giu.AlignRight,
							giu.Style().
								SetStyle(giu.StyleVarFramePadding, 5, 5).
								To(giu.Button("Télécharger")),
							135,
							true,
						),
					),
					giu.TabBar().TabItems(
						giu.TabItem("Fiche").Layout(
							giu.Style().
								SetFontSize(20).
								To(
									giu.Row(
										giu.Label("Titre original : "),
										giu.InputText(&movie.OriginalTitle),
									),
									giu.Row(
										giu.Image(texture).Size(600, 800),
										giu.Column(
											giu.Table().
												Rows(
													giu.TableRow(
														giu.Label(""),
														giu.Label(""),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Date de sortie : "),
															135,
															true,
														),
														giu.InputText(&movie.DateSortie).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Réalisateur : "),
															115,
															true,
														),
														giu.InputText(&movie.Directors).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Durée : "),
															73,
															true,
														),
														giu.InputInt(&movie.Duration).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Age minimum : "),
															135,
															true,
														),
														giu.InputText(&movie.AgeMini).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Nationalités : "),
															120,
															true,
														),
														giu.InputText(&movie.Countries).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Genres : "),
															85,
															true,
														),
														giu.InputText(&movie.Genres).Size(200),
													),
													giu.TableRow(
														giu.AlignManually(
															giu.AlignRight,
															giu.Label("Acteurs : "),
															86,
															true,
														),
														giu.InputText(&movie.Actors).Size(200),
													),
												).
												Flags(giu.TableFlagsNoBordersInBody),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Label(""),
											giu.Button("Changer"),
										),
									),
									giu.Label("Synopsis :"),
									giu.InputTextMultiline(&movie.Synopsis).
										Size(-1, -1),
								),
						),

						giu.TabItem("Détails").Layout(),

						giu.TabItem("Infos").Layout(),
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
	appRep, err = getAppPath()
	if err != nil {
		panic(fmt.Sprint("  init:getAppPath : ", err))
	}

	moviesDBDir = filepath.Join(appRep, "Collections")
	dbCollecsName = filepath.Join(moviesDBDir, "Collections.sqlite")
	fmt.Println(dbCollecsName)

	myUI := &appUI{}
	if utils.Exists(dbCollecsName) {
		modele.OpenDBCollec(dbCollecsName)
		collection := modele.GetCollecByID(1)
		myUI.collec = collection.Name
		myUI.dbType = collection.Type
		myUI.dbName = filepath.Join(moviesDBDir, collection.Name+".sqlite")
		myUI.pathImg = filepath.Join(moviesDBDir, collection.Name)
	} else {
		modele.OpenDBCollec(dbCollecsName)
		modele.CreateCollecTable()
	}
	modele.OpenDB(myUI.dbName)

	wnd := giu.NewMasterWindow("MoviesDB", 1280, 1150, 0)
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
	bigFont = giu.AddFont("Arial.ttf", 25)
	fmt.Println(font)

	wnd.Run(myUI.loop)
}

// récupère le répertoire de l'application
func getAppPath() (string, error) {
	appPath, err := os.Executable()
	if err == nil {
		appPath = filepath.Dir(appPath)
		fmt.Println("appPath - os.Executable()           : " + appPath)

		appPath, err = filepath.EvalSymlinks(appPath)
		if err == nil {
			fmt.Println("appPath - filepath.EvalSymlinks(...): " + appPath)

			parts := strings.Split(appPath, string(os.PathSeparator))
			fmt.Println(parts)
			fmt.Println(filepath.Join(parts...))
			if parts[1] == "tmp" && strings.Contains(parts[2], "go-build") {
				fmt.Println("Session de DEV.")
				appPath = appDevPath
			}
			if parts[len(parts)-1] == "__debug_bin" {
				fmt.Println("Session de DEBUG")
				appPath = appDevPath
			}
			if strings.Contains(appPath, appDevPath) {
				fmt.Println("Session de DEBUG")
				appPath = appDevPath
			}
		}
	}
	return appPath, err
}
