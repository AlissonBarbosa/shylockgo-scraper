package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/AlissonBarbosa/shylockgo-scraper/common"
	"github.com/AlissonBarbosa/shylockgo-scraper/controllers"
	"github.com/AlissonBarbosa/shylockgo-scraper/models"
)

func main()  {
  l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
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
}
