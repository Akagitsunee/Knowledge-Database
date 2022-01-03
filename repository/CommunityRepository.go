package repository

import (
	"fmt"
	"github.com/blockloop/scan"
	"kdb/model"
)

type CommunityRepository struct {
}

func (repository *CommunityRepository) GetAllCommunitys() ([]model.Community, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.Community;")

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Community

	scan.Rows(&list, rows)

	return list, nil
}

func (repository *CommunityRepository) GetCommunityById(id string) (*model.Community, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.Community WHERE kdb.Community.CommunityId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entry model.Community

	scan.Row(&entry, rows)

	return &entry, nil
}

func (repository *CommunityRepository) GetCommunityByName(name string) (*model.Community, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM TestSchema.Community WHERE name = @name;")

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return &model.Community{}, err
	}

	defer rows.Close()

	var entry model.Community

	scan.Row(entry, rows)

	return &entry, nil
}

func (repository *CommunityRepository) AddCommunity(community *model.Community) (*model.Community, error) {
	checkIfAlive()

	query := fmt.Sprintf("INSERT INTO kdb.Community([Name],[AdminId]) VALUES('%s', %d)", community.Name, community.AdminId)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return community, nil
}

func (repository *CommunityRepository) DeleteCommunity(id string) error {
	checkIfAlive()

	query := fmt.Sprintf("DELETE FROM [kdb].[Community] WHERE CommunityId = %s", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (repository *CommunityRepository) ModifyCommunity(id string, community *model.Community) (*model.Community, error) {
	checkIfAlive()

	query := fmt.Sprintf("UPDATE [kdb].[Community] SET [Name] = '%s',[AdminId] = %d WHERE [kdb].[Community].CommunityId = %s",
		community.Name, community.AdminId, id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	community.CommunityId = id
	return community, nil
}

func (repository *CommunityRepository) AddEmployee(cId string, eId string) error {
	checkIfAlive()

	query := fmt.Sprintf("INSERT INTO kdb.EmployeeCommunity VALUES(%s, %s)", eId, cId)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (repository *CommunityRepository) RemoveEmployee(cId string, eId string) error {
	checkIfAlive()

	query := fmt.Sprintf("DELETE FROM [kdb].[EmployeeCommunity] WHERE CommunityId = %s AND EmployeeId = %s", cId, eId)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (repository *CommunityRepository) GetAllCommunityMembersById(id string) ([]model.Employee, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT e.EmployeeId, e.Name, e.UserId FROM kdb.Employee e JOIN kdb.EmployeeCommunity ec ON ec.EmployeeID = e.EmployeeId JOIN kdb.Community c ON ec.CommunityID = c.CommunityId WHERE ec.CommunityID = %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var l []model.Employee

	scan.Rows(&l, rows)

	return l, nil
}

func (repository *CommunityRepository) InitCommunityRepository() CommunityRepository {
	return CommunityRepository{}
}
