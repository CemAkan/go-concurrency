package crawler

import "exercise3/models"

type PageResult struct {
	URL      string
	Elements []models.Element
}
