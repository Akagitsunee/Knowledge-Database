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

	query := fmt.Sprintf("INSERT INTO kdb.KnowledgeEntry([Title],[Content],[CreatorId],[CommunityId]) VALUES('%s', '%s', %d,%d)",
		entry.Title, entry.Content, entry.CreatorId, entry.CommunityId)

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

	query := fmt.Sprintf("UPDATE [kdb].[KnowledgeEntry] SET [Deleted] = 1 WHERE [kdb].[KnowledgeEntry].KnowledgeEntryId = %s", id)

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

	query := fmt.Sprintf("INSERT INTO kdb.KnowledgeEntryVersion([Title],[Content],[EditorId],[EntryId]) VALUES('%s', '%s', %d, %s)",
		entry.Title, entry.Content, entry.EditorId, id)

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

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntryVersion WHERE kdb.KnowledgeEntryVersion.KnowledgeEntryVersionId = %s;", id)

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

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntryVersion WHERE kdb.KnowledgeEntryVersion.EntryId = %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var versionList []model.KnowledgeEntryVersion

	scan.Rows(&versionList, rows)

	return versionList, nil
}

func (repository *KnowledgeEntryRepository) GetKnowledgeEntryVersionsByCommunity(id string) ([]model.KnowledgeEntry, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.KnowledgeEntry ke JOIN kdb.Community c ON ke.CommunityId = c.CommunityId WHERE c.CommunityId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var l []model.KnowledgeEntry

	scan.Rows(&l, rows)

	return l, nil
}

func (repository *KnowledgeEntryRepository) InitKnowledgeEntryRepository() KnowledgeEntryRepository {
	return KnowledgeEntryRepository{}
}
