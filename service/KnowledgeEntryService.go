package service

import (
	"kdb/model"
	"kdb/repository"
)

var knowledgeEntryRepository repository.KnowledgeEntryRepository

func init() {
	knowledgeEntryRepository = knowledgeEntryRepository.InitKnowledgeEntryRepository()
}

func GetAllKnowledgeEntries() []model.KnowledgeEntry {
	return knowledgeEntryRepository.GetAllKnowledgeEntries()
}

func GetKnowledgeEntryById(id int) *model.KnowledgeEntry {
	knowledgeEntryRepository.GetKnowledgeEntryById(id)
}

func GetKnowledgeEntryByFullName(title string) *model.KnowledgeEntry{
	knowledgeEntryRepository.GetKnowledgeEntryByTitle(title)
}

func CreateKnowledgeEntry(s *model.KnowledgeEntry) *model.KnowledgeEntry {
	knowledgeEntryRepository.AddKnowledgeEntry(s)
}

func ModifyKnowledgeEntry(id int, s *model.KnowledgeEntry) *model.KnowledgeEntry {
	knowledgeEntryRepository.ModifyKnowledgeEntry(id, s)
}

func DeleteKnowledgeEntry(id int) {
	knowledgeEntryRepository.DeleteKnowledgeEntry(id)
}
