package model


import(
  // "log"
  "github.com/LINK/service"
)

func GetHomeContent() (service.HomeContent) {
  banners, _ := GetAllContentCards()
  events, _ := GetAllEventCards()
  rankings, _ := GetSortedContentCards("browsed_num")
  newContents, _ := GetSortedContentCards("new")
  histories := service.ContentCards{}

  homeContent := service.HomeContent {
    Banners: banners,
    Events: events,
    Rankings: rankings,
    NewContents: newContents,
    Histories: histories,
  }
  return homeContent
}
