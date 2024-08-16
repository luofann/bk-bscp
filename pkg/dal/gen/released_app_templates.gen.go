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

func newReleasedAppTemplate(db *gorm.DB, opts ...gen.DOOption) releasedAppTemplate {
	_releasedAppTemplate := releasedAppTemplate{}

	_releasedAppTemplate.releasedAppTemplateDo.UseDB(db, opts...)
	_releasedAppTemplate.releasedAppTemplateDo.UseModel(&table.ReleasedAppTemplate{})

	tableName := _releasedAppTemplate.releasedAppTemplateDo.TableName()
	_releasedAppTemplate.ALL = field.NewAsterisk(tableName)
	_releasedAppTemplate.ID = field.NewUint32(tableName, "id")
	_releasedAppTemplate.ReleaseID = field.NewUint32(tableName, "release_id")
	_releasedAppTemplate.TemplateSpaceID = field.NewUint32(tableName, "template_space_id")
	_releasedAppTemplate.TemplateSpaceName = field.NewString(tableName, "template_space_name")
	_releasedAppTemplate.TemplateSetID = field.NewUint32(tableName, "template_set_id")
	_releasedAppTemplate.TemplateSetName = field.NewString(tableName, "template_set_name")
	_releasedAppTemplate.TemplateID = field.NewUint32(tableName, "template_id")
	_releasedAppTemplate.Name = field.NewString(tableName, "name")
	_releasedAppTemplate.Path = field.NewString(tableName, "path")
	_releasedAppTemplate.TemplateRevisionID = field.NewUint32(tableName, "template_revision_id")
	_releasedAppTemplate.IsLatest = field.NewBool(tableName, "is_latest")
	_releasedAppTemplate.TemplateRevisionName = field.NewString(tableName, "template_revision_name")
	_releasedAppTemplate.TemplateRevisionMemo = field.NewString(tableName, "template_revision_memo")
	_releasedAppTemplate.FileType = field.NewString(tableName, "file_type")
	_releasedAppTemplate.FileMode = field.NewString(tableName, "file_mode")
	_releasedAppTemplate.User = field.NewString(tableName, "user")
	_releasedAppTemplate.UserGroup = field.NewString(tableName, "user_group")
	_releasedAppTemplate.Privilege = field.NewString(tableName, "privilege")
	_releasedAppTemplate.Signature = field.NewString(tableName, "signature")
	_releasedAppTemplate.ByteSize = field.NewUint64(tableName, "byte_size")
	_releasedAppTemplate.OriginSignature = field.NewString(tableName, "origin_signature")
	_releasedAppTemplate.OriginByteSize = field.NewUint64(tableName, "origin_byte_size")
	_releasedAppTemplate.Md5 = field.NewString(tableName, "md5")
	_releasedAppTemplate.BizID = field.NewUint32(tableName, "biz_id")
	_releasedAppTemplate.AppID = field.NewUint32(tableName, "app_id")
	_releasedAppTemplate.Creator = field.NewString(tableName, "creator")
	_releasedAppTemplate.Reviser = field.NewString(tableName, "reviser")
	_releasedAppTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_releasedAppTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")

	_releasedAppTemplate.fillFieldMap()

	return _releasedAppTemplate
}

type releasedAppTemplate struct {
	releasedAppTemplateDo releasedAppTemplateDo

	ALL                  field.Asterisk
	ID                   field.Uint32
	ReleaseID            field.Uint32
	TemplateSpaceID      field.Uint32
	TemplateSpaceName    field.String
	TemplateSetID        field.Uint32
	TemplateSetName      field.String
	TemplateID           field.Uint32
	Name                 field.String
	Path                 field.String
	TemplateRevisionID   field.Uint32
	IsLatest             field.Bool
	TemplateRevisionName field.String
	TemplateRevisionMemo field.String
	FileType             field.String
	FileMode             field.String
	User                 field.String
	UserGroup            field.String
	Privilege            field.String
	Signature            field.String
	ByteSize             field.Uint64
	OriginSignature      field.String
	OriginByteSize       field.Uint64
	Md5                  field.String
	BizID                field.Uint32
	AppID                field.Uint32
	Creator              field.String
	Reviser              field.String
	CreatedAt            field.Time
	UpdatedAt            field.Time

	fieldMap map[string]field.Expr
}

