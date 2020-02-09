package main

import (
	"fmt"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	window *widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	// Create the main window struct
	mw := new(MainWindow)
	// Create the main Qt main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("spotify-qt")
	window.SetWindowIcon(gui.QIcon_FromTheme("spotify"))
	window.Resize2(1280, 720)
	window.SetCentralWidget(mw.NewCentralWidget())
	// 1=left, 2=right, 4=top, 8=bottom
	window.AddToolBar(0x4, mw.NewToolBar())
	// Set and return final window
	mw.window = window
	return mw
}

func LayoutToWidget(layout widgets.QLayout_ITF) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	return widget
}

func CreateGroupBox(widget ...widgets.QWidget_ITF) *widgets.QGroupBox {
	group := widgets.NewQGroupBox(nil)
	layout := widgets.NewQVBoxLayout()
	for _, w := range widget {
		layout.AddWidget(w, 0, 0)
	}
	group.SetLayout(layout)
	return group
}

func NewMenuAction(icon, text string, shortcut gui.QKeySequence__StandardKey) *widgets.QAction {
	action := widgets.NewQAction3(gui.QIcon_FromTheme(icon), text, nil)
	if shortcut != 0 {
		action.SetShortcut(gui.NewQKeySequence5(shortcut))
	}
	return action
}

func (mw *MainWindow) NewCentralWidget() widgets.QWidget_ITF {
	container := widgets.NewQSplitter(nil)
	// Sidebar with playlists etc.
	sidebar := widgets.NewQVBoxLayout()
	libraryList := widgets.NewQListWidget(nil)
	playlistsList := widgets.NewQListWidget(nil)
	// Library
	libraryList.AddItems([]string{
		"Made For You", "Recently Played", "Liked Songs", "Albums", "Artists",
	})
	libraryList.ConnectItemPressed(func(item *widgets.QListWidgetItem) {
		if item != nil {
			playlistsList.SetCurrentRow(-1)
		}
	})
	library := CreateGroupBox(libraryList)
	library.SetTitle("Library")
	sidebar.AddWidget(library, 0, 0)
	// Playlists
	playlistsList.AddItems([]string{
		"p01", "p02", "p03", "p04", "p05", "p06", "p07", "p08", "p09",
	})
	playlistsList.SetCurrentRow(0)
	playlistsList.ConnectItemPressed(func(item *widgets.QListWidgetItem) {
		if item != nil {
			libraryList.SetCurrentRow(-1)
		}
	})
	playlists := CreateGroupBox(playlistsList)
	playlists.SetTitle("Playlists")
	sidebar.AddWidget(playlists, 1, 0)
	// Now playing song
	nowPlaying := widgets.NewQHBoxLayout()
	nowPlayingArt := widgets.NewQLabel(nil, 0)
	nowPlayingArt.SetFixedSize2(64, 64)
	nowPlayingArt.SetPixmap(
		gui.QIcon_FromTheme("media-optical-audio").Pixmap(nowPlayingArt.Size(), 0, 0))
	nowPlaying.AddWidget(nowPlayingArt, 0, 0)
	nowPlayingText := widgets.NewQLabel2("No music playing", nil, 0)
	nowPlaying.AddWidget(nowPlayingText, 1, 0)
	sidebar.AddWidget(LayoutToWidget(nowPlaying), 0, 0)

	sidebarWidget := LayoutToWidget(sidebar)
	sidebarWidget.SetMaximumWidth(250)
	container.AddWidget(sidebarWidget)
	// Table with songs
	songs := widgets.NewQTreeWidget(nil)
	songs.SetColumnCount(4)
	songs.SetEditTriggers(0)
	songs.SetSelectionBehavior(1)
	songs.SetSortingEnabled(true)
	songs.SetRootIsDecorated(false)
	songs.SetAllColumnsShowFocus(true)
	songs.SortByColumn(0, 0)
	headers := []string{
		"Title", "Artist", "Album", "Length",
	}
	songs.SetHeaderLabels(headers)
	for i := range headers {
		songs.SetColumnWidth(i, 200)
	}
	for i := 0; i < 20; i++ {
		item := widgets.NewQTreeWidgetItem2([]string{
			fmt.Sprintf("song%02v-title",  i),
			fmt.Sprintf("song%02v-artist", i),
			fmt.Sprintf("song%02v-album",  i),
			fmt.Sprintf("song%02v-length", i),
		}, 0)
		if i == 0 {
			item.SetIcon(0, gui.QIcon_FromTheme("media-playback-pause"))
		}
		songs.InsertTopLevelItem(i, item)

	}
	container.AddWidget(songs)
	return container
}

