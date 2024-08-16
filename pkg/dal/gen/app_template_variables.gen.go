// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newAppTemplateVariable(db *gorm.DB, opts ...gen.DOOption) appTemplateVariable {
	_appTemplateVariable := appTemplateVariable{}

	_appTemplateVariable.appTemplateVariableDo.UseDB(db, opts...)
	_appTemplateVariable.appTemplateVariableDo.UseModel(&table.AppTemplateVariable{})

	tableName := _appTemplateVariable.appTemplateVariableDo.TableName()
	_appTemplateVariable.ALL = field.NewAsterisk(tableName)
	_appTemplateVariable.ID = field.NewUint32(tableName, "id")
	_appTemplateVariable.Variables = field.NewField(tableName, "variables")
	_appTemplateVariable.BizID = field.NewUint32(tableName, "biz_id")
	_appTemplateVariable.AppID = field.NewUint32(tableName, "app_id")
	_appTemplateVariable.Creator = field.NewString(tableName, "creator")
	_appTemplateVariable.Reviser = field.NewString(tableName, "reviser")
	_appTemplateVariable.CreatedAt = field.NewTime(tableName, "created_at")
	_appTemplateVariable.UpdatedAt = field.NewTime(tableName, "updated_at")

	_appTemplateVariable.fillFieldMap()

	return _appTemplateVariable
}

type appTemplateVariable struct {
	appTemplateVariableDo appTemplateVariableDo

	ALL       field.Asterisk
	ID        field.Uint32
	Variables field.Field
	BizID     field.Uint32
	AppID     field.Uint32
	Creator   field.String
	Reviser   field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a appTemplateVariable) Table(newTableName string) *appTemplateVariable {
	a.appTemplateVariableDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appTemplateVariable) As(alias string) *appTemplateVariable {
	a.appTemplateVariableDo.DO = *(a.appTemplateVariableDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appTemplateVariable) updateTableName(table string) *appTemplateVariable {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.Variables = field.NewField(table, "variables")
	a.BizID = field.NewUint32(table, "biz_id")
	a.AppID = field.NewUint32(table, "app_id")
	a.Creator = field.NewString(table, "creator")
	a.Reviser = field.NewString(table, "reviser")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *appTemplateVariable) WithContext(ctx context.Context) IAppTemplateVariableDo {
	return a.appTemplateVariableDo.WithContext(ctx)
}

func (a appTemplateVariable) TableName() string { return a.appTemplateVariableDo.TableName() }

func (a appTemplateVariable) Alias() string { return a.appTemplateVariableDo.Alias() }

func (a appTemplateVariable) Columns(cols ...field.Expr) gen.Columns {
	return a.appTemplateVariableDo.Columns(cols...)
}

func (a *appTemplateVariable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appTemplateVariable) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["variables"] = a.Variables
	a.fieldMap["biz_id"] = a.BizID
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["creator"] = a.Creator
	a.fieldMap["reviser"] = a.Reviser
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a appTemplateVariable) clone(db *gorm.DB) appTemplateVariable {
	a.appTemplateVariableDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appTemplateVariable) replaceDB(db *gorm.DB) appTemplateVariable {
	a.appTemplateVariableDo.ReplaceDB(db)
	return a
}

type appTemplateVariableDo struct{ gen.DO }