func (r releasedAppTemplate) Table(newTableName string) *releasedAppTemplate {
	r.releasedAppTemplateDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r releasedAppTemplate) As(alias string) *releasedAppTemplate {
	r.releasedAppTemplateDo.DO = *(r.releasedAppTemplateDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *releasedAppTemplate) updateTableName(table string) *releasedAppTemplate {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint32(table, "id")
	r.ReleaseID = field.NewUint32(table, "release_id")
	r.TemplateSpaceID = field.NewUint32(table, "template_space_id")
	r.TemplateSpaceName = field.NewString(table, "template_space_name")
	r.TemplateSetID = field.NewUint32(table, "template_set_id")
	r.TemplateSetName = field.NewString(table, "template_set_name")
	r.TemplateID = field.NewUint32(table, "template_id")
	r.Name = field.NewString(table, "name")
	r.Path = field.NewString(table, "path")
	r.TemplateRevisionID = field.NewUint32(table, "template_revision_id")
	r.IsLatest = field.NewBool(table, "is_latest")
	r.TemplateRevisionName = field.NewString(table, "template_revision_name")
	r.TemplateRevisionMemo = field.NewString(table, "template_revision_memo")
	r.FileType = field.NewString(table, "file_type")
	r.FileMode = field.NewString(table, "file_mode")
	r.User = field.NewString(table, "user")
	r.UserGroup = field.NewString(table, "user_group")
	r.Privilege = field.NewString(table, "privilege")
	r.Signature = field.NewString(table, "signature")
	r.ByteSize = field.NewUint64(table, "byte_size")
	r.OriginSignature = field.NewString(table, "origin_signature")
	r.OriginByteSize = field.NewUint64(table, "origin_byte_size")
	r.Md5 = field.NewString(table, "md5")
	r.BizID = field.NewUint32(table, "biz_id")
	r.AppID = field.NewUint32(table, "app_id")
	r.Creator = field.NewString(table, "creator")
	r.Reviser = field.NewString(table, "reviser")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *releasedAppTemplate) WithContext(ctx context.Context) IReleasedAppTemplateDo {
	return r.releasedAppTemplateDo.WithContext(ctx)
}

func (r releasedAppTemplate) TableName() string { return r.releasedAppTemplateDo.TableName() }

func (r releasedAppTemplate) Alias() string { return r.releasedAppTemplateDo.Alias() }

func (r releasedAppTemplate) Columns(cols ...field.Expr) gen.Columns {
	return r.releasedAppTemplateDo.Columns(cols...)
}

func (r *releasedAppTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *releasedAppTemplate) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 29)
	r.fieldMap["id"] = r.ID
	r.fieldMap["release_id"] = r.ReleaseID
	r.fieldMap["template_space_id"] = r.TemplateSpaceID
	r.fieldMap["template_space_name"] = r.TemplateSpaceName
	r.fieldMap["template_set_id"] = r.TemplateSetID
	r.fieldMap["template_set_name"] = r.TemplateSetName
	r.fieldMap["template_id"] = r.TemplateID
	r.fieldMap["name"] = r.Name
	r.fieldMap["path"] = r.Path
	r.fieldMap["template_revision_id"] = r.TemplateRevisionID
	r.fieldMap["is_latest"] = r.IsLatest
	r.fieldMap["template_revision_name"] = r.TemplateRevisionName
	r.fieldMap["template_revision_memo"] = r.TemplateRevisionMemo
	r.fieldMap["file_type"] = r.FileType
	r.fieldMap["file_mode"] = r.FileMode
	r.fieldMap["user"] = r.User
	r.fieldMap["user_group"] = r.UserGroup
	r.fieldMap["privilege"] = r.Privilege
	r.fieldMap["signature"] = r.Signature
	r.fieldMap["byte_size"] = r.ByteSize
	r.fieldMap["origin_signature"] = r.OriginSignature
	r.fieldMap["origin_byte_size"] = r.OriginByteSize
	r.fieldMap["md5"] = r.Md5
	r.fieldMap["biz_id"] = r.BizID
	r.fieldMap["app_id"] = r.AppID
	r.fieldMap["creator"] = r.Creator
	r.fieldMap["reviser"] = r.Reviser
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r releasedAppTemplate) clone(db *gorm.DB) releasedAppTemplate {
	r.releasedAppTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r releasedAppTemplate) replaceDB(db *gorm.DB) releasedAppTemplate {
	r.releasedAppTemplateDo.ReplaceDB(db)
	return r
}

type releasedAppTemplateDo struct{ gen.DO }

