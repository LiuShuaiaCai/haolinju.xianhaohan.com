//go:build wireinject
// +build wireinject

package core

import (
	"github.com/google/wire"
	"haolinju.xianhaohan.com/internal/app/controllers"
	"haolinju.xianhaohan.com/internal/app/models"
	"haolinju.xianhaohan.com/internal/app/services"
)

//go:generate
func InitApp() (app *App, fc func(), err error) {
	panic(wire.Build(Init, models.Provider, services.Provider, controllers.Provider, NewApp))
}

//go:generate
func InitJob() (job *Job, fc func(), err error) {
	panic(wire.Build(Init, models.Provider, services.Provider, NewJob))
}
