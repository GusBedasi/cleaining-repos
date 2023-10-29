package domain

type OptionsBuilder func(*Options)

type Whitelist []string
type Owner string

type Options struct {
	Whitelist             Whitelist
	Owner                 Owner
	ListRepositoryOptions ListRepositoryOptions
	DeleteEnabled         bool
	Filename              string
}

type Type string

const (
	All     Type = "all"
	Private Type = "private"
)

type ListRepositoryOptions struct {
	Quantity int
	Type     Type
}

func WithWhitelist(repos ...string) OptionsBuilder {
	return func(options *Options) {
		options.Whitelist = append(options.Whitelist, repos...)
	}
}

func WithOwner(owner string) OptionsBuilder {
	return func(options *Options) {
		options.Owner = Owner(owner)
	}
}

func WithListRepositoryOptions(quantity int, repositoryType Type) OptionsBuilder {
	return func(options *Options) {
		options.ListRepositoryOptions = ListRepositoryOptions{
			Quantity: quantity,
			Type:     repositoryType,
		}
	}
}

func WithDeleteEnabled(enable bool) OptionsBuilder {
	return func(options *Options) {
		options.DeleteEnabled = enable
	}
}

func WithFilename(fileName string) OptionsBuilder {
	return func(options *Options) {
		options.Filename = fileName
	}
}

func NewOptions(opts ...OptionsBuilder) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
