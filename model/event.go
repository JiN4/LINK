package model

import (
	"github.com/LINK/service"
)

// Create Eventを作成する
func CreateEvent(event Event) (Event, error) {
	err := db.Create(&event).Error
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func GetAllEvents() (Events, error) {
	events := Events{}
	err := db.Find(&events).Error
	return events, err
}

// EventをIDで指定して取得する
func GetEventById(eventId uint) (service.Event, error) {
	eventCard := service.EventCard{}
	err := db.Table("events").Find(&eventCard, eventId).Error
	companyCard, _ := GetCompanyCardById(eventCard.CompanyID)
	eventCard.Company = companyCard
	contentCards, _ := GetContentCardsFilteredByEventId(eventCard.ID)

	return service.Event{EventCard: eventCard, Contents: contentCards}, err
}

// Eventをすべてカード形式で取得
func GetAllEventCards() (service.EventCards, error) {

	// JOINして速くした
	eventsOnCompany := EventsOnCompany{}
	sql := "SELECT e.id, e.title, e.description AS event_description, e.thumbnail AS event_thumbnail, e.created_at, c.id AS company_id, c.thumbnail AS company_thumbnail, c.name, c.description AS company_description, c.url, c.address FROM events AS e INNER JOIN companies AS c ON e.company_id = c.id"
	err := db.Raw(sql).Scan(&eventsOnCompany).Error

	return eventsOnCompany.Cards(), err
	// return eventCards, err
}

func (events Events) Cards() service.EventCards {
	var cards service.EventCards
	for _, event := range events {
		company, _ := GetCompanyById(event.CompanyID)
		cards = append(cards, service.EventCard{
			ID:          event.ID,
			Company:     company.Card(),
			Title:       event.Title,
			Description: event.Description,
			Thumbnail:   event.Thumbnail,
			CompanyID:   event.CompanyID,
			CreatedAt:   event.CreatedAt,
		})
	}
	return cards
}
