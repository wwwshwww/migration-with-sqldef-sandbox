package session_finder

type FilteringOptions struct {
	UserId *string
}

type SortingOptions struct {
	Orders []SortingOption
}

type SortingOption bool

type SortingCreatedAtAsc SortingOption
