package repository

import "kdb/model"

type KnowledgeEntryRepository struct {

}

func (repository *KnowledgeEntryRepository) GetAllKnowledgeEntries() []model.KnowledgeEntry {

}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryById(id int) model.KnowledgeEntry  {

}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryByTitle(title string) model.KnowledgeEntry {

}

func (repository *KnowledgeEntryRepository) AddKnowledgeEntry(space *model.KnowledgeEntry) model.KnowledgeEntry {

}

func (repository *KnowledgeEntryRepository) DeleteKnowledgeEntry(id int) {

}

func (repository *KnowledgeEntryRepository) ModifyKnowledgeEntry(id int,space *model.KnowledgeEntry) model.KnowledgeEntry {

}

func (repository *KnowledgeEntryRepository) InitKnowledgeEntryRepository() KnowledgeEntryRepository {
	return KnowledgeEntryRepository{}
}