package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/AlissonBarbosa/shylockgo-scraper/src/common"
	"github.com/AlissonBarbosa/shylockgo-scraper/src/controllers"
	"github.com/AlissonBarbosa/shylockgo-scraper/src/models"
)

func main() {
  filelog, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    slog.Error("Error creating log file")
    return
  }
  defer filelog.Close()

	l := slog.New(slog.NewJSONHandler(filelog, nil))
	slog.SetDefault(l)

	models.ConnectDatabase()

	provider, err := common.GetProvider()
	if err != nil {
		slog.Error("Error getting provider:", err)
		return
	}

	serversFlag := flag.Bool("servers", false, "Execute servers functions")
	projectsFlag := flag.Bool("projects", false, "Execute projects functions")
	flavorsFlag := flag.Bool("flavors", false, "Execute flavors functions")
	usersFlag := flag.Bool("users", false, "Execute users functions")
	allFlag := flag.Bool("all", false, "Execute all functions")
	flag.Parse()

	if *serversFlag {
		slog.Info("Executing servers functions")
		controllers.SaveServersDesc(provider)
		controllers.SaveServersSpec(provider)
		controllers.SaveServersUsage(provider)
		controllers.SaveServersOwnership(provider)
	}

	if *projectsFlag {
		slog.Info("Executing projects functions")
		controllers.SaveProjectsDesc(provider)
		controllers.SaveProjectQuota(provider)
		controllers.SaveProjectUsage(provider)
	}

	if *flavorsFlag {
		slog.Info("Executing flavors functions")
		controllers.SaveFlavorDesc(provider)
		controllers.SaveFlavorSpec(provider)
	}

	if *usersFlag {
		slog.Info("Executing users functions")
		controllers.SaveUserDesc(provider)
		controllers.SaveUserProjects(provider)
	}

  if *allFlag {
		slog.Info("Executing servers functions")
		controllers.SaveServersDesc(provider)
		controllers.SaveServersSpec(provider)
		controllers.SaveServersUsage(provider)
		controllers.SaveServersOwnership(provider)
		slog.Info("Executing projects functions")
		controllers.SaveProjectsDesc(provider)
		controllers.SaveProjectQuota(provider)
		controllers.SaveProjectUsage(provider)
		slog.Info("Executing flavors functions")
		controllers.SaveFlavorDesc(provider)
		controllers.SaveFlavorSpec(provider)
		slog.Info("Executing users functions")
		controllers.SaveUserDesc(provider)
		controllers.SaveUserProjects(provider)
	}
}
