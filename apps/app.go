package apps

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"wails_hashcat/global"
	"wails_hashcat/model"
)

var (
	selectFilePath string
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called at application Startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// DomReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

// BeforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// Shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
func (a *App) ChooseFile() string {
	var err error
	selectFilePath, err = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return ""
	}
	global.SugarLogger.Info(selectFilePath)
	fileName := filepath.Base(selectFilePath)
	return fileName
}

func (a *App) Confirm(fileType string) {
	fileName := filepath.Base(selectFilePath)
	db := global.DB
	var sourceFile = model.SourceHashFile{
		FileName: fileName,
		FilePath: selectFilePath,
		FileType: fileType,
	}
	db.Save(&sourceFile)
}
