package format

type Formatter interface {
	Format(v interface{}) error
}
