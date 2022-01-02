package repository

import (
	"fmt"
	"github.com/blockloop/scan"
	"kdb/model"
)

type KnowledgeEntryRepository struct {
}

func (repository *KnowledgeEntryRepository) GetAllKnowledgeEntries() ([]model.KnowledgeEntry, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntry;")

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.KnowledgeEntry

	scan.Rows(&list, rows)

	return list, nil
}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryById(id string) (*model.KnowledgeEntry, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntry WHERE kdb.KnowledgeEntry.KnowledgeEntryId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entry model.KnowledgeEntry

	scan.Row(&entry, rows)

	return &entry, nil
}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryByTitle(title string) (*model.KnowledgeEntry, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM TestSchema.KnowledgeEntry WHERE id = @title;")

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return &model.KnowledgeEntry{}, err
	}

	defer rows.Close()

	var entry model.KnowledgeEntry

	scan.Row(entry, rows)

	return &entry, nil
}

func (repository *KnowledgeEntryRepository) AddKnowledgeEntry(entry *model.KnowledgeEntry) (*model.KnowledgeEntry, error) {
	checkIfAlive()

	query := fmt.Sprintf("INSERT INTO kdb.KnowledgeEntry([Title],[Content],[Creator],[Community]) VALUES('%s', '%s', %d,%d)",
		entry.Title, entry.Content, entry.Creator, entry.CommunityId)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return entry, nil
}

func (repository *KnowledgeEntryRepository) DeleteKnowledgeEntry(id string) error {
	checkIfAlive()

	query := fmt.Sprintf("DELETE FROM [kdb].[KnowledgeEntry] WHERE KnowledgeEntryId = %s", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (repository *KnowledgeEntryRepository) ModifyKnowledgeEntry(id string, entry *model.KnowledgeEntryVersion) (*model.KnowledgeEntryVersion, error) {
	checkIfAlive()

	query := fmt.Sprintf("INSERT INTO kdb.KnowledgeEntryVersion([Title],[Content],[Editor],[EntryId]) VALUES('%s', '%s' %d,%d)",
		entry.Title, entry.Content, entry.Editor, id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return entry, nil
}

func (repository *KnowledgeEntryRepository) GetEntryVersionById(id string) (*model.KnowledgeEntryVersion, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntryVersion WHERE kdb.KnowledgeEntryVersion.VersionId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entry model.KnowledgeEntryVersion

	scan.Row(&entry, rows)

	return &entry, nil
}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryVersionsById(id string) ([]model.KnowledgeEntryVersion, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntryVersion WHERE kdb.KnowledgeEntryVersion.EntryId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var versionList []model.KnowledgeEntryVersion

	scan.Row(versionList, rows)

	return versionList, nil
}

func (repository *KnowledgeEntryRepository) InitKnowledgeEntryRepository() KnowledgeEntryRepository {
	return KnowledgeEntryRepository{}
}
