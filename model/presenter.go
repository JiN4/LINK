package model

import (
	"time"

	"github.com/LINK/service"
)

type BrowsedHistories []service.ContentCard

type EventOnCompany struct {
	ID                 uint
	Title              string
	EventDescription   string
	EventThumbnail     string
	CreatedAt          time.Time
	CompanyID          uint
	CompanyThumbnail   string
	Name               string
	CompanyDescription string
	Url                string
	Address            string
}

type EventsOnCompany []EventOnCompany

func (eventsOnCompany EventsOnCompany) Cards() service.EventCards {
	eventCards := service.EventCards{}
	for _, e := range eventsOnCompany {
		eventCards = append(eventCards, service.EventCard{
			ID: e.ID,
			Company: service.CompanyCard{
				Id:          e.CompanyID,
				Thumbnail:   e.CompanyThumbnail,
				Name:        e.Name,
				Description: e.CompanyDescription,
				Url:         e.Url,
				Address:     e.Address,
			},
			Title:       e.Title,
			Description: e.EventDescription,
			Thumbnail:   e.EventThumbnail,
			CreatedAt:   e.CreatedAt,
		})
	}
	return eventCards
}

type ContentOnCompany struct {
	ID                 uint
	Author             string
	Title              string
	Description        string
	Thumbnail          string
	IsBookMarked       bool
	CompanyID          uint
	CreatedAt          time.Time
	CompanyThumbnail   string
	Name               string
	CompanyDescription string
	Url                string
	Address            string
}

type ContentsOnCompany []ContentOnCompany

func (contentsOnCompany ContentsOnCompany) Cards() service.ContentCards {
	contentCards := service.ContentCards{}
	for _, c := range contentsOnCompany {
		contentCards = append(contentCards, service.ContentCard{
			Id: c.ID,
			Company: service.CompanyCard{
				Id:          c.CompanyID,
				Thumbnail:   c.CompanyThumbnail,
				Name:        c.Name,
				Description: c.CompanyDescription,
				Url:         c.Url,
				Address:     c.Address,
			},
			Author:       c.Author,
			Title:        c.Title,
			Description:  c.Description,
			Thumbnail:    c.Thumbnail,
			IsBookMarked: c.IsBookMarked,
			CreatedAt:    c.CreatedAt,
		})
	}
	return contentCards
}
