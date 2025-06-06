// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/danielmesquitta/library-manager/internal/model"
)

func newBook(db *gorm.DB, opts ...gen.DOOption) book {
	_book := book{}

	_book.bookDo.UseDB(db, opts...)
	_book.bookDo.UseModel(&model.Book{})

	tableName := _book.bookDo.TableName()
	_book.ALL = field.NewAsterisk(tableName)
	_book.ID = field.NewField(tableName, "id")
	_book.ISBN = field.NewString(tableName, "isbn")
	_book.Title = field.NewString(tableName, "title")
	_book.AuthorID = field.NewField(tableName, "author_id")
	_book.Copies = field.NewUint(tableName, "copies")
	_book.CreatedAt = field.NewTime(tableName, "created_at")
	_book.UpdatedAt = field.NewTime(tableName, "updated_at")
	_book.DeletedAt = field.NewField(tableName, "deleted_at")
	_book.Author = bookBelongsToAuthor{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Author", "model.Author"),
		Books: struct {
			field.RelationField
			Author struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Author.Books", "model.Book"),
			Author: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Author.Books.Author", "model.Author"),
			},
		},
	}

	_book.fillFieldMap()

	return _book
}

type book struct {
	bookDo bookDo

	ALL       field.Asterisk
	ID        field.Field
	ISBN      field.String
	Title     field.String
	AuthorID  field.Field
	Copies    field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Author    bookBelongsToAuthor

	fieldMap map[string]field.Expr
}

func (b book) Table(newTableName string) *book {
	b.bookDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b book) As(alias string) *book {
	b.bookDo.DO = *(b.bookDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *book) updateTableName(table string) *book {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewField(table, "id")
	b.ISBN = field.NewString(table, "isbn")
	b.Title = field.NewString(table, "title")
	b.AuthorID = field.NewField(table, "author_id")
	b.Copies = field.NewUint(table, "copies")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")

	b.fillFieldMap()

	return b
}

func (b *book) WithContext(ctx context.Context) *bookDo { return b.bookDo.WithContext(ctx) }

func (b book) TableName() string { return b.bookDo.TableName() }

func (b book) Alias() string { return b.bookDo.Alias() }

func (b book) Columns(cols ...field.Expr) gen.Columns { return b.bookDo.Columns(cols...) }

func (b *book) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *book) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 9)
	b.fieldMap["id"] = b.ID
	b.fieldMap["isbn"] = b.ISBN
	b.fieldMap["title"] = b.Title
	b.fieldMap["author_id"] = b.AuthorID
	b.fieldMap["copies"] = b.Copies
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt

}

func (b book) clone(db *gorm.DB) book {
	b.bookDo.ReplaceConnPool(db.Statement.ConnPool)
	b.Author.db = db.Session(&gorm.Session{Initialized: true})
	b.Author.db.Statement.ConnPool = db.Statement.ConnPool
	return b
}

func (b book) replaceDB(db *gorm.DB) book {
	b.bookDo.ReplaceDB(db)
	b.Author.db = db.Session(&gorm.Session{})
	return b
}

type bookBelongsToAuthor struct {
	db *gorm.DB

	field.RelationField

	Books struct {
		field.RelationField
		Author struct {
			field.RelationField
		}
	}
}

func (a bookBelongsToAuthor) Where(conds ...field.Expr) *bookBelongsToAuthor {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a bookBelongsToAuthor) WithContext(ctx context.Context) *bookBelongsToAuthor {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a bookBelongsToAuthor) Session(session *gorm.Session) *bookBelongsToAuthor {
	a.db = a.db.Session(session)
	return &a
}

func (a bookBelongsToAuthor) Model(m *model.Book) *bookBelongsToAuthorTx {
	return &bookBelongsToAuthorTx{a.db.Model(m).Association(a.Name())}
}

func (a bookBelongsToAuthor) Unscoped() *bookBelongsToAuthor {
	a.db = a.db.Unscoped()
	return &a
}

type bookBelongsToAuthorTx struct{ tx *gorm.Association }

func (a bookBelongsToAuthorTx) Find() (result *model.Author, err error) {
	return result, a.tx.Find(&result)
}

func (a bookBelongsToAuthorTx) Append(values ...*model.Author) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a bookBelongsToAuthorTx) Replace(values ...*model.Author) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a bookBelongsToAuthorTx) Delete(values ...*model.Author) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a bookBelongsToAuthorTx) Clear() error {
	return a.tx.Clear()
}

func (a bookBelongsToAuthorTx) Count() int64 {
	return a.tx.Count()
}

func (a bookBelongsToAuthorTx) Unscoped() *bookBelongsToAuthorTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type bookDo struct{ gen.DO }

func (b bookDo) Debug() *bookDo {
	return b.withDO(b.DO.Debug())
}

func (b bookDo) WithContext(ctx context.Context) *bookDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookDo) ReadDB() *bookDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookDo) WriteDB() *bookDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookDo) Session(config *gorm.Session) *bookDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookDo) Clauses(conds ...clause.Expression) *bookDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookDo) Returning(value interface{}, columns ...string) *bookDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookDo) Not(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookDo) Or(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookDo) Select(conds ...field.Expr) *bookDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookDo) Where(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookDo) Order(conds ...field.Expr) *bookDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookDo) Distinct(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookDo) Omit(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookDo) Join(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookDo) RightJoin(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookDo) Group(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookDo) Having(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookDo) Limit(limit int) *bookDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookDo) Offset(offset int) *bookDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bookDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookDo) Unscoped() *bookDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookDo) Create(values ...*model.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookDo) CreateInBatches(values []*model.Book, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookDo) Save(values ...*model.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookDo) First() (*model.Book, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Take() (*model.Book, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Last() (*model.Book, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Find() ([]*model.Book, error) {
	result, err := b.DO.Find()
	return result.([]*model.Book), err
}

func (b bookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Book, err error) {
	buf := make([]*model.Book, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookDo) FindInBatches(result *[]*model.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookDo) Attrs(attrs ...field.AssignExpr) *bookDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookDo) Assign(attrs ...field.AssignExpr) *bookDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookDo) Joins(fields ...field.RelationField) *bookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookDo) Preload(fields ...field.RelationField) *bookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookDo) FirstOrInit() (*model.Book, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) FirstOrCreate() (*model.Book, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) FindByPage(offset int, limit int) (result []*model.Book, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookDo) Delete(models ...*model.Book) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookDo) withDO(do gen.Dao) *bookDo {
	b.DO = *do.(*gen.DO)
	return b
}
