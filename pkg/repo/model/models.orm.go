package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// ModelsN is a Models object, all fields are nullable
type ModelsN struct {
	original    *modelsOriginal
	modelsModel *ModelsModel

	Id            null.Int    `json:"id"`
	ModelId       null.String `json:"model_id"`
	Name          null.String `json:"name"`
	ShortName     null.String `json:"short_name,omitempty"`
	Description   null.String `json:"description,omitempty"`
	AvatarUrl     null.String `json:"avatar_url,omitempty"`
	Status        null.Int    `json:"status"`
	VersionMin    null.String `json:"version_min,omitempty"`
	VersionMax    null.String `json:"version_max,omitempty"`
	MetaJson      null.String `json:"-"`
	ProvidersJson null.String `json:"-"`
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *ModelsN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for Models
func (inst *ModelsN) SetModel(modelsModel *ModelsModel) {
	inst.modelsModel = modelsModel
}

// modelsOriginal is an object which stores original Models from database
type modelsOriginal struct {
	Id            null.Int
	ModelId       null.String
	Name          null.String
	ShortName     null.String
	Description   null.String
	AvatarUrl     null.String
	Status        null.Int
	VersionMin    null.String
	VersionMax    null.String
	MetaJson      null.String
	ProvidersJson null.String
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// Staled identify whether the object has been modified
func (inst *ModelsN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &modelsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.ModelId != inst.original.ModelId {
			return true
		}
		if inst.Name != inst.original.Name {
			return true
		}
		if inst.ShortName != inst.original.ShortName {
			return true
		}
		if inst.Description != inst.original.Description {
			return true
		}
		if inst.AvatarUrl != inst.original.AvatarUrl {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.VersionMin != inst.original.VersionMin {
			return true
		}
		if inst.VersionMax != inst.original.VersionMax {
			return true
		}
		if inst.MetaJson != inst.original.MetaJson {
			return true
		}
		if inst.ProvidersJson != inst.original.ProvidersJson {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "model_id":
				if inst.ModelId != inst.original.ModelId {
					return true
				}
			case "name":
				if inst.Name != inst.original.Name {
					return true
				}
			case "short_name":
				if inst.ShortName != inst.original.ShortName {
					return true
				}
			case "description":
				if inst.Description != inst.original.Description {
					return true
				}
			case "avatar_url":
				if inst.AvatarUrl != inst.original.AvatarUrl {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "version_min":
				if inst.VersionMin != inst.original.VersionMin {
					return true
				}
			case "version_max":
				if inst.VersionMax != inst.original.VersionMax {
					return true
				}
			case "meta_json":
				if inst.MetaJson != inst.original.MetaJson {
					return true
				}
			case "providers_json":
				if inst.ProvidersJson != inst.original.ProvidersJson {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *ModelsN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &modelsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.ModelId != inst.original.ModelId {
			kv["model_id"] = inst.ModelId
		}
		if inst.Name != inst.original.Name {
			kv["name"] = inst.Name
		}
		if inst.ShortName != inst.original.ShortName {
			kv["short_name"] = inst.ShortName
		}
		if inst.Description != inst.original.Description {
			kv["description"] = inst.Description
		}
		if inst.AvatarUrl != inst.original.AvatarUrl {
			kv["avatar_url"] = inst.AvatarUrl
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.VersionMin != inst.original.VersionMin {
			kv["version_min"] = inst.VersionMin
		}
		if inst.VersionMax != inst.original.VersionMax {
			kv["version_max"] = inst.VersionMax
		}
		if inst.MetaJson != inst.original.MetaJson {
			kv["meta_json"] = inst.MetaJson
		}
		if inst.ProvidersJson != inst.original.ProvidersJson {
			kv["providers_json"] = inst.ProvidersJson
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "model_id":
				if inst.ModelId != inst.original.ModelId {
					kv["model_id"] = inst.ModelId
				}
			case "name":
				if inst.Name != inst.original.Name {
					kv["name"] = inst.Name
				}
			case "short_name":
				if inst.ShortName != inst.original.ShortName {
					kv["short_name"] = inst.ShortName
				}
			case "description":
				if inst.Description != inst.original.Description {
					kv["description"] = inst.Description
				}
			case "avatar_url":
				if inst.AvatarUrl != inst.original.AvatarUrl {
					kv["avatar_url"] = inst.AvatarUrl
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "version_min":
				if inst.VersionMin != inst.original.VersionMin {
					kv["version_min"] = inst.VersionMin
				}
			case "version_max":
				if inst.VersionMax != inst.original.VersionMax {
					kv["version_max"] = inst.VersionMax
				}
			case "meta_json":
				if inst.MetaJson != inst.original.MetaJson {
					kv["meta_json"] = inst.MetaJson
				}
			case "providers_json":
				if inst.ProvidersJson != inst.original.ProvidersJson {
					kv["providers_json"] = inst.ProvidersJson
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *ModelsN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.modelsModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.modelsModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a models
func (inst *ModelsN) Delete(ctx context.Context) error {
	if inst.modelsModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.modelsModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *ModelsN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type modelsScope struct {
	name  string
	apply func(builder query.Condition)
}

var modelsGlobalScopes = make([]modelsScope, 0)
var modelsLocalScopes = make([]modelsScope, 0)

// AddGlobalScopeForModels assign a global scope to a model
func AddGlobalScopeForModels(name string, apply func(builder query.Condition)) {
	modelsGlobalScopes = append(modelsGlobalScopes, modelsScope{name: name, apply: apply})
}

// AddLocalScopeForModels assign a local scope to a model
func AddLocalScopeForModels(name string, apply func(builder query.Condition)) {
	modelsLocalScopes = append(modelsLocalScopes, modelsScope{name: name, apply: apply})
}

func (m *ModelsModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range modelsGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range modelsLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ModelsModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ModelsModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type Models struct {
	Id            int64  `json:"id"`
	ModelId       string `json:"model_id"`
	Name          string `json:"name"`
	ShortName     string `json:"short_name,omitempty"`
	Description   string `json:"description,omitempty"`
	AvatarUrl     string `json:"avatar_url,omitempty"`
	Status        int64  `json:"status"`
	VersionMin    string `json:"version_min,omitempty"`
	VersionMax    string `json:"version_max,omitempty"`
	MetaJson      string `json:"-"`
	ProvidersJson string `json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (w Models) ToModelsN(allows ...string) ModelsN {
	if len(allows) == 0 {
		return ModelsN{

			Id:            null.IntFrom(int64(w.Id)),
			ModelId:       null.StringFrom(w.ModelId),
			Name:          null.StringFrom(w.Name),
			ShortName:     null.StringFrom(w.ShortName),
			Description:   null.StringFrom(w.Description),
			AvatarUrl:     null.StringFrom(w.AvatarUrl),
			Status:        null.IntFrom(int64(w.Status)),
			VersionMin:    null.StringFrom(w.VersionMin),
			VersionMax:    null.StringFrom(w.VersionMax),
			MetaJson:      null.StringFrom(w.MetaJson),
			ProvidersJson: null.StringFrom(w.ProvidersJson),
			CreatedAt:     null.TimeFrom(w.CreatedAt),
			UpdatedAt:     null.TimeFrom(w.UpdatedAt),
		}
	}

	res := ModelsN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "model_id":
			res.ModelId = null.StringFrom(w.ModelId)
		case "name":
			res.Name = null.StringFrom(w.Name)
		case "short_name":
			res.ShortName = null.StringFrom(w.ShortName)
		case "description":
			res.Description = null.StringFrom(w.Description)
		case "avatar_url":
			res.AvatarUrl = null.StringFrom(w.AvatarUrl)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "version_min":
			res.VersionMin = null.StringFrom(w.VersionMin)
		case "version_max":
			res.VersionMax = null.StringFrom(w.VersionMax)
		case "meta_json":
			res.MetaJson = null.StringFrom(w.MetaJson)
		case "providers_json":
			res.ProvidersJson = null.StringFrom(w.ProvidersJson)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w Models) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *ModelsN) ToModels() Models {
	return Models{

		Id:            w.Id.Int64,
		ModelId:       w.ModelId.String,
		Name:          w.Name.String,
		ShortName:     w.ShortName.String,
		Description:   w.Description.String,
		AvatarUrl:     w.AvatarUrl.String,
		Status:        w.Status.Int64,
		VersionMin:    w.VersionMin.String,
		VersionMax:    w.VersionMax.String,
		MetaJson:      w.MetaJson.String,
		ProvidersJson: w.ProvidersJson.String,
		CreatedAt:     w.CreatedAt.Time,
		UpdatedAt:     w.UpdatedAt.Time,
	}
}

// ModelsModel is a model which encapsulates the operations of the object
type ModelsModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var modelsTableName = "models"

// ModelsTable return table name for Models
func ModelsTable() string {
	return modelsTableName
}

const (
	FieldModelsId            = "id"
	FieldModelsModelId       = "model_id"
	FieldModelsName          = "name"
	FieldModelsShortName     = "short_name"
	FieldModelsDescription   = "description"
	FieldModelsAvatarUrl     = "avatar_url"
	FieldModelsStatus        = "status"
	FieldModelsVersionMin    = "version_min"
	FieldModelsVersionMax    = "version_max"
	FieldModelsMetaJson      = "meta_json"
	FieldModelsProvidersJson = "providers_json"
	FieldModelsCreatedAt     = "created_at"
	FieldModelsUpdatedAt     = "updated_at"
)

// ModelsFields return all fields in Models model
func ModelsFields() []string {
	return []string{
		"id",
		"model_id",
		"name",
		"short_name",
		"description",
		"avatar_url",
		"status",
		"version_min",
		"version_max",
		"meta_json",
		"providers_json",
		"created_at",
		"updated_at",
	}
}

func SetModelsTable(tableName string) {
	modelsTableName = tableName
}

// NewModelsModel create a ModelsModel
func NewModelsModel(db query.Database) *ModelsModel {
	return &ModelsModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           modelsTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *ModelsModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *ModelsModel) clone() *ModelsModel {
	return &ModelsModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ModelsModel) WithoutGlobalScopes(names ...string) *ModelsModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ModelsModel) WithLocalScopes(names ...string) *ModelsModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *ModelsModel) Condition(builder query.SQLBuilder) *ModelsModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ModelsModel) Find(ctx context.Context, id int64) (*ModelsN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ModelsModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ModelsModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ModelsModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]ModelsN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ModelsModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]ModelsN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"model_id",
			"name",
			"short_name",
			"description",
			"avatar_url",
			"status",
			"version_min",
			"version_max",
			"meta_json",
			"providers_json",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "model_id":
			selectFields = append(selectFields, f)
		case "name":
			selectFields = append(selectFields, f)
		case "short_name":
			selectFields = append(selectFields, f)
		case "description":
			selectFields = append(selectFields, f)
		case "avatar_url":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "version_min":
			selectFields = append(selectFields, f)
		case "version_max":
			selectFields = append(selectFields, f)
		case "meta_json":
			selectFields = append(selectFields, f)
		case "providers_json":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*ModelsN, []interface{}) {
		var modelsVar ModelsN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &modelsVar.Id)
			case "model_id":
				scanFields = append(scanFields, &modelsVar.ModelId)
			case "name":
				scanFields = append(scanFields, &modelsVar.Name)
			case "short_name":
				scanFields = append(scanFields, &modelsVar.ShortName)
			case "description":
				scanFields = append(scanFields, &modelsVar.Description)
			case "avatar_url":
				scanFields = append(scanFields, &modelsVar.AvatarUrl)
			case "status":
				scanFields = append(scanFields, &modelsVar.Status)
			case "version_min":
				scanFields = append(scanFields, &modelsVar.VersionMin)
			case "version_max":
				scanFields = append(scanFields, &modelsVar.VersionMax)
			case "meta_json":
				scanFields = append(scanFields, &modelsVar.MetaJson)
			case "providers_json":
				scanFields = append(scanFields, &modelsVar.ProvidersJson)
			case "created_at":
				scanFields = append(scanFields, &modelsVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &modelsVar.UpdatedAt)
			}
		}

		return &modelsVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	modelss := make([]ModelsN, 0)
	for rows.Next() {
		modelsReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		modelsReal.original = &modelsOriginal{}
		_ = query.Copy(modelsReal, modelsReal.original)

		modelsReal.SetModel(m)
		modelss = append(modelss, *modelsReal)
	}

	return modelss, nil
}

// First return first result for given query
func (m *ModelsModel) First(ctx context.Context, builders ...query.SQLBuilder) (*ModelsN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new models to database
func (m *ModelsModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all modelss to database
func (m *ModelsModel) SaveAll(ctx context.Context, modelss []ModelsN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, models := range modelss {
		id, err := m.Save(ctx, models)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a models to database
func (m *ModelsModel) Save(ctx context.Context, models ModelsN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, models.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new models or update it when it has a id > 0
func (m *ModelsModel) SaveOrUpdate(ctx context.Context, models ModelsN, onlyFields ...string) (id int64, updated bool, err error) {
	if models.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, models.Id.Int64, models, onlyFields...)
		return models.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, models, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ModelsModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ModelsModel) Update(ctx context.Context, builder query.SQLBuilder, models ModelsN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, models.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *ModelsModel) UpdateById(ctx context.Context, id int64, models ModelsN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, models.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *ModelsModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *ModelsModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
