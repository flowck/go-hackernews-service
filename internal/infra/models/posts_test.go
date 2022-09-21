// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPosts(t *testing.T) {
	t.Parallel()

	query := Posts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPostsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Posts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PostExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Post exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PostExists to return true, but got false.")
	}
}

func testPostsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	postFound, err := FindPost(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if postFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPostsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Posts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPostsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Posts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPostsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Posts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPostsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func postBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Post) error {
	*o = Post{}
	return nil
}

func testPostsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Post{}
	o := &Post{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, postDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Post object: %s", err)
	}

	AddPostHook(boil.BeforeInsertHook, postBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	postBeforeInsertHooks = []PostHook{}

	AddPostHook(boil.AfterInsertHook, postAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	postAfterInsertHooks = []PostHook{}

	AddPostHook(boil.AfterSelectHook, postAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	postAfterSelectHooks = []PostHook{}

	AddPostHook(boil.BeforeUpdateHook, postBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	postBeforeUpdateHooks = []PostHook{}

	AddPostHook(boil.AfterUpdateHook, postAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	postAfterUpdateHooks = []PostHook{}

	AddPostHook(boil.BeforeDeleteHook, postBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	postBeforeDeleteHooks = []PostHook{}

	AddPostHook(boil.AfterDeleteHook, postAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	postAfterDeleteHooks = []PostHook{}

	AddPostHook(boil.BeforeUpsertHook, postBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	postBeforeUpsertHooks = []PostHook{}

	AddPostHook(boil.AfterUpsertHook, postAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	postAfterUpsertHooks = []PostHook{}
}

func testPostsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(postColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Post
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PostSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Post)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPostToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Post
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Posts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testPostToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Post
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.UserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Posts) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPostsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPostsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPostsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Posts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	postDBTypes = map[string]string{`ID`: `uuid`, `Title`: `character varying`, `UserID`: `uuid`, `Votes`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testPostsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postDBTypes, true, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPostsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postDBTypes, true, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(postAllColumns, postPrimaryKeyColumns) {
		fields = postAllColumns
	} else {
		fields = strmangle.SetComplement(
			postAllColumns,
			postPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PostSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPostsUpsert(t *testing.T) {
	t.Parallel()

	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Post{}
	if err = randomize.Struct(seed, &o, postDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, postDBTypes, false, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err = Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
