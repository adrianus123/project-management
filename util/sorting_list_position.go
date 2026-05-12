package util

import (
	"github.com/adrianus123/project-management/model"
	"github.com/google/uuid"
)

func SortListByPosition(lists []model.List, order []uuid.UUID) []model.List {
	if len(order) == 0 {
		return lists
	}

	ordered := make([]model.List, 0, len(order))

	listMap := make(map[uuid.UUID]model.List)
	for _, list := range lists {
		listMap[list.PublicID] = list
	}

	for _, publicID := range order {
		if list, found := listMap[publicID]; found {
			ordered = append(ordered, list)
		}
	}

	return ordered
}
