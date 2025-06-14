// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: recipe.sql

package repository

import (
	"context"
)

const filterRecipesByTagNamesAndParams = `-- name: FilterRecipesByTagNamesAndParams :many
SELECT r.id, r.name, r.time, r.difficulty
FROM recipes r
WHERE
  -- Min preparation time (optional)
  ($1::int = 0 OR r.time >= $1::int)

  -- Max preparation time (optional)
  AND ($2::int = 0 OR r.time <= $2::int)

  -- Min difficulty (optional)
  AND ($3::int = 0 OR r.difficulty >= $3::int)

  -- Max difficulty (optional)
  AND ($4::int = 0 OR r.difficulty <= $4::int)

  -- Type 1 (Dieta): OR within, AND across types
  AND ($5::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 1
      AND t.name = ANY($5::text[])
  ))

  -- Type 2 (Region)
  AND ($6::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 2
      AND t.name = ANY($6::text[])
  ))

  -- Type 3 (Rodzaj)
  AND ($7::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 3
      AND t.name = ANY($7::text[])
  ))

  -- Type 4 (Alergie): must NOT include any of these
  AND ($8::text[] IS NULL OR NOT EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 4
      AND t.name = ANY($8::text[])
  ))

  -- Type 5 (Składniki_odżywcze)
  AND ($9::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 5
      AND t.name = ANY($9::text[])
  ))

  -- Type 6 (Inne)
  AND ($10::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 6
      AND t.name = ANY($10::text[])
  ))

ORDER BY r.id LIMIT $12::int OFFSET $11::int
`

type FilterRecipesByTagNamesAndParamsParams struct {
	MinTime       int32    `json:"min_time"`
	MaxTime       int32    `json:"max_time"`
	MinDifficulty int32    `json:"min_difficulty"`
	MaxDifficulty int32    `json:"max_difficulty"`
	Diet          []string `json:"diet"`
	Region        []string `json:"region"`
	RecipeType    []string `json:"recipe_type"`
	Allergies     []string `json:"allergies"`
	Nutrients     []string `json:"nutrients"`
	Others        []string `json:"others"`
	RecipesOffset int32    `json:"recipes_offset"`
	RecipesLimit  int32    `json:"recipes_limit"`
}

type FilterRecipesByTagNamesAndParamsRow struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	Time       int32  `json:"time"`
	Difficulty int32  `json:"difficulty"`
}

func (q *Queries) FilterRecipesByTagNamesAndParams(ctx context.Context, arg FilterRecipesByTagNamesAndParamsParams) ([]FilterRecipesByTagNamesAndParamsRow, error) {
	rows, err := q.db.Query(ctx, filterRecipesByTagNamesAndParams,
		arg.MinTime,
		arg.MaxTime,
		arg.MinDifficulty,
		arg.MaxDifficulty,
		arg.Diet,
		arg.Region,
		arg.RecipeType,
		arg.Allergies,
		arg.Nutrients,
		arg.Others,
		arg.RecipesOffset,
		arg.RecipesLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterRecipesByTagNamesAndParamsRow
	for rows.Next() {
		var i FilterRecipesByTagNamesAndParamsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Time,
			&i.Difficulty,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecipeWithId = `-- name: GetRecipeWithId :one
SELECT id, name, recipe, ingredients, time, difficulty FROM recipes WHERE id = $1
`

func (q *Queries) GetRecipeWithId(ctx context.Context, id int32) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipeWithId, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Recipe,
		&i.Ingredients,
		&i.Time,
		&i.Difficulty,
	)
	return i, err
}
