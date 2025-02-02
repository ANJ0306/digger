// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newRepo(db *gorm.DB, opts ...gen.DOOption) repo {
	_repo := repo{}

	_repo.repoDo.UseDB(db, opts...)
	_repo.repoDo.UseModel(&model.Repo{})

	tableName := _repo.repoDo.TableName()
	_repo.ALL = field.NewAsterisk(tableName)
	_repo.ID = field.NewInt64(tableName, "id")
	_repo.CreatedAt = field.NewTime(tableName, "created_at")
	_repo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_repo.DeletedAt = field.NewField(tableName, "deleted_at")
	_repo.Name = field.NewString(tableName, "name")
	_repo.OrganizationID = field.NewString(tableName, "organization_id")
	_repo.DiggerConfig = field.NewString(tableName, "digger_config")
	_repo.RepoName = field.NewString(tableName, "repo_name")
	_repo.RepoFullName = field.NewString(tableName, "repo_full_name")
	_repo.RepoOrganisation = field.NewString(tableName, "repo_organisation")
	_repo.RepoURL = field.NewString(tableName, "repo_url")

	_repo.fillFieldMap()

	return _repo
}

type repo struct {
	repoDo

	ALL              field.Asterisk
	ID               field.Int64
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Name             field.String
	OrganizationID   field.String
	DiggerConfig     field.String
	RepoName         field.String
	RepoFullName     field.String
	RepoOrganisation field.String
	RepoURL          field.String

	fieldMap map[string]field.Expr
}

func (r repo) Table(newTableName string) *repo {
	r.repoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r repo) As(alias string) *repo {
	r.repoDo.DO = *(r.repoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *repo) updateTableName(table string) *repo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.Name = field.NewString(table, "name")
	r.OrganizationID = field.NewString(table, "organization_id")
	r.DiggerConfig = field.NewString(table, "digger_config")
	r.RepoName = field.NewString(table, "repo_name")
	r.RepoFullName = field.NewString(table, "repo_full_name")
	r.RepoOrganisation = field.NewString(table, "repo_organisation")
	r.RepoURL = field.NewString(table, "repo_url")

	r.fillFieldMap()

	return r
}

func (r *repo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *repo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 11)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["name"] = r.Name
	r.fieldMap["organization_id"] = r.OrganizationID
	r.fieldMap["digger_config"] = r.DiggerConfig
	r.fieldMap["repo_name"] = r.RepoName
	r.fieldMap["repo_full_name"] = r.RepoFullName
	r.fieldMap["repo_organisation"] = r.RepoOrganisation
	r.fieldMap["repo_url"] = r.RepoURL
}

func (r repo) clone(db *gorm.DB) repo {
	r.repoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r repo) replaceDB(db *gorm.DB) repo {
	r.repoDo.ReplaceDB(db)
	return r
}

type repoDo struct{ gen.DO }

type IRepoDo interface {
	gen.SubQuery
	Debug() IRepoDo
	WithContext(ctx context.Context) IRepoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRepoDo
	WriteDB() IRepoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRepoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRepoDo
	Not(conds ...gen.Condition) IRepoDo
	Or(conds ...gen.Condition) IRepoDo
	Select(conds ...field.Expr) IRepoDo
	Where(conds ...gen.Condition) IRepoDo
	Order(conds ...field.Expr) IRepoDo
	Distinct(cols ...field.Expr) IRepoDo
	Omit(cols ...field.Expr) IRepoDo
	Join(table schema.Tabler, on ...field.Expr) IRepoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRepoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRepoDo
	Group(cols ...field.Expr) IRepoDo
	Having(conds ...gen.Condition) IRepoDo
	Limit(limit int) IRepoDo
	Offset(offset int) IRepoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRepoDo
	Unscoped() IRepoDo
	Create(values ...*model.Repo) error
	CreateInBatches(values []*model.Repo, batchSize int) error
	Save(values ...*model.Repo) error
	First() (*model.Repo, error)
	Take() (*model.Repo, error)
	Last() (*model.Repo, error)
	Find() ([]*model.Repo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Repo, err error)
	FindInBatches(result *[]*model.Repo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Repo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRepoDo
	Assign(attrs ...field.AssignExpr) IRepoDo
	Joins(fields ...field.RelationField) IRepoDo
	Preload(fields ...field.RelationField) IRepoDo
	FirstOrInit() (*model.Repo, error)
	FirstOrCreate() (*model.Repo, error)
	FindByPage(offset int, limit int) (result []*model.Repo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRepoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r repoDo) Debug() IRepoDo {
	return r.withDO(r.DO.Debug())
}

func (r repoDo) WithContext(ctx context.Context) IRepoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r repoDo) ReadDB() IRepoDo {
	return r.Clauses(dbresolver.Read)
}

func (r repoDo) WriteDB() IRepoDo {
	return r.Clauses(dbresolver.Write)
}

func (r repoDo) Session(config *gorm.Session) IRepoDo {
	return r.withDO(r.DO.Session(config))
}

func (r repoDo) Clauses(conds ...clause.Expression) IRepoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r repoDo) Returning(value interface{}, columns ...string) IRepoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r repoDo) Not(conds ...gen.Condition) IRepoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r repoDo) Or(conds ...gen.Condition) IRepoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r repoDo) Select(conds ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r repoDo) Where(conds ...gen.Condition) IRepoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r repoDo) Order(conds ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r repoDo) Distinct(cols ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r repoDo) Omit(cols ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r repoDo) Join(table schema.Tabler, on ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r repoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRepoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r repoDo) RightJoin(table schema.Tabler, on ...field.Expr) IRepoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r repoDo) Group(cols ...field.Expr) IRepoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r repoDo) Having(conds ...gen.Condition) IRepoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r repoDo) Limit(limit int) IRepoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r repoDo) Offset(offset int) IRepoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r repoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRepoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r repoDo) Unscoped() IRepoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r repoDo) Create(values ...*model.Repo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r repoDo) CreateInBatches(values []*model.Repo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r repoDo) Save(values ...*model.Repo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r repoDo) First() (*model.Repo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repo), nil
	}
}

func (r repoDo) Take() (*model.Repo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repo), nil
	}
}

func (r repoDo) Last() (*model.Repo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repo), nil
	}
}

func (r repoDo) Find() ([]*model.Repo, error) {
	result, err := r.DO.Find()
	return result.([]*model.Repo), err
}

func (r repoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Repo, err error) {
	buf := make([]*model.Repo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r repoDo) FindInBatches(result *[]*model.Repo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r repoDo) Attrs(attrs ...field.AssignExpr) IRepoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r repoDo) Assign(attrs ...field.AssignExpr) IRepoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r repoDo) Joins(fields ...field.RelationField) IRepoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r repoDo) Preload(fields ...field.RelationField) IRepoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r repoDo) FirstOrInit() (*model.Repo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repo), nil
	}
}

func (r repoDo) FirstOrCreate() (*model.Repo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repo), nil
	}
}

func (r repoDo) FindByPage(offset int, limit int) (result []*model.Repo, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r repoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r repoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r repoDo) Delete(models ...*model.Repo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *repoDo) withDO(do gen.Dao) *repoDo {
	r.DO = *do.(*gen.DO)
	return r
}
