package user_finder

type FilteringOptions struct {
	NamePrefix  *string
	NamePartial *string
	NameExact   *string
}

type SortingOptions struct {
	Orders []SortingOption
}

type SortingOption bool

type SortingNameID SortingOption
type SortingNameAsc SortingOption