type IReleasedAppTemplateDo interface {
	gen.SubQuery
	Debug() IReleasedAppTemplateDo
	WithContext(ctx context.Context) IReleasedAppTemplateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReleasedAppTemplateDo
	WriteDB() IReleasedAppTemplateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReleasedAppTemplateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReleasedAppTemplateDo
	Not(conds ...gen.Condition) IReleasedAppTemplateDo
	Or(conds ...gen.Condition) IReleasedAppTemplateDo
	Select(conds ...field.Expr) IReleasedAppTemplateDo
	Where(conds ...gen.Condition) IReleasedAppTemplateDo
	Order(conds ...field.Expr) IReleasedAppTemplateDo
	Distinct(cols ...field.Expr) IReleasedAppTemplateDo
	Omit(cols ...field.Expr) IReleasedAppTemplateDo
	Join(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo
	Group(cols ...field.Expr) IReleasedAppTemplateDo
	Having(conds ...gen.Condition) IReleasedAppTemplateDo
	Limit(limit int) IReleasedAppTemplateDo
	Offset(offset int) IReleasedAppTemplateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedAppTemplateDo
	Unscoped() IReleasedAppTemplateDo
	Create(values ...*table.ReleasedAppTemplate) error
	CreateInBatches(values []*table.ReleasedAppTemplate, batchSize int) error
	Save(values ...*table.ReleasedAppTemplate) error
	First() (*table.ReleasedAppTemplate, error)
	Take() (*table.ReleasedAppTemplate, error)
	Last() (*table.ReleasedAppTemplate, error)
	Find() ([]*table.ReleasedAppTemplate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedAppTemplate, err error)
	FindInBatches(result *[]*table.ReleasedAppTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ReleasedAppTemplate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReleasedAppTemplateDo
	Assign(attrs ...field.AssignExpr) IReleasedAppTemplateDo
	Joins(fields ...field.RelationField) IReleasedAppTemplateDo
	Preload(fields ...field.RelationField) IReleasedAppTemplateDo
	FirstOrInit() (*table.ReleasedAppTemplate, error)
	FirstOrCreate() (*table.ReleasedAppTemplate, error)
	FindByPage(offset int, limit int) (result []*table.ReleasedAppTemplate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReleasedAppTemplateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r releasedAppTemplateDo) Debug() IReleasedAppTemplateDo {
	return r.withDO(r.DO.Debug())
}

func (r releasedAppTemplateDo) WithContext(ctx context.Context) IReleasedAppTemplateDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r releasedAppTemplateDo) ReadDB() IReleasedAppTemplateDo {
	return r.Clauses(dbresolver.Read)
}

func (r releasedAppTemplateDo) WriteDB() IReleasedAppTemplateDo {
	return r.Clauses(dbresolver.Write)
}

func (r releasedAppTemplateDo) Session(config *gorm.Session) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Session(config))
}

func (r releasedAppTemplateDo) Clauses(conds ...clause.Expression) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r releasedAppTemplateDo) Returning(value interface{}, columns ...string) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r releasedAppTemplateDo) Not(conds ...gen.Condition) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r releasedAppTemplateDo) Or(conds ...gen.Condition) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r releasedAppTemplateDo) Select(conds ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r releasedAppTemplateDo) Where(conds ...gen.Condition) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r releasedAppTemplateDo) Order(conds ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r releasedAppTemplateDo) Distinct(cols ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r releasedAppTemplateDo) Omit(cols ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r releasedAppTemplateDo) Join(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r releasedAppTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r releasedAppTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r releasedAppTemplateDo) Group(cols ...field.Expr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r releasedAppTemplateDo) Having(conds ...gen.Condition) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r releasedAppTemplateDo) Limit(limit int) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r releasedAppTemplateDo) Offset(offset int) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r releasedAppTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r releasedAppTemplateDo) Unscoped() IReleasedAppTemplateDo {
	return r.withDO(r.DO.Unscoped())
}

func (r releasedAppTemplateDo) Create(values ...*table.ReleasedAppTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r releasedAppTemplateDo) CreateInBatches(values []*table.ReleasedAppTemplate, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r releasedAppTemplateDo) Save(values ...*table.ReleasedAppTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r releasedAppTemplateDo) First() (*table.ReleasedAppTemplate, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedAppTemplate), nil
	}
}

func (r releasedAppTemplateDo) Take() (*table.ReleasedAppTemplate, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedAppTemplate), nil
	}
}

func (r releasedAppTemplateDo) Last() (*table.ReleasedAppTemplate, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedAppTemplate), nil
	}
}

func (r releasedAppTemplateDo) Find() ([]*table.ReleasedAppTemplate, error) {
	result, err := r.DO.Find()
	return result.([]*table.ReleasedAppTemplate), err
}

func (r releasedAppTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedAppTemplate, err error) {
	buf := make([]*table.ReleasedAppTemplate, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r releasedAppTemplateDo) FindInBatches(result *[]*table.ReleasedAppTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r releasedAppTemplateDo) Attrs(attrs ...field.AssignExpr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r releasedAppTemplateDo) Assign(attrs ...field.AssignExpr) IReleasedAppTemplateDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r releasedAppTemplateDo) Joins(fields ...field.RelationField) IReleasedAppTemplateDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r releasedAppTemplateDo) Preload(fields ...field.RelationField) IReleasedAppTemplateDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r releasedAppTemplateDo) FirstOrInit() (*table.ReleasedAppTemplate, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedAppTemplate), nil
	}
}

func (r releasedAppTemplateDo) FirstOrCreate() (*table.ReleasedAppTemplate, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedAppTemplate), nil
	}
}

func (r releasedAppTemplateDo) FindByPage(offset int, limit int) (result []*table.ReleasedAppTemplate, count int64, err error) {
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

func (r releasedAppTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r releasedAppTemplateDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r releasedAppTemplateDo) Delete(models ...*table.ReleasedAppTemplate) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *releasedAppTemplateDo) withDO(do gen.Dao) *releasedAppTemplateDo {
	r.DO = *do.(*gen.DO)
	return r
}