type IAppTemplateVariableDo interface {
	gen.SubQuery
	Debug() IAppTemplateVariableDo
	WithContext(ctx context.Context) IAppTemplateVariableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppTemplateVariableDo
	WriteDB() IAppTemplateVariableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppTemplateVariableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppTemplateVariableDo
	Not(conds ...gen.Condition) IAppTemplateVariableDo
	Or(conds ...gen.Condition) IAppTemplateVariableDo
	Select(conds ...field.Expr) IAppTemplateVariableDo
	Where(conds ...gen.Condition) IAppTemplateVariableDo
	Order(conds ...field.Expr) IAppTemplateVariableDo
	Distinct(cols ...field.Expr) IAppTemplateVariableDo
	Omit(cols ...field.Expr) IAppTemplateVariableDo
	Join(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo
	Group(cols ...field.Expr) IAppTemplateVariableDo
	Having(conds ...gen.Condition) IAppTemplateVariableDo
	Limit(limit int) IAppTemplateVariableDo
	Offset(offset int) IAppTemplateVariableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppTemplateVariableDo
	Unscoped() IAppTemplateVariableDo
	Create(values ...*table.AppTemplateVariable) error
	CreateInBatches(values []*table.AppTemplateVariable, batchSize int) error
	Save(values ...*table.AppTemplateVariable) error
	First() (*table.AppTemplateVariable, error)
	Take() (*table.AppTemplateVariable, error)
	Last() (*table.AppTemplateVariable, error)
	Find() ([]*table.AppTemplateVariable, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AppTemplateVariable, err error)
	FindInBatches(result *[]*table.AppTemplateVariable, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.AppTemplateVariable) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppTemplateVariableDo
	Assign(attrs ...field.AssignExpr) IAppTemplateVariableDo
	Joins(fields ...field.RelationField) IAppTemplateVariableDo
	Preload(fields ...field.RelationField) IAppTemplateVariableDo
	FirstOrInit() (*table.AppTemplateVariable, error)
	FirstOrCreate() (*table.AppTemplateVariable, error)
	FindByPage(offset int, limit int) (result []*table.AppTemplateVariable, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppTemplateVariableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appTemplateVariableDo) Debug() IAppTemplateVariableDo {
	return a.withDO(a.DO.Debug())
}

func (a appTemplateVariableDo) WithContext(ctx context.Context) IAppTemplateVariableDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appTemplateVariableDo) ReadDB() IAppTemplateVariableDo {
	return a.Clauses(dbresolver.Read)
}

func (a appTemplateVariableDo) WriteDB() IAppTemplateVariableDo {
	return a.Clauses(dbresolver.Write)
}

func (a appTemplateVariableDo) Session(config *gorm.Session) IAppTemplateVariableDo {
	return a.withDO(a.DO.Session(config))
}

func (a appTemplateVariableDo) Clauses(conds ...clause.Expression) IAppTemplateVariableDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appTemplateVariableDo) Returning(value interface{}, columns ...string) IAppTemplateVariableDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appTemplateVariableDo) Not(conds ...gen.Condition) IAppTemplateVariableDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appTemplateVariableDo) Or(conds ...gen.Condition) IAppTemplateVariableDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appTemplateVariableDo) Select(conds ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appTemplateVariableDo) Where(conds ...gen.Condition) IAppTemplateVariableDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appTemplateVariableDo) Order(conds ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appTemplateVariableDo) Distinct(cols ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appTemplateVariableDo) Omit(cols ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appTemplateVariableDo) Join(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appTemplateVariableDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appTemplateVariableDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appTemplateVariableDo) Group(cols ...field.Expr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appTemplateVariableDo) Having(conds ...gen.Condition) IAppTemplateVariableDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appTemplateVariableDo) Limit(limit int) IAppTemplateVariableDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appTemplateVariableDo) Offset(offset int) IAppTemplateVariableDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appTemplateVariableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppTemplateVariableDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appTemplateVariableDo) Unscoped() IAppTemplateVariableDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appTemplateVariableDo) Create(values ...*table.AppTemplateVariable) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appTemplateVariableDo) CreateInBatches(values []*table.AppTemplateVariable, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appTemplateVariableDo) Save(values ...*table.AppTemplateVariable) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appTemplateVariableDo) First() (*table.AppTemplateVariable, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateVariable), nil
	}
}

func (a appTemplateVariableDo) Take() (*table.AppTemplateVariable, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateVariable), nil
	}
}

func (a appTemplateVariableDo) Last() (*table.AppTemplateVariable, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateVariable), nil
	}
}

func (a appTemplateVariableDo) Find() ([]*table.AppTemplateVariable, error) {
	result, err := a.DO.Find()
	return result.([]*table.AppTemplateVariable), err
}

func (a appTemplateVariableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AppTemplateVariable, err error) {
	buf := make([]*table.AppTemplateVariable, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appTemplateVariableDo) FindInBatches(result *[]*table.AppTemplateVariable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appTemplateVariableDo) Attrs(attrs ...field.AssignExpr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appTemplateVariableDo) Assign(attrs ...field.AssignExpr) IAppTemplateVariableDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appTemplateVariableDo) Joins(fields ...field.RelationField) IAppTemplateVariableDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appTemplateVariableDo) Preload(fields ...field.RelationField) IAppTemplateVariableDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appTemplateVariableDo) FirstOrInit() (*table.AppTemplateVariable, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateVariable), nil
	}
}

func (a appTemplateVariableDo) FirstOrCreate() (*table.AppTemplateVariable, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateVariable), nil
	}
}

func (a appTemplateVariableDo) FindByPage(offset int, limit int) (result []*table.AppTemplateVariable, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appTemplateVariableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appTemplateVariableDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appTemplateVariableDo) Delete(models ...*table.AppTemplateVariable) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appTemplateVariableDo) withDO(do gen.Dao) *appTemplateVariableDo {
	a.DO = *do.(*gen.DO)
	return a
}
