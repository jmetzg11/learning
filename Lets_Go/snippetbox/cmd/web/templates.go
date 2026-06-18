package main 

import "snippetbox.jmetzg11/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}