func (mw *MainWindow) NewToolBar() widgets.QToolBar_ITF {
	toolBar := widgets.NewQToolBar2(nil)
	toolBar.SetMovable(false)
	// Menu
	menu := widgets.NewQToolButton(nil)
	menu.SetText("Menu")
	menu.SetIcon(gui.QIcon_FromTheme("application-menu"))
	menu.SetPopupMode(2)
	menu.SetMenu(mw.NewMenu())
	toolBar.AddWidget(menu)
	toolBar.AddSeparator()
	// Media controls
	toolBar.AddAction2(gui.QIcon_FromTheme("media-skip-backward"),  "Previous")
	toolBar.AddAction2(gui.QIcon_FromTheme("media-playback-start"), "Play")
	toolBar.AddAction2(gui.QIcon_FromTheme("media-playback-stop"),  "Stop")
	toolBar.AddAction2(gui.QIcon_FromTheme("media-skip-forward"),   "Next")
	// Progress
	progress := widgets.NewQSlider(nil)
	progress.SetOrientation(0x1)
	toolBar.AddSeparator()
	toolBar.AddWidget(progress)
	toolBar.AddSeparator()
	toolBar.AddWidget(widgets.NewQLabel2("0:00/0:00", nil, 0))
	toolBar.AddSeparator()
	// Repeat and shuffle toggles
	toolBar.AddAction2(gui.QIcon_FromTheme("media-playlist-repeat"),   "Repeat").SetCheckable(true)
	toolBar.AddAction2(gui.QIcon_FromTheme("media-playlist-shuffle"),  "Shuffle").SetCheckable(true)
	// Volume slider
	volume := widgets.NewQSlider(nil)
	volume.SetOrientation(0x1)
	volume.SetMaximumWidth(100)
	volume.SetMinimum(0)
	volume.SetMaximum(20)
	volume.SetValue(20)
	toolBar.AddWidget(volume)
	// Return final toolbar
	return toolBar
}

func (mw *MainWindow) NewMenu() widgets.QMenu_ITF {
	// Create root menu
	menu := widgets.NewQMenu(nil)
	// About submenu
	aboutMenu := widgets.NewQMenu(nil)
	aboutMenu.SetTitle("About")
	aboutMenu.SetIcon(gui.QIcon_FromTheme("help-about"))
	aboutMenu.AddActions([]*widgets.QAction{
		NewMenuAction("spotify", "About spotify-qt", 0),
		NewMenuAction("qt",      "About Qt",         0),
	})
	// Check for updates / GC subsection
	aboutMenu.AddSeparator()
	aboutMenu.AddActions([]*widgets.QAction{
		NewMenuAction("download",  "Check for updates", 0),
		NewMenuAction("run-clean", "Run GC",            0),
	})
	menu.AddMenu(aboutMenu)
	// Refresh and settings subsection
	menu.AddActions([]*widgets.QAction{
		NewMenuAction("view-refresh", "Refresh",     gui.QKeySequence__Refresh),
		NewMenuAction("settings",     "Settings...", gui.QKeySequence__Preferences),
	})
	// Log out and quit subsection
	menu.AddSeparator()
	menu.AddActions([]*widgets.QAction{
		NewMenuAction("im-user-away",     "Log out", 0),
		NewMenuAction("application-exit", "Quit", gui.QKeySequence__Quit),
	})
	// Return final menu
	return menu
